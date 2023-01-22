这是寒假在家的第一个 Go 实验项目, 也就是实现一个 [Sinatra](http://sinatrarb.com/) 风格的 Web 框架.

首先我们来看一下路由定义的方式:

    server.Get("/test/get", getHandler)
    server.Post("/test/post", postHandler)

那么我们该怎么存储这两个路由呢, 首先我们需要明确的是, 一个路由系统要做的任务就是把请求打到合适的 handler 上, 那么最简单的方式, 就是定义一个 Map, key 是路由的路径, 而 value 是 handler, 这样的话, 每次请求的进来的时候, 我们可以立刻拿到对应的 hanlder, 但是这样的方式其实有两个明显的缺点:

1. 首先是需要储存所有的路径字符串, 浪费内存;
2. 请求路径和路由定义需要完全匹配, 无法处理 `/user/:id/infos` 这样的带参数路径.

而在路由系统的实际应用中, 我们会根据资源的区别定义很多前缀相通的路径, 这种含有大量共同前缀的字符串系统, 那么我们可以使用 trie 树来存储, 这样一来首先可以节约内存, 其次因为对路径各级分别操作, 可以处理带参数路径的情况.

那么我们先定义结构体 `Router`:

    const routeUndefined = "Route undefined"

    type Handler func(http.ResponseWriter, *http.Request)

    type Router struct {
      handlers map[string]Handler
      children map[string]*Router
    }

    var defaultIndexRouter = &Router{
      handlers: map[string]Handler{
        "GET": func(w http.ResponseWriter, r *http.Request) {
          w.Write([]byte(routeUndefined))
        },
      },
      children: make(map[string]*Router),
    }

我们来解释一下这两个字段的作用:

1. `handlers map[string]Handler`, 根据不同的请求方法将请求打到不同的 handler 上;
2. `children map[string]*Router`, 将当前路由的下级资源存储到这个 map 里.

那么我们再定义一个结构体 `Engine` 来储存当前的路由系统:

    type Engine struct {
      router *Router
    }

    func DefaultEngine() *Engine {
      engine := &Engine{
        router: &Router{
          handlers: make(map[string]Handler),
          children: map[string]*Router{"/": defaultIndexRouter},
        },
      }
      return engine
    }

当我们使用 `DefaultEngine` 来获得一个 Engine 实例的时候, 其实就是初始化了 trie 树的根, 不过和传统的 trie 树不同的是, 这个树的根节点并不是空, 而是指向 `/` 的路由.

首先我们实现一个工具方法, 就是把一个路径字符串划分成一个数组, 这个数组就是 trie 树上从上向下的路径:

    func getPatterns(path string) []string {
      patternString := strings.Trim(path, "/")
      if len(patternString) == 0 {
        return []string{}
      }
      return strings.Split(patternString, "/")
    }

那么我们再实现一个方法, 就是往这个 trie 树里存储路由的方法:

    func (e *Engine) storeRoute(patterns []string, method string, handler Handler) {
      var router = e.router.children["/"]
      for _, pattern := range patterns {
        _, ok := router.children[pattern]
        if !ok {
          router.children[pattern] = &Router{
            handlers: make(map[string]Handler),
            children: make(map[string]*Router),
          }
        }
        router = router.children[pattern]
      }
      router.handlers[method] = handler
    }

然后我们就可以实现定义路由的 `Get` 和 `Post` 方法了:

    func (e *Engine) Get(path string, handler Handler) {
      e.storeRoute(getPatterns(path), "GET", handler)
    }

    func (e *Engine) Post(path string, handler Handler) {
      e.storeRoute(getPatterns(path), "POST", handler)
    }

然后我们需要使用一种方法, 使接收到的网络请求能够被我们的 trie 树引导, 这样我们就需要用到 `http.ListenAndServe` 方法的第二个参数, 这个参数需要实现如下这个接口:

    type Handler interface {
      ServeHTTP(ResponseWriter, *Request)
    }

这样一来, 所有的请求都会通过 `ServeHTTP` 这个方法分发下去, 那么接下来就好说了, 我们直接在 `Engine` 上实现这个方法就好了, 具体的做法就是将请求路径分隔出来, 再在 trie 树上找到对应的 hanler 即可:

    func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
      patterns := getPatterns(r.URL.Path)
      router := e.router.children["/"]
      for _, pattern := range patterns {
        if _, ok := router.children[pattern]; !ok {
          w.Write([]byte(routeUndefined))
          return
        }
        router = router.children[pattern]
      }

      if _, ok := router.handlers[r.Method]; !ok {
        w.Write([]byte(routeUndefined))
        return
      }
      router.handlers[r.Method](w, r)
    }

[完整的实现代码](https://github.com/MrHuxu/x-go-lab/tree/master/web/engine)

---
