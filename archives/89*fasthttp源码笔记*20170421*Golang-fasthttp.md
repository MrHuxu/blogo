最近状态吃屎, 想不到要做什么, 每天陷入了一种碌碌无为的状态, 准确来说就是瞎JB忙又不知道在忙什么, 所以决定看点源码提高一下自己.

前段时间看到一个观点, 蜻蜓点水式的阅读作用基本等于0, 所以还是写下来的好. 因为是边看边写, 可能想到哪儿就写到哪儿, 所以行文可能会比较乱, 不过鉴于这个blog从来也没什么人看, 所以就先这么放着了.

---

这次选的源码是[fasthttp](https://github.com/valyala/fasthttp), 首先简介就挺唬人的:

> Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http

简单的说就是这个库重新实现了Golang标准库中的net/http部分, 然后速度快的令人发指(当然后面的benchmark表明这句话并没有吹牛), 看到这里大家肯定会有疑问, 因为按照一般的常识来说, 一个框架封装的层数和速度绝对是成反比的, 比如在这个例子里, net/http包是建立在源码基础上的, 而一个建立在net/http之上的包怎么做到比net/http还快的呢? 那就只能钻进源码里看了.

首先上测试代码:

    package main

    import (
      "fmt"
      "github.com/valyala/fasthttp"
      "net/http"
    )

    func httpHandler(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello World")
    }

    func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
      fmt.Fprintf(ctx, "Hello World")
    }

    func main() {

      go func() {
        http.HandleFunc("/", httpHandler)
        http.ListenAndServe(":8080", nil)
      }()

      fasthttp.ListenAndServe(":8081", fastHTTPHandler)
    }

我们可以看到, net/http对请求是根据path来处理的, 针对一个path我们指定一个函数`func(http.ResponseWriter, *http.Request){}`, 对response的操作是第一个参数, request的是第二个参数; 而fasthttp把请求的处理都放到了一个统一的`fasthttp.RequestCtx`的参数里.

参数的内部挺复杂的, 暂时就不看了, 先往后看吧. 接下来就是`ListenAndServe`方法了:

    func ListenAndServe(addr string, handler RequestHandler) error {
      s := &Server{
        Handler: handler,
      }
      return s.ListenAndServe(addr)
    }

在这个方法内部实例化了一个`Server`对象, 并且调用这个对象的`ListenAndServe`方法.

继续进入Server结构体内部:

    func (s *Server) ListenAndServe(addr string) error {
      ln, err := net.Listen("tcp4", addr)
      if err != nil {
        return err
      }
      return s.Serve(ln)
    }

可以看到, 这个方法内部是调用net包的原生`Listen`方法来监听地址, 再把监听器作为参数执行`Serve`方法.

---

在开始阅读fasthttp的`Serve`方法前, 我们不妨先看一下net/http的原生Serve方法:

    func (srv *Server) Serve(l net.Listener) error {
      defer l.Close()

      //...

      for {
        rw, e := l.Accept()
        if e != nil {
          ...
        }
        tempDelay = 0

        c := srv.newConn(rw)
        c.setState(c.rwc, StateNew) // before Serve can return
        go c.serve(ctx)
      }
    }

我们可以看到, 在原生的`Serve`方法里, 使用了一个for...select来监听listener, 每一次请求都会生成一个`net.Conn`对象, 并且启动一个新的goroutine, 在goroutine执行这个对象的`serve`方法.

这里我们可以看到, 当我们开始监听一个端口后, 因为http协议无状态的特性, 对于每一次请求, 都会在一个独立的goroutine中处理, 这使得Golang原生就支持高性能并发地处理网络请求.

---

那我们再来看看fasthttp的`Serve`方法:

    func (s *Server) Serve(ln net.Listener) error {
      // ...
      maxWorkersCount := s.getConcurrency()
      s.concurrencyCh = make(chan struct{}, maxWorkersCount)
      wp := &workerPool{
        WorkerFunc:      s.serveConn,
        MaxWorkersCount: maxWorkersCount,
        LogAllErrors:    s.LogAllErrors,
        Logger:          s.logger(),
      }
      wp.Start()

      for {
        if c, err = acceptConn(s, ln, &lastPerIPErrorTime); err != nil {
          wp.Stop()
          if err == io.EOF {
            return nil
          }
          return err
        }
        if !wp.Serve(c) {
          //...
        }
        c = nil
      }
    }

精简掉部分代码之后, 我们可以看到, fasthttp和net/http最主要的区别是, fasthttp并没有直接执行`net.Conn.serve`方法, 而是通过初始化一个`workerPool`对象后, 使用`wp.Serve`方法来处理这次请求, 那我们继续深入这个方法:

    func (wp *workerPool) Serve(c net.Conn) bool {
      ch := wp.getCh()
      if ch == nil {
        return false
      }
      ch.ch <- c
      return true
    }

这里我们可以看到, 当fasthttp处理的时候一个`net.Conn`对象的时候, 并没有直接开始处理, 而是把这个对象扔到了一个channel里, 那么为什么要这么做呢, 我想起之前看克神[这篇博客](http://legendtkl.com/2016/09/06/go-pool/)的时候的时候, 看到这样一段叙述:

> golang中的goroutine通过go来启动，goroutine资源和临时对象池不一样，不能放回去再取出来。所以goroutine应该是一直运行着的。需要的时候就运行，不需要的时候就阻塞，这样对其他的goroutine的调度影响也不是很大。而goroutine的任务可以通过channel来传递就ok了。

看来这个channel就是fasthttp里协程池一个worker的入口了, 那么重点就是`getCh`方法:

    func (wp *workerPool) getCh() *workerChan {
      var ch *workerChan
      createWorker := false

      wp.lock.Lock()
      ready := wp.ready  // ready []*workerChan
      n := len(ready) - 1
      if n < 0 {
        if wp.workersCount < wp.MaxWorkersCount {
          createWorker = true
          wp.workersCount++
        }
      } else {
        ch = ready[n]
        ready[n] = nil
        wp.ready = ready[:n]
      }
      wp.lock.Unlock()

      if ch == nil {
        if !createWorker {
          return nil
        }
        vch := wp.workerChanPool.Get()
        if vch == nil {
          vch = &workerChan{
            ch: make(chan net.Conn, workerChanCap),
          }
        }
        ch = vch.(*workerChan)
        go func() {
          wp.workerFunc(ch)
          wp.workerChanPool.Put(vch)
        }()
      }
      return ch
    }

我们可以看到,`wp.ready`是一个保存了可用的worker的channel的数组, 每次会从这里拿一个可用的channel, 当然如果ready数组里没有内容的话, WorkerPool 使用`sync.Pool`声明了一个pool, 每次都会从这个pool里取一个可用channel出来, 如果没有的话, 就创建一个, 在一个新的goroutine里监听这个channel, 并且把这个channel放进pool里.

那么我们再来深入一下`wp.workerFunc`方法:

    func (wp *workerPool) workerFunc(ch *workerChan) {
      var c net.Conn

      var err error
      for c = range ch.ch {
        if c == nil {
          break
        }

        // ...
        c = nil

        if !wp.release(ch) {
          break
        }
      }

      wp.lock.Lock()
      wp.workersCount--
      wp.lock.Unlock()
    }

    func (wp *workerPool) release(ch *workerChan) bool {
      ch.lastUseTime = CoarseTimeNow()
      wp.lock.Lock()
      if wp.mustStop {
        wp.lock.Unlock()
        return false
      }
      wp.ready = append(wp.ready, ch)
      wp.lock.Unlock()
      return true
    }

果然和克神讲的一样, 这里就是会有一个for loop在监听channel, 如果有连接进来就开始处理, 没有的话就阻塞在这里. 在执行完一次连接之后, 就会通过`wp.release`方法把当前的channel放到`wp.ready`数组里, 以供下次使用.

---

以上就是这次关于fasthttp协程池部分的粗略解读, 配合上克神的文章感觉真是获益匪浅, 感想不少, 在这里写一下:

1. Go和NodeJS支持并发实在是太方便了, 一个`go func`一个扔回调就OK, 对于web开发简直是开挂一样的存在;
2. Channel真是一种让菜鸟懵逼, 高手却能玩出花儿的存在, 是玩好Go之前迈不过的坎;
2. Go的异步真是比NodeJS高不知道哪里去, 所以能衍生出协程池这样的高端玩法, 与之相比NodeJS在ES6之前基于libuv对用户透明的异步简直就是玩具了.
