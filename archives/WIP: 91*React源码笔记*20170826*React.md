今天决定看看React的源码, 这个念头其实很早就开始有了, 只是一直没有动力付诸行动, 好好看一下, 算是知其然也知其所以然吧.

这次看的代码基于GitHub上的 [React](https://github.com/facebook/react), commit号为 `c0a41196e68dea0128c887bef8f125925d4777d2`.

虽然 `package.json` 文件没有显式的声明 `main` 字段, 不过通过 `scripts/rollup/build.js` 这个构建脚本, 我们还是可以找到项目的入口是 `src/isomorphic/ReactEntry.js` 文件, 这次我们主要来看看Component相关的源码:

    var ReactBaseClasses = require('ReactBaseClasses');
    ...

    var React = {
      ...
      Component: ReactBaseClasses.Component,
      PureComponent: ReactBaseClasses.PureComponent,
      unstable_AsyncComponent: ReactBaseClasses.AsyncComponent,
      ...
    };
    ...
    module.exports = React;

我们继续找到 `src/isomorphic/modern/class/ReactBaseClasses.js` 文件, 这个文件里可以看到 `Component` 的详细定义:

