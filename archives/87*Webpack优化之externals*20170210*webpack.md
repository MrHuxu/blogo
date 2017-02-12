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

    > webpack-test@0.0.1 webpack /Users/xhu/temp/webpack-test
    > ./node_modules/webpack/bin/webpack.js --config webpack.config.js

    Hash: c7bb12f114f762f31092
    Version: webpack 2.2.1
    Time: 5235ms
        Asset    Size  Chunks                    Chunk Names
    bundle.js  307 kB       0  [emitted]  [big]  main

这里可以看出来, 即使一个简单的组件, 在引入jQuery和React之后, 什么都不做生成的文件都达到了307kB, 