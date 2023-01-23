前段时间看了[Koa2](https://github.com/koajs/koa)的代码, 代码没看多少, 倒是看到一个很有意思的现象, 就是在Koa2的代码里有很多的 `get property()` 的使用, 当然这并不是很高深的东西, 就是JS里的 `Setter/Getter`, 一般稍微有点基础的都能看懂. 不过在Koa2里使用的实在太多, 所以我今天就来好好看一下这个 `get`.

首先让我们看一下官网上对于 `get` 方法的定义:

> The `get` syntax binds an object property to a function that will be called when that property is looked up.

也就是说, get方法可以把对象的一个属性绑定到一个函数上, 然后在获取这个属性的时候自动调用这个函数.

![js getter 1](https://img.xhu.me/blog/js%20getter%201.png)

通过Chrome里的终端我们可以看到, 当我们用getter声明一个属性 `test` 的时候, 其实是声明了一个名为 `test` 的占位属性, 以及名为 `get test` 的一个隐藏函数, 获取test属性的时候, 其实就是执行这个函数来获得结果.

## 性能

首先我们来看看Setter/Getter的性能如何:

    var obj = {
      val: 1,

      valByFunc: function () {
        return this.val;
      },

      setValByFunc: function (newVal) {
        this.val = newVal;
      },

      get valByGetter () {
        return this.val
      },

      set valByGetter (newVal) {
        this.val = newVal;
      }
    };

    var start, end;

    start = new Date();
    for (var i = 0; i < 1000000; ++i) {
      let test = obj.val;
    }
    end = new Date();
    console.log(`get val directly: ${end - start} ms`);


    start = new Date();
    for (var i = 0; i < 1000000; ++i) {
      let test = obj.valByFunc();
    }
    end = new Date();
    console.log(`get val by function: ${end - start} ms`);


    start = new Date();
    for (var i = 0; i < 1000000; ++i) {
      let test = obj.valByGetter;
    }
    end = new Date();
    console.log(`get val by getter: ${end - start} ms`);

    console.log('\n**********************************************\n');

    start = new Date();
    for (var i = 0; i < 1000000; ++i) {
      obj.val = 2;
    }
    end = new Date();
    console.log(`set val directly: ${end - start} ms`);


    start = new Date();
    for (var i = 0; i < 1000000; ++i) {
      obj.setValByFunc(2);
    }
    end = new Date();
    console.log(`set val by function: ${end - start} ms`);


    start = new Date();
    for (var i = 0; i < 1000000; ++i) {
      obj.valByGetter = 2;
    }
    end = new Date();
    console.log(`set val by setter: ${end - start} ms`)

执行这个脚本, 可以看到结果为

    get val directly: 7 ms
    get val by function: 19 ms
    get val by getter: 109 ms

    **********************************************

    set val directly: 8 ms
    set val by function: 18 ms
    set val by setter: 131 ms

通过这个结果我们可以看到, 通过函数读取以及改变属性的值, 性能还是能接受的, 因为V8在解释JavaScript代码的时候, 会把没有上下文依赖的函数内联到代码里, 减少运行时创建Scope的性能损失, 最后效果和操作字面量差别不大.

而我们通过getter/setter来对属性进行读写的时候, 却和直接读写有十多倍的差距,
类似于[这样的例子](https://github.com/facebook/immutable-js/issues/21)也看到了几个, 也就是说, 还是有不少人对JavaScript的getter的性能是存在concern的, 但是Google了一下却没看到具体原因, 个人猜测, 这里的慢, 一部分原因是因为我们在调用一个Setter/Getter声明的属性的时候, 解释器并不知道应该去调用get函数还是set函数, 只有当整个AST都完备的时候, 才能知道当前的操作对属性是读值还是写值, 这样一来必然是会存在一定的性能损耗.

## 应用场景

既然Setter/Getter性能有这么明显的问题, 那这种处理属性的方式在什么情况下可以使用呢?

首先就是需要只读属性的时候, 在JavaScript里给对象创建只读属性有几种方法:

1. 使用 `Object.defineProperty`, 代码太繁琐
2. 使用函数, 将函数返回值当做属性值, 调用的时候要加上括号, 总感觉哪里不对
3. 使用Getter, 用函数的方式声明, 取值的时候也比函数更像一个 `属性`

至于什么叫 `更像一个属性`, 我们可以用这段代码来解释:

    > var obj = { a: 1, b: () => 2, get c () { return 3 } };
    undefined
    > obj
    { a: 1, b: [Function: b], c: [Getter] }
    > JSON.stringify(obj)
    '{"a":1,"c":3}'

而通过下面这段代码我们可以发现, 每次调用Getter属性的时候, 其实就是直接执行了一次函数:

    > var obj = { a: 1, b: () => 2, get c () { console.log('test'); return 3 } };
    undefined
    > JSON.stringify(obj)
    test
    '{"a":1,"c":3}'
    > typeof obj.c
    test
    'number'

这样带来的好处就是, Getter的结果虽然是通过函数来获取, 不过却是被当做一个字面量存在的. 不过同时我们上面也提到过, 频繁调用Getter属性其实存在性能隐患, 所以这里也需要小心.

## 总结

1. Getter比函数慢, 大量调用的话要当心, 不过一般使用没问题
2. 如果需要有 `inspect` 这类的方法来打印对象内容的话, 就直接用Getter吧, 感觉除此之外也很难有更好的办法了.

总结一下, Getter的存在让我们可以在一定程度上对对象属性的操作变得优雅, 暴露只读属性轻松随意, 至少不需要函数那样让代码里括号满天飞, 但是性能问题也是要考虑的. 这两者很难兼得, 真正开发的话, 还是需要考虑到应用场景, 尽量找到解决实际问题的平衡点吧.
