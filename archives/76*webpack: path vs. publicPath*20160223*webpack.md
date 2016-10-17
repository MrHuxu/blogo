# webpack: path vs. publicPath

在配置webpack的时候, output配置有两个很容易混淆的项```path```和```publicPath```, 这两个项对于js文件里引用资源文件的配置非常重要, 如果配置的不对, 就很容易发生页面上图片显示不出来的情况.

首先看官方文档的定义:

> output.path

>> The output directory as absolute path (required).

> output.publicPath

>> The publicPath specifies the public URL address of the output files when referenced in a browser

官方的定义很简单, ```path```指定的是打包文件生成的绝对路径, ```publicPath```指定的浏览器中资源文件的相对路径, 也就是通过```/```之后能访问到的路径.

干说是很难理解的, OK, 开始踩坑!

一开始我对这两个项的理解不深, 导致在编写[MySnippets](https://github.com/MrHuxu/MySnippets)的时候都是在碰运气, 因为这个项目有两个运行模式, 分别是通过网页访问以及打包成electron app的运行. 通过网页访问的时候我是通过[webpack-dev-server]来实现对资源文件的打包和hot reloader, 而通过electron app运行的时候, 则是加载webpack打包好的静态文件, 所以配置文件也需要有一点区别.

一开始我的配置文件是这样的:

    // in server.js for webpack-dev-server
    new WebpackDevServer(webpack(config), {
      publicPath         : config.output.publicPath,
      hot                : true,
      historyApiFallback : true
    })

    // in webpack.config.js for webpack
    output: {
      path: path.join(__dirname, 'vendor', 'dist'),
      filename: 'bundle.js',
      publicPath: '/assets/'
    }

这样配置之后, 出现的情况是在网页上能看到图片, 但是在electron app里看不到图片.

在代码中, 我是这样import这个图片的:

    css: require('../vendor/lang-icons/css.svg')

我们先看一下electron的出错提示
    
    GET file:///assets/daf56877fc199c232ae5b82a33c76744.svg net::ERR_FILE_NOT_FOUND


而通过网页查看的时候, 相应的元素是

    <img src="/assets/daf56877fc199c232ae5b82a33c76744.svg" ... />

这里我们可以看到, 不论是webpack-dev-server还是webpack本身, 打包之后都去了```/assets/```这个目录下去找图片文件了, 为什么会这样呢? 干想是没有结果的, 直接看打包之后的bundle.js吧.

首先找到import图片的地方, 我们可以看到webpack打包之后的语句是:

    /* 429 */
    /***/ function(module, exports, __webpack_require__) {

      eval(" ... css: __webpack_require__(446) ... ");

    /***/ },

每一个通过webpack打包的文件其实都是像这样被包裹在一个作用域里并且通过```eval```函数来执行, 这个作用域自己模拟了CommonJS的模块化功能, ```__webpack_require__```就是引入文件的函数, 第一行语句```/* 429 */```中的数字可以看作是这个文件在webpack中的ID, 把这个数字作为参数传入```__webpack_require__```函数就能引入这个文件export的内容.

那么我们看一下446这个文件导出的是什么

    /* 446 */
    /***/ function(module, exports, __webpack_require__) {

      eval("module.exports = __webpack_require__.p + \"68164d05c63b6ef931beeb3a2abb7768.svg\" ...");

    /***/ },

也就是说, 446这个文件export了一个字符串, 字符串的内容就是图片的地址, 而这个字符串是用```__webpack_require__.p```这个变量和图片名称构成的, 那前面这个变量是什么呢? 再往前找:

    /******/  // __webpack_public_path__
    /******/  __webpack_require__.p = "/assets/";

我们可以看到```__webpack_require__.p```这个变量其实就是```publicPath```, 原来这个图片的位置都被放到这个地址下了.

那么为什么会这样处理图片的地址呢, 当然这不是webpack做的, 而是因为我使用了官方的[file-loader](https://github.com/webpack/file-loader)

这是这个loader的示例:

    var url = require("file!./file.png");
    // => emits file.png as file in the output directory and returns the public url
    // => returns i. e. "/public-path/0dcbbaa701328a3c262cfd45869e351f.png"

也就是说, 通过这个loader加载的文件会被放到publicPath的地址下.

这样原因就很明了了, 首先我们看一下webpack打包之后dist目录的结构:

    dist 
      |- 21e42bb0c1fe0de25c5e0e31ab55addb.svg
      |- 486c5fd143446fcd97b8ef382c42a3a6.svg
      |- 68164d05c63b6ef931beeb3a2abb7768.svg
      |- 8bfd5143504e30f2caa6fa5bee6280e9.svg
      |- bundle.css
      |- bundle.js
      |- c0e47f0589e4e8dacc17f3d6795e0866.svg
      |- cbd37e34205327b10f66573a350a0bc3.svg
      |- daf56877fc199c232ae5b82a33c76744.svg
      |- index.html
      |- main.js

我们可以看到, webpack打包完成后, 图片和js/css文件都被打包进了```path```这个目录下, 而且是在同一个级别的, 通过electron运行项目的时候, 其实就是通过一个浏览器浏览index.html, 而图片文件和index.html是在一个目录下的, 自然是无法通过```/assets/```来访问的, 而应该直接通过```/```来访问.

而webpack-dev-server是通过[webpack-dev-middleware](https://webpack.github.io/docs/webpack-dev-middleware.html)来提供打包后的文件, 在除了```outpub.publicPath```之外还有一个直接叫做```publicPath```的参数, 这个参数就是通过网页serve打包后文件的地址, 在这个例子里, 这个地址恰好就是```/assets/```, 所以就可以访问了.

知道原因之后, 解决方案就简单了, 网页访问没问题所以不用改了, 然后我们可以把打包后的静态文件都放到index.html所在目录更下一级的assets目录中, 这样通过electron app也可以访问了.

更新之后的配置文件如下:

    // in server.js for webpack-dev-server
    config.output.publicPath = '/assets/';

    new WebpackDevServer(webpack(config), {
      publicPath         : config.output.publicPath,
      hot                : true,
      historyApiFallback : true
    })

    // in webpack.config.js for webpack
    output: {
      path: path.join(__dirname, 'vendor', 'dist'),
      filename: 'bundle.js'
    },

关于这两个参数的解析就到这里了. 按照惯例, 写个总结:

首先, ```output.publicPath```的值有几种情况:
1. 不设值, 那么资源文件会从相对的根目录加载, electron是html文件的同级, 网页的话则是```/```
2. 值是路径形式
    - 通过```file://```打开网页, 是通过绝对根目录```/```往下寻找路径
    - 通过```http(s)://```打开网页, 是通过网页的```/```往下寻找路径
3. 值是```http(s)://```这样的URL路径, 会直接去该路径下加载文件

然后是一些感想:

1. webpack打包之后的文件其实并不复杂, 除开一些特有的语法, 基本上还是CommonJS的那一套, 只要没做Uglify, 追溯某些问题也是很容易从打包之后的js文件里看出线索的.
2. ```output```里的```publicPath```是访问资源文件的位置, 外面的那个是webpack-dev-server打包文件的访问位置, 这个配置其实还有一个更实用的地方, 就是很容易和CDN结合起来, 也就是上面的第三种情况, 只要把这个值设置为CDN的URL, 那么通过```file-loader```转译的资源文件就都会从CDN加载了.
