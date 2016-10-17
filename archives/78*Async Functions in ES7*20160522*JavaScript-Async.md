# Async Functions in ES7

上次在[这篇文章](http://blog.xhu.me/post/67*ES6:%20%E5%9B%9E%E8%B0%83%E5%B0%86%E6%AD%BB,%20Promise%E6%B0%B8%E7%94%9F*20151018*JavaScript-Promise.md)里介绍了[Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise)的应用, 这是ES6里关于异步操作的解决方案, 但是这个方案其实有一些问题:

1. 如果使用的话, 那么基本上Promise的过程会'污染'整个代码, 整个代码会被then填满
2. 基于```resolve```/```reject```函数的处理方式, 看起来还是不够直观

所以在JS的世界里, 很需要一种侵入性不那么强而且易读的异步回调解决方案, ES7的规范里, ECMAScript又提出了async/await, 使用这个方案仍然需要使用Promise来改写异步过程, 但是需要用then来改写执行过程, 而是可以像同步代码一样将结果直接作为执行结果返回, 可阅读性也非常好. 这个方案的[spec](https://tc39.github.io/ecmascript-asyncawait/)已经完全确定下来了, 这里我就对基于这个方案做一个简单的讲解吧.


### 语法

1. 在使用async/await的时候, 应该用async来声明函数/方法/箭头函数
2. await只能用在async函数里, 也就是说, 下面这样的写法是会报错的
    
        function test(promise) {
          await promise;
        }

3. await后面应该接一个promise, 表示会等待这个promise的返回值, 当然也可以接同步方法, 不过就没有意义了

### 示例

    function chainAnimationsPromise(elem, animations) {
      let ret = null;
      let p = currentPromise;
      for(const anim of animations) {
        p = p.then(function(val) {
          ret = val;
          return anim(elem);
        })
      }
      return p.catch(function(e) {
        /* ignore and keep going */
      }).then(function() {
        return ret;
      });
    }
    
首先是Promise, 在这个例子里, 在for循环内部形成了Promise链, 而不像同步编程里,for循环里每一次结构都是等价的,最后结果的获取也是通过Promise进行的, 这样就对传统的代码编写方式进行了很大改动, 而且代码量也多了不少.

    function chainAnimationsGenerator(elem, animations) {
      return spawn(function*() {
        let ret = null;
        try {
          for(const anim of animations) {
            ret = yield anim(elem);
          }
        } catch(e) { /* ignore and keep going */ }
        return ret;
      });
    }

然后是```Generator```的方式, 这个例子在返回一个generator函数后, 虽然代码少了不少, 但是在调用的时候我们需要手动执行```next```方法, 并且一般使用Generator还需要在真正的星号函数外面加一个wrapper层, 这样一来, 还是不算简便.

    async function chainAnimationsAsync(elem, animations) {
      let ret = null;
      try {
        for(const anim of animations) {
          ret = await anim(elem);
        }
      } catch(e) { /* ignore and keep going */ }
      return ret;
    }

最后就是```async/await```方案了, 使用起来非常简单, 就是在声明有异步过程的函数的时候, 加上```async```关键字, 在执行异步操作的地方加上```await```关键字, 这样基本就可以用同步的写法来处理JS里的异步过程了.

### Tips

1. 即使使用async/await, await后面跟的异步过程不能基于回调, 而是需要以Promise的方式来进行, 所以学会Promise是很重要的, 好消息是ES6内置的Promise语法简约而且强大, 可以看[这篇文章](http://blog.xhu.me/post/67*ES6:%20%E5%9B%9E%E8%B0%83%E5%B0%86%E6%AD%BB,%20Promise%E6%B0%B8%E7%94%9F*20151018*JavaScript-Promise.md)
2. 每一次使用await, 都需要在它最近的一个函数声明前加上async声明, 比较绕口是吧, 下面这段代码就是例子:

        # wrong, Unexpected token
        async () => {
          str = [1, 2];
          str.forEach(i => {
            await i
          });
        }

        # correct
        () => {
          str = [1, 2];
          str.forEach(async i => {
            await i
          });
        }
3. await并不是多线程, 内部还是按照JavaScript本身的单线程模型来执行的, 可以看作是上层的语法糖, 所以没有并发加锁的问题存在.

### 总结

从ES6的Promise/Generator到ES7的Async Functions, 我们可以看到ECMAScript已经开始在保证JS的高可用性同时, 也在努力提高JS的可阅读性.

但是其实在我个人看来, 在底层逻辑不变的情况下, 在这么短时间内增加上层语法来掩盖底层模型, 这种做法的合理性其实是有待商榷的, ES6/ES7的很多改动不可谓不是大刀阔斧, 可是在浏览器本身想要支持这些特性要走的路还很远, 而且有些做法我个人是比较不喜欢的, 比如硬要在```对象 <-> 原型```这种集成方式上加上类的概念, 这样完全隐藏了JS这门语言的本质而强行迎合了一把OOP.

但是从另一个方面看, 这样做更容易获取使用传统面向对象编程语言进行同步编程的程序员的芳心, 而且我个人眼界浅薄, 说不定未来就被打脸了, 总之这些改动成败与否, 就交给时间来验证吧.

### refs:
- [Async Functions](https://tc39.github.io/ecmascript-asyncawait/)
- [ES7 Async/Await](http://rossboucher.com/await/#/)

