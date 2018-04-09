## 使用场景

在使用 goroutine 的时候, 我们经常需要对 goroutine 进行超时控制, 一般是通过在 select 中加上超时条件来完成.

那么假设我们需要对一组 goroutine 来进行控制呢? 这时就可以使用 `context` 包.

几个常见用法:

1. 需要对一组 goroutine 进行手动取消控制, 使用 `WithCancel` 返回的 `cancelFunc`;
2. 需要对一组 goroutine 进行超时控制, 使用 `WithTimeout` 或者 `WithDeadline`, 其实前者的底层实现是基于后者的;
3. 需要像下传值, 使用 `WithValue`.

使用的时候, 一般是吧 context 作为 goroutine 的第一个参数, 然后使用 select 监听 `Done()` 方法, 然后就可以在外部对 goroutine 进行同一控制:

    package main

	import (
		"context"
		"fmt"
		"time"
	)

	func main() {
		ctx1, cancel1 := context.WithCancel(context.Background())
		go func(ctx context.Context) {
			select {
			case <-ctx.Done():
				println("the goroutine is terminated by the context1")
			}
		}(ctx1)
		cancel1()
		time.Sleep(time.Second / 10)

		ctx2, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
		go func(ctx context.Context) {
			select {
			case <-ctx.Done():
				println("the goroutine is terminated by the context2")
			}
		}(ctx2)
		time.Sleep(time.Second * 2)

		ctx3, _ := context.WithTimeout(context.Background(), time.Second)
		go func(ctx context.Context) {
			select {
			case <-ctx.Done():
				println("the goroutine is terminated by the context3")
			}
		}(ctx3)
		time.Sleep(time.Second * 2)

		ctx4 := context.WithValue(context.Background(), "name", "xhu")
		go func(ctx context.Context) {
			fmt.Printf("the value of key %s is %s\n", "name", ctx.Value("name"))
		}(ctx4)
		time.Sleep(time.Second / 10)

		ctx5 := context.Background()
		ctx6, cancel2 := context.WithCancel(ctx5)
		ctx7 := ctx6
		ctx8, _ := context.WithCancel(ctx7)
		go func(ctx context.Context) {
			select {
			case <-ctx.Done():
				println("the goroutine is terminated by the context6")
			}
		}(ctx8)
		cancel2()
		time.Sleep(time.Second / 10)
	}

## 源码阅读

context 包在 1.7 版本就被加入 Go 标准库, 源码在 `go/src/context/context.go` , 我们今天重点看一下 cancel context 的实现.

首先是 Context 的定义:

	type Context interface {
		Deadline() (deadline time.Time, ok bool)
		Done() <-chan struct{}
		Err() error
		Value(key interface{}) interface{}
	}

当我们来执行 `context.Background()` 的时候, 其实是创建了一个空的 context, 其定义如下:

	type emptyCtx int

	func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
		return
	}

	func (*emptyCtx) Done() <-chan struct{} {
		return nil
	}

	func (*emptyCtx) Err() error {
		return nil
	}

	func (*emptyCtx) Value(key interface{}) interface{} {
		return nil
	}

至于这个地方为什么用 int 而不是一个空的 `struct`, 注释的解释是这样可以保证每个变量的地址不一样, 不过其实因为后面都是指针操作, 所以这块儿其实用 `type emptyCtx struct{}` 也是可以正常工作的.

然后我们可以看到 `Done()` 返回的是一个 nil, 通过上一篇我们可以知道, 使用 `emptyCtx.Done()` 做 select...case 分支的话, 是会一直阻塞下去的.

然后当我们用 `WithCancel` 来创造出一个可以 cancel 的 context 的时候, 调用的代码如下:

	type cancelCtx struct {
		Context                         // 用来放父 context

		mu       sync.Mutex             // 给数据修改加锁
		done     chan struct{}          // 用来表示 Done 信号的 chnnel
		children map[canceler]struct{}  // 存储基于当前 context 的子 context
		err      error                  // context 已经 cancel, 或者其他出错情况, 置上这个字段, 否则为 nil
	}
	
	func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	    c := newCancelCtx(parent)
	    propagateCancel(parent, &c)
	    return &c, func() { c.cancel(true, Canceled) }
    }

    func newCancelCtx(parent Context) cancelCtx {
    	return cancelCtx{Context: parent}
    }

也就是说, 每次我们调用 withCancel 的时候, 都会创建出一个新的 `cancelCtx` 实例, 给这个实例嵌入了父 context, 这样一来的层层嵌套结构, 对于 cancelCtx/deadlineCtx 不需要的方法, 直接使用 emptyCtx 的默认实现就好了.

这段代码最重要的是通过 `propagateCancel` 来传递 parent 的 cancel 信号.

	func propagateCancel(parent Context, child canceler) {
		if parent.Done() == nil {
			return // parent 无法被 cancel
		}
		if p, ok := parentCancelCtx(parent); ok {
			p.mu.Lock()
			if p.err != nil {
				// parent 已经被 cancel
				child.cancel(false, p.err)
			} else {
				if p.children == nil {
					p.children = make(map[canceler]struct{})
				}
				p.children[child] = struct{}{}
			}
			p.mu.Unlock()
		} else {
			go func() {
				select {
				case <-parent.Done():
					child.cancel(false, parent.Err())
				case <-child.Done():
				}
			}()
		}
	}
	
	func parentCancelCtx(parent Context) (*cancelCtx, bool) {
		for {
			switch c := parent.(type) {
			case *cancelCtx:
				return c, true
			case *timerCtx:
				return &c.cancelCtx, true
			case *valueCtx:
				parent = c.Context
			default:
				return nil, false
			}
		}
	}

首先, 当 parent 无法被 cancel 的时候是不需要传递 cancel 信号的, 直接返回即可.

对于下面的条件语句, 最好结合 `parentCancelCtx` 来一起理解, 这个函数就是寻找 parent 所属的 cancelCtx. 对于 cancelCtx 实例, 最近的一个 cancelCtx 就是其本身, 对于 `timerCtx` 实例, 最近的是其成员变量 `cancelCtx`, 对于 `valueCtx`, 就通过 for 循环继续向上追溯. 如果都不是, 第二个返回值就是 false.

1. 当我们可以找到这个 cancelCtx 时:

    首先给当前操作加锁.
    
    - 如果 parent 已经被 cancel, 直接 cancel 子 context 即可
    - 如果 parent 没有被 cancel, 将子 context 加入到 parent 的 `children` 成员变量里.

2. 如果我们找不到 cancelCtx, 就起一个协程来监听 parent 的 Done(), 当有消息时直接 cancel 子 context.

那然后我们再看看 cancelCtx 上 `cancel` 这个函数:

	func (c *cancelCtx) cancel(removeFromParent bool, err error) {
		if err == nil {
			panic("context: internal error: missing cancel error")
		}
		c.mu.Lock()
		if c.err != nil {
			c.mu.Unlock()
			return // 已经被 cancel 了
		}
		c.err = err
		if c.done == nil {
			c.done = closedchan
		} else {
			close(c.done)
		}
		for child := range c.children {
			child.cancel(false, err)
		}
		c.children = nil
		c.mu.Unlock()

		if removeFromParent {
			removeChild(c.Context, c)
		}
	}
	
	func removeChild(parent Context, child canceler) {
	    p, ok := parentCancelCtx(parent)
	    if !ok {
		    return
	    }
	    p.mu.Lock()
	    if p.children != nil {
		    delete(p.children, child)
	    }
	    p.mu.Unlock()
    }

`cancel` 函数的操作, 就是在我们 cancel 一个 context 的时候, 首先将其自身的 `done` 给关掉, 然后将 `children` 的 context 给 cancel 掉, 然后根据 `removeFromParent` 参数决定是否需要从 parent 的 children 中移除当前 context.

当我们手动去 cancel 一个 context 的时候, 是需要将其从 parent 的 children 中移除的, 因为重复 close 一个 channel 会导致 panic, 而这个 context 的 children 就不用移除了, 因为一个 context 是无法被重复 cancel 的, 这样也避免了多余的内存操作.

进行上述操作的时候, 也别忘了加锁以避免并发冲突.

对于 cancelCtx 的讲解就到此为止, context 包的主体部分应该就差不多了, 这个包的设计还是很有意思的, 比如对 `valueCtx`:

	type valueCtx struct {
		Context
		key, val interface{}
	}

	func (c *valueCtx) Value(key interface{}) interface{} {
		if c.key == key {
			return c.val
		}
		return c.Context.Value(key)
	}

并不是我想的会有一个 `map[string]interface{}`, 而也是通过一层层嵌套来构建, 在取值的时候用递归来查询, 不得不说这个包真的是把嵌套/递归这种数据组合和操作方式玩出了花儿.