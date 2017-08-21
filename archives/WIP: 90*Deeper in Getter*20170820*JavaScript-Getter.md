在我之前的文章[Getter in JavaScript](http://blog.xhu.me/post/Getter%20in%20JavaScript)里, 已经讲了一些 `Getter` 和传统使用成员方法取值的不同, 这段时间写前端代码比较多, 而且有很多类似于这种 `getXxxx()` 长相的方法, 就想用Getter来体换一下, 结果发现了Getter和方法之间的另一个不同.

首先看一下这段代码:

    const method = {
      getPlus () {
        return this.base + 1;
      }
    }

    const obj = {
      base: 1,
      ...method
    };

    console.log(obj.getPlus());

答案应该显而易见, 输出应该是 `2`.

那么下面这段代码呢?

    const getter = {
      get plus () {
        return this.base + 1;
      }
    };

    const obj = {
      base: 1,
      ...getter
    };

    console.log(obj.plus);

别急, 还有一段代码:

    const base = {
      base: 1
    };

    const obj = {
      ...base,
      get plus () {
        return this.base + 1;
      }
    }

    console.log(obj.plus);

这两段代码看上去和第一段代码没有太大区别, 输出应该还是 `2` 吧? 这个结论正确吗?

我们在终端里用Node来跑一下这段程序就能发现, 地一段代码的结果并不是 `2`, 而是 `NaN`, 而第二段的结果是 `2`, 至于为什么会这样, 我们在代码中打断点输出之后可以发现, `get plus ()` 中, `this.base` 并没有取到值, 而是 `undefined`.