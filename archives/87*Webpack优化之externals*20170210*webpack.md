在我们对前端代码文件模块进行优化的时候, 一般来说文件的大小和数量都是要考虑的, 原因是:

1. 单个文件太大的话, 加载一个文件的时间过长
2. 文件不大但是文件过多的话, 由于浏览器请求并发限制, 同样会有耗时

不过因为目前的网络环境对于加载一个几百kb的文件并不成问题, 而且出于代码混淆的目的, 所以一般来说, 我们还是会把前端文件压缩成一个, 这时文件的大小就很敏感了.

现在我们一般用Webpack打包文件, 这个工具在让我们享受用ES6的方式构建项目的同时, 也带来了一个问题, 当我们简单的用```import```来引入一个库的时候, 这个构建工具会把整个项目的代码打包进来. 特别是当我们使用jQuery, React这些巨无霸的时候, 以前我们可以通过CDN来引入这些库并且进行缓存, 现在如果把这些代码打包进目标文件, 文件的大小会非常可怕.

我们来做一个实验:

    // src.js
    import $ from 'jquery';
    import React from 'react';
    import { render } from 'react-dom';

    const App = ({}) => (
      <div>
        <h1> this is a test </h1>
      </div>
    );

    render(App, document.getElementById('test'));

    // webpack.config.js
    var path    = require('path');
    var webpack = require('webpack');

    module.exports = {
      entry : ['./src'],

      output : {
        path       : __dirname,
        filename   : 'bundle.js',
      },

      resolve : {
        extensions : ['.js', '.jsx']
      },

      module : {
        loaders : [
          { test: /\.(js|jsx)$/, exclude: /node_modules/, loaders: ['babel-loader'] }
        ]
      },

      plugins : [
        new webpack.optimize.UglifyJsPlugin({minimize: true}),
        new webpack.optimize.OccurrenceOrderPlugin()
      ]
    };

然后我们在终端里把这个文件编译一下:

    >> ./node_modules/webpack/bin/webpack.js --config webpack.config.js
    Hash: c7bb12f114f762f31092
    Version: webpack 2.2.1
    Time: 5235ms
        Asset    Size  Chunks                    Chunk Names
    bundle.js  307 kB       0  [emitted]  [big]  main

这里可以看出来, 即使一个简单的组件, 在引入jQuery和React之后, 即使什么都不做生成的文件都达到了307kB, 实在是有点不能接受.

既然不能接受, 那么我们也需要想一些解决方法, 好在Webpack已经有了成熟的方案了, 就是今天要讲的```externals```. 首先我们来看官方的解释:

> ```externals``` allows you to specify dependencies for your library that are not resolved by webpack, but become dependencies of the output. This means they are imported from the environment during runtime.

也就是说, 写在externals里的外部依赖不会在打包的时候引入, 而是会在运行时通过环境中的变量引入. 这样一来, 我们也可以用CDN来加速这些依赖了. 比如, 我们已经通过CDN在页面上加入了jQuery, React和ReactDom:

    <script src="http://cdn.bootcss.com/jquery/3.1.0/jquery.min.js"></script>
    <script src="http://cdn.bootcss.com/react/15.4.2/react-with-addons.js"></script>
    <script src="http://cdn.bootcss.com/react/15.4.2/react-dom.js"></script>

也就是说, 我们代码的运行环境里, 已经有了```jQuery```, ```React```和```ReactDom```这三个全局变量, 那么我们就可以通过设置externals属性, 让打包之后的代码使用这三个全局变量:

    // webpack.config.js
    var path    = require('path');
    var webpack = require('webpack');

    module.exports = {
      // ...

      externals : {
        'jquery'                        : 'jQuery',
        'react'                         : 'React',
        'react-dom'                     : 'ReactDOM'
      },

      // ...
    };

这时我们可以看一下打包之后的文件大小:

    >> $ ./node_modules/webpack/bin/webpack.js --config webpack.config.js
    Hash: a58c1aca3d27b437ac77
    Version: webpack 2.2.1
    Time: 391ms
        Asset       Size  Chunks             Chunk Names
    bundle.js  741 bytes       0  [emitted]  main

可以看出, 打包之后的文件大小几乎已经可以忽略不计了. 然后我们可以去掉配置文件里的压缩选项, 看一下打包之后的代码:

    // built.js

    // ...

    var _jquery = __webpack_require__(3);

    var _jquery2 = _interopRequireDefault(_jquery);

    var _react = __webpack_require__(1);

    var _react2 = _interopRequireDefault(_react);

    var _reactDom = __webpack_require__(2);

    // ...

    /* 1 */
    /***/ (function(module, exports) {

    module.exports = React;

    /***/ }),
    /* 2 */
    /***/ (function(module, exports) {

    module.exports = ReactDOM;

    /***/ }),
    /* 3 */
    /***/ (function(module, exports) {

    module.exports = jQuery;

    /***/ }),

    // ...

这里就可以很清楚的看出来, 这三个写到了externals里的库, 直接就是用了相应的全局变量了.

关于externals的讲解到这里就结束了, 但是其实一个库被引入的时候, 除了像React这样暴露一个对象, 然后所有的变量和方法都通过这个对象来获取的方式之外, 还有一种方式, 就是在发布的时候, 已经在项目内部把不同的变量和方法通过目录划分好了(比如[Material-UI](http://material-ui.com/), 这样的话, externals就没用了, 不过因为我们通过目录加载只加载需要的模块, 所以对于打包结果大小的影响还是可以接受的.

这两种方式总结一下就是:

1. ```import xxx from 'yyy'```或```import { xxx } from 'yyy'```, 这样引入可以通过将yyy暴露为全局变量以及设置webpack的externals来进行优化
2. ```import xxx from 'yyy/zzz'```, 这样引入打包的时候只会加载zzz目录下的代码, 可以接受, 但是需要小心在zzz里引入了体积很大的库, 必要的话可以对zzz的依赖用方式1来优化(比如我们引入Material-UI的时候, 我们需要的模块本身的代码不多, 但是却依赖React, 不过因为我们已经把React放到了externals里, 所以打包的文件不会增加多少)

最后还有一点, 因为添加externals之后, 依赖的库只有在运行时才能得到, 一般来说是没问题的, 但是有一个特殊情况, 就是当我们使用```react-hot-loader```的时候, 会出现依赖错误, 原因就是react-hot-loader会在本地渲染React组件以对比是否有改变, 需要在打包的时候把库引入进来, 我在Google上找了一下, 解决的方案比较复杂, 所以我的建议就是, 开发和上线使用两套webpack config, 比如我的[这个项目](https://github.com/MrHuxu/bar), 两个模式互不干扰, 这样就一劳永逸了.