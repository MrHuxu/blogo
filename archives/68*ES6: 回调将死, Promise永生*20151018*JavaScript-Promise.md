# ES6: 回调将死, Promise永生

首先我必须要承认标题是有哗众取宠的嫌疑，回调这个概念在JavaScript里已经存在了很久，而且必然会长期存在下去，而```Promise```是我在认真学习Javascript之后才知道的名词，[这里](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise)是MDN上关于ES6规范中关于Promise的详细介绍，首先我们看Promise的简介

> The Promise object is used for deferred and asynchronous computations. A Promise represents an operation that hasn't completed yet, but is expected in the future.

简单的说，Promise是一种处理异步流程的新方式，对于一个异步操作，我们把这个操作看成一个承诺(Promise)，而对于这个承诺，我们事先已经准备好了```解决(resolve)```和```拒绝(reject)```的函数。当操作完成时，传统的操作是把结果作为参数传入回调函数，并且手动检测结果中有没有error来进行下一步操作，而在Promise里，正确的操作结果会进入resolve函数处理，而error会自动进入reject处理。

Promise的出现主要解决了```回调地狱```的问题，因为在传统写法里的callback套callback的方式，变成了Promise的连续的then的方式。

话不多说进正题，ES6作为最新的Javascript规范，提供的Promise其实是很基础的，首先我们看一段代码

    var p = new Promise((resolve, reject) => {
      setTimeout(() => {
        resolve('hello');
      }, 2000);
    });

    p.then((str) => {
      console.log(str + ' world');
    });

在这段代码里，我们声明并定义了一个Promise对象，这个对象初始化的时候接收一个函数，在这个函数里包含一个用```setTimeout```模拟的异步过程，这个函数接受两个参数，分别是```resolve```和```reject```，也就是我们对这个Promise给出的解决方案，在这个例子里，我们设定在2秒之后调用resolve函数。

~~然后接下来我们通过调用```then```方法来执行这个异步过程，并且定义了相应的resolve函数，这样当2秒之后，这个函数被执行并被传入了hello作为参数，打印出hello world。~~

为了不误导人, 我本来想把上面这段给删了, 但是这样的误解几乎是新人必然会有的, 所以我还是留了下来, 这段话最大的错误就是, Promise里的异步过程并不是在我们调用```then```方法的时候执行的, 而是在我们声明完Promise之后就立即执行了的.

也就是说整个过程并不是我们通过控制异步过程 -> 拿到结果 -> 执行resolve或reject这么一个```拉取(pull)```结果的过程, 而是执行异步过程 -> 进入```then``` -> 通知resolve或reject函数执行, 这其实是一个```推送(push)```结果的过程.

这里面其实包含着```Reactive Programming```的思想, 一个比较不错的文档请戳[这里](http://www.tuicool.com/articles/73YNNbu).

下面是一个在Promise中再次返回一个Promise的例子：

    var p = new Promise((resolve, reject) => {
      setTimeout(() => {
        resolve('hello');
      }, 2000);
    });

    p.then((str) => {
      console.log(str);
      return new Promise((resolve, reject) => {
        setTimeout(() => {
          resolve(str + ' world');
        }, 2000);
      });
    }).then((str) => {
      console.log(str);
    });
    
    // print 'hello' after 2000ms
    // print 'hello world' after 4000ms
    
当然，如果仅仅只有一两个异步操作的时候，回调的写法也不是不能接受的，但是一旦异步操作的数量一多，层层嵌套的写法实在是惨不忍睹，因此在提供最基本的resolve/reject之外，ES6的Promise还有两个很有用的可以用来优化多个异步操作的原生函数: ```Promise#race()```和```Promise#all()```.
    
1. ```race```函数接受一个Promise数组并且返回一个Promise，只要这个集合里任何一个项完成了操作，便立即执行相应的resolve/reject。

2. ```all```函数接受一个Promise数组并且返回一个Promise，只有这个集合里所有的项都完成操作并且执行resolve函数，才回把结果生成一个数组作为参数执行resolve；一旦有任意一个操作触发了reject，all函数会立即结束并且执行reject。

下面是一段示例代码：

    var p1 = new Promise((resolve, reject) => setTimeout(resolve, 500, 'this is promise1'));
    var p2 = new Promise((resolve, reject) => setTimeout(resolve, 300, 'this is promise2'));
    var p3 = new Promise((resolve, reject) => setTimeout(reject, 100, 'this is promise3'));
    var p4 = new Promise((resolve, reject) => setTimeout(reject, 400, 'this is promise4'));

    Promise.race([p1, p2]).then(value => console.log(value), err => console.log(err));
    Promise.race([p1, p2, p3]).then(value => console.log(value), err => console.log(err));
    Promise.all([p1, p2]).then(value => console.log(value), err => console.log(err));
    Promise.all([p1, p2, p4]).then(value => console.log(value), err => console.log(err));
    /*
      this is promise3   // from 2nd Promise#race
      this is promise4   // from 2nd Promise#all
      this is promise2   // from 1st Promise#race
      [ 'this is promise1', 'this is promise2' ]   // from 1st Promise#all
    */

  [1]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise