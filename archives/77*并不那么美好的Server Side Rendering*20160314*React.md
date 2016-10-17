# 并不那么美好的Server Side Rendering

玩React玩久了, 就不免遇到需要服务端渲染的情况, 本来觉得应该是挺简单个事儿, 其实真正要实现的话, 坑也是不少, 折腾了两天, 倒也是弄出一点成果了, 具体的代码已经放到了[GayHub](https://github.com/MrHuxu/server-rendering-demo)上, 目前已经完成了React和react-router在服务端渲染的任务, Redux尚在探索中, 弄好了会再写一篇的.

之所以想学习服务端渲染, 最主要的原因有两个

1. 现在这个blog不论是对搜索引擎不是那么友好, 文章是通过js异步加载的, 这样很难被搜索引擎收录
2. 移动端不能接受, 不论怎么压缩前端代码打包之后都太大了, 手机看一次1M流量就没了, 简直要命

不过丑话说在前面, 服务端渲染可以解决搜索引擎收录的问题, 但是对减小流量没有任何效果, 如果你的需求是后者的话, 这篇文章就不用看了= =

### Server

首先从服务端开始, 一个比较爽的地方是, ```react-router```已经完全支持服务端渲染了, 那么其实可以抛弃express的那一套路由了, 完全由react-router来接管,

首先是路由文件, 里面将```/```和```/test/```这两个path分别定向到了[Home.jsx](https://github.com/MrHuxu/server-rendering-demo/blob/master/components/Home.jsx)和[Test.jsx](https://github.com/MrHuxu/server-rendering-demo/blob/master/components/Test.jsx):

    // routes/index.js
    import Home from '../components/Home.jsx';
    import Test from '../components/Test.jsx';

    export default {
      path: '/',
      component: Home,
      childRoutes: [{
        path      : 'test',
        component : Test
      }]
    };

然后参考react-router的[官方示例](https://github.com/reactjs/react-router/blob/master/docs/guides/ServerRendering.md)编写的server, 使用```react-router.match```来匹配路由并用```ReactDom.renderToString```方法将对应的JSX文件渲染成HTML并且填入index.ejs里.

    // server.js
    import path from 'path';
    import React from 'react';
    import { renderToString } from 'react-dom/server';
    import { match, RouterContext } from 'react-router';
    import routes from './routes';

    import express from 'express';
    var app = express();

    app.set('views', path.join(__dirname, 'views'));
    app.set('view engine', 'ejs');

    app.use((req, res) => {
      match({ routes, location: req.url }, (error, redirectLocation, renderProps) => {
        if (error) {
          res.status(500).send(error.message)
        } else if (redirectLocation) {
          res.redirect(302, redirectLocation.pathname + redirectLocation.search)
        } else if (renderProps) {
          res.status(200).render('index', {
            markup: renderToString(<RouterContext {...renderProps} />)
          });
        } else {
          res.status(404).send('Not found')
        }
      })
    });

    export default app;

这是项目的启动文件:

    // index.js
    #!/usr/bin/env node
    require('babel-register')({
      presets: ['es2015', 'react']
    });
    var app = require('./server').default;

    var port = process.env.PORT || 16311;
    app.listen(port, () => {
      console.log('==> 🌎  Listening on port %s. Open up http://localhost:%s/ in your browser.', port, port);
    });

这样使用```node .```, 然后打开```http://localhost:16311```就可以看到页面了, 并且点击页面上的链接也是可以跳转的, 但是点击Home页面上的两个button却并没有效果, 为什么呢? 我们看一下网页源代码, 发现只有HTML相关的内容:


    ...
    <div id="container">
      <div data-reactid=".15qqqcmsl4w" data-react-checksum="-170493387"><h3 data-reactid=".15qqqcmsl4w.0"> Home </h3><button data-reactid=".15qqqcmsl4w.1"> set blue </button><button data-reactid=".15qqqcmsl4w.2"> set orange </button><div style="width:100px;height:50px;background-color:#e57373;" data-reactid=".15qqqcmsl4w.3"></div><a class="" href="/test/" data-reactid=".15qqqcmsl4w.4"> to test </a></div>
    </div>
    ...

完全没有js相关的代码, 这两个button肯定就没效果了.

### Client

之所以会出现上面所说的情况, 就是因为我们缺少了客户端渲染这一步骤, 如果没有这一步骤的话, 那么得到的页面只是一个静态页面, 是没有任何动态效果的= =

那就把客户端代码也加上吧, 这个就比较简单了, 引入同样的routes文件渲染一遍就行:

    import React from 'react';
    import { render } from 'react-dom';
    import { Router, browserHistory } from 'react-router';
    import routes from './routes';

    var container = document.getElementById("container");
    render(<Router routes={routes} history={browserHistory} />, container);

使用webpack打包引入之后, 前端页面终于也有了动态效果了.

桥豆麻袋!!  我想你应该已经发现奇怪的地方了, 妈蛋既然我需要重新在客户端渲染, 那么还是需要webpack打包出一个巨大的js文件, 这样服务端渲染还有什么意义啊?

### Explain

我们可以简单的画个流程图来看看目前这个blog的工作形式:

![client_render](https://raw.githubusercontent.com/MrHuxu/img-repo/master/blog/client_render.jpg)

这个可以说是简单SPA的基本流程了, 很大一部分工作都交给浏览器来完成, 这样一来, 当使用curl这样的工具来爬取网站时, 得到的只有一个类似```<div id="container"></div>```的东西而没有具体内容, 这样如果搜索引擎如果没有智能到执行js代码的话, 是获取不到内容的, 同理, 如果在移动端禁用了js代码, 也看不到内容了, sigh.

刚开始知道服务端渲染, 我以为流程是这样的:

![server_render_imagine](https://raw.githubusercontent.com/MrHuxu/img-repo/master/blog/server_render_imagine.jpg)

也就是说, 我以为所谓的后端渲染能把前端相关的代码使用```script```标签这种形式放进HTML代码里并返回, 这样返回的代码不需要任何操作就可以支持React的行为了, [这里](https://github.com/MrHuxu/server-rendering-demo/tree/e5cf0c5b62cb619a9ef9ad5bb6e4b91d9d6e0936)是我一开始写的代码.

然而理想太丰满, react-dom库的```renderToString```方法并没有如我想的一般强大, 这个方法只是将JSX当做一个普通的模板语言, 把React的component转成了HTML的元素而已, 页面上是可以渲染出来, 但是却不带任何动态效果. 如果需要动态效果, 应该怎么做呢? 看了各种对的错的文档之后, 我得到一个很失望的答案, 那就是, 如果需要动态效果, 需要在客户端把React的组件重新渲染一次QAQ

这是现实中的服务端渲染流程:

![server_render_real](https://raw.githubusercontent.com/MrHuxu/img-repo/master/blog/server_render_real.jpg)

正因为存在后面的需要在客户端重新渲染的流程, 所以即使进行过服务端渲染, 客户端还是需要拿到完整的js代码, 所以对于减小流量, 帮助不大, 所以还是只能通过```uglify```和```gzip```这样的方法来压缩js文件体积.

另外, 其实说是重新```渲染```一遍不太恰当, 因为实际上React并没有重绘页面, 在后端渲染好页面的时候, 每个元素都带上了```data-reactid```这个属性,  这样当在客户端执行```ReactDom.render```的时候, 可以快速的将虚拟dom和实际dom进行对比来决定是否更新页面, 所以如果在服务端和客户端使用同一份props来初始化, 那么数据没有改变所以并不会重绘页面, 只是给组件加上动态效果, 性能也不会有太大损失.

### Conclusion

目前能想到的使用服务端渲染的优点有:

1. 搜索引擎友好
2. 服务端渲染页面, 不用客户端绘制dom, 对于比较复杂的页面能有性能提升
3. 客户端和服务端同构(比如, 使用同一套router)

当然, 目前的这个demo还比较简单, 也没有使用Redux来管理数据, 其实在服务端渲染中集成Redux我是有点疑问的:

1. 后端页面已经带上了数据, 那么肯定是不需要用额外的action来fetch数据了, 那么action这个部分感觉没有存在的必要了.
2. 按照Redux官方的做法, 需要对每个request都生成Redux的initialState, 如果每个页面的initialState都不一样, 那么在data -> initialState -> component这样的过程中, 强行加上中间那一环是否有意义呢?

当然这些想法有待验证, 如果有新的成果, 将会在我的下一篇博文里详细说明.╭(●｀∀´●)╯

