# 当Promise遇上闭包

这个博客的Projects页面，数据并不是写死的，而是从GitHub抓取过来的，因为官方的Nodejs GitHub API库并不支持Promise的方式，所以我一开始是这样写的:

    var fetchReposPromise = new Promise((resolve, reject) => {
      githubAPI.request('...', (data) => {
        resolve(data);
      });
    });

然后在需要请求数据的时候，直接```fetchReposPromise.then```就行了，对这段代码我期望的行为是，每次进入resolve函数的时候，都能重新发一次请求，然而事实并不是这样，我发现每次最后进入resolve函数的都是同一份数据，于是我写了下面一段代码来验证:

    // server
    var express = require('express');
    var app = express();

    app.get('/', (req, res) => {
      res.send({num: Math.random()});
    });

    app.listen(3000, () => {
      console.log('Server is running on port 3000');
    });

    // client
    var request = require('request');

    var promise =  new Promise((resolve, reject) => {
      request.get('http://localhost:3000', function (err, res, body) {
        resolve(body);
      });
    });

    promise.then(value => console.log(value));
    promise.then(value => console.log(value));
    // {"num":0.6926063213031739}
    // {"num":0.6926063213031739}

事实果然是这样，两次then的结果并没有区别，也就是说这个get请求只执行了一次，所以进入resolve函数的data还是一样的。

于是我又写了下面这段代码来验证我的想法:

    var a = 1;
    var p = new Promise(resolve => resolve(a));

    a = 2;
    p.then(value => console.log(value));   // => 1

这个结果出来后，终于算是真相大白了，原来又是闭包！

回到第一段代码，这段代码中的Promise封装了一个异步过程，然后把结果作为参数传入了resolve函数，这段代码看起来每次resolve执行前都会经历请求数据的过程，其实并不是这样，对于一个Promise的resolve和reject函数，之前的执行过程都是**完全不可见**的，当Promise里的异步过程执行完毕时，最后```resolve(data)```其实生成了一个闭包，```data```成了这个闭包的内部变量，之后再调用then函数的时候，都是使用这个data作为参数，所以每次执行的结果都是一样的。

知道了原因，改起来也就轻松了，我们只需要在这个Promise外面套一个function，这样每次执行这个函数的时候都会生成一个新的闭包，以第二段代码中的clien为例:

    var request = require('request');

    var generatePromise = () => {
      return new Promise((resolve, reject) => {
        request.get('http://localhost:3000', function (err, res, body) {
          resolve(body);
        });
      });
    };

    generatePromise().then(value => console.log(value));
    generatePromise().then(value => console.log(value));
    // => {"num":0.4341859801206738}
    // => {"num":0.3852167946752161}