首先我们看一下今天需要解决的问题: 实现 express 的 BodyParser 中间件.

在解决这个问题之前, 我们先看一下, 一个不带任何中间件的 express 服务器, 接收到 post 请求时的情况:

    /*
     *  index.js
     */

    const app = require('express')();

    app.post('/', (req, res) => {
      console.log(req.body);
      res.send('hello world');
    });

    app.listen(8125);

    /*
     *  client:
     *  $ curl --data "{\"test\":\"test\"}" http://localhost:8125/
     *    hello world
     *
     *  server:
     *  $ node index.js
     *    undefined
     */


我们可以看到请求的确打到了 express 服务器上, 但是我们从 `req` 对象里拿到的 `body` 属性却是 undefined, 也就是说, 这个问题其实分为两个部分:

1. 我们需要用一个种方法从 `req` 中读取数据;
2. 用一种文本格式协议比如 json 来 parse 读取的数据并且存到 `body` 属性中.

首先我们来解决第一个问题, 我们先来看看 Nodejs 原生的 `http` 组件的 [createServer](https://nodejs.org/api/http.html#http_http_createserver_requestlistener) 方法, 这个方法创建的对象可以响应 `request` 事件:

    const server = require('http').createServer()
    server.on('request', (req, res) => {});

当然根据文档的内容, 这个函数其实是可以接受 `requestListener` 作为参数, 也就是说, 我们也可以把上面代码中的回调函数直接作为参数创建服务器:

    const server = require('http').createServer((req, res) => {});

而我们再看这个方法返回的 [http.Server](https://nodejs.org/api/http.html#http_class_http_server) 实例, 可以看到参数中 `req` 对象是 [http.IncomingMessage](https://nodejs.org/api/http.html#http_class_http_incomingmessage) 实例, 而这个对象实现了 [stream.Readable](https://nodejs.org/api/stream.html#stream_class_stream_readable) 这个接口, 那么我们就可以通过流读取的方式从 `req` 里读取数据, 代码如下:

    const { createServer } = require('http');

    createServer((req, res) => {
      let chunk = '';
      req.on('data', data => chunk += data.toString());
      req.on('end', () => {
        console.log(chunk);
        res.end('hello world');
      });
    }).listen(8125);

而通过 express 的 [文档](http://expressjs.com/en/4x/api.html#req) 我们可以看到, express 里的 `req` 对象是对原生 `http.ClientRequest` 的扩展, 支持所有的原生方法, 自然也实现了 `stream.Readable` 接口, 所以我们可以用同样的方法获取数据:

    const app = require('express')();

    app.post('/', (req, res) => {
      req.on('data', data => console.log(data.toString()));
      res.send('hello world');
    });

    app.listen(8125);

这样我们就解决第一个问题了. 对于第二个问题, 首先我们来学习一下 express 的中间件工作方式. 根据官方 [文档](http://expressjs.com/en/guide/using-middleware.html) 我们可以知道 express 的中间件工作方式类似于下图一样先进后出


>               --------------------------------------
>               |            middleware1              |
>               |    ----------------------------     |
>               |    |       middleware2         |    |
>               |    |    -------------------    |    |
>               |    |    |  middleware3    |    |    |
>               |    |    |                 |    |    |
>             next next next  ———————————   |    |    |
>      request ————————————> |  handler  | — after handler callback->|
>     response <———————————  |     G     |  |    |    |
>               | A  | C  | E ——————————— F |  D |  B |
>               |    |    |                 |    |    |
>               |    |    -------------------    |    |
>               |    ----------------------------     |
>               --------------------------------------

而上面这个图也可以和官方的 middleware 示例代码结合起来:

    var app = require('express')();

    app.use(function (req, res, next) {
      console.log('Time:', Date.now());
      next();
    });

在这个例子里, 我们给 `app` 添加了一个中间件, 功能是打印出当前的时间, 执行完成后, 就通过 `next` 函数进入上图中的下一层, 如果中间没有别的中间件的话, 就会直接进入我们创建服务器时使用的回调函数.

那么这样我们也可以很容易写出我们自己的 BodyParser:

    const bodyParser = (req, res, next) => {
      let chunk = '';
      req.on('data', data => chunk += data.toString());
      req.on('end', () => {
        req.body = JSON.parse(chunk);
        next();
      });
    };

再将这个中间件和上面的 server 结合起来, 就可以通过 body 属性拿到 post 请求中的数据了, 代码如下:

    const app = require('express')();

    app.use(bodyParser);

    app.post('/', (req, res) => {
      console.log(req.body);
      res.send('hello world');
    });

    app.listen(8125);



### refs:

- [Node.js Documentation](https://nodejs.org/api/)
- [Express API Reference](http://expressjs.com/en/4x/api.html)
- [Middleware Onion Model](https://github.com/kenberkeley/redux-simple-tutorial/blob/master/middleware-onion-model.md)