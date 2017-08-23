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

其实关于Getter中的 `this` 归属, 我翻了MDN也没找到明确的定义, 不过根据上面的代码, 我们其实可以得出一个结论, 那就是:

> Getter是随着对象一起定义的, `this` 指向的是最初定义时的对象, 即使通过解构嵌入新对象的方式, `this` 的指向不会改变, 而普通的属性和方法, 在结构嵌入之后, 将会被应用于新的对象.

那么有没有方法, 在一个定义好的对象上加上Getter呢? 当然有了, 这时就要用到 `defineProperty` 大法了:

    const obj = {
      base: 1
    };

    Object.defineProperty(obj, 'plus', { get: function () {
      return this.base + 1;
    } });

    console.log(obj.plus);

这样一来, 我们就成功的给对象加上了一个新的Getter.