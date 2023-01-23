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

    var ReactNoopUpdateQueue = require('ReactNoopUpdateQueue');

    function ReactComponent(props, context, updater) {
      this.props = props;
      this.context = context;
      this.refs = emptyObject;
      this.updater = updater || ReactNoopUpdateQueue;
    }

    ReactComponent.prototype.isReactComponent = {};

    ReactComponent.prototype.setState = function(partialState, callback) {
      ...
      this.updater.enqueueSetState(this, partialState, callback, 'setState');
    };

    ReactComponent.prototype.forceUpdate = function(callback) {
      this.updater.enqueueForceUpdate(this, callback, 'forceUpdate');
    };

这里可以看到, 一个组件就是一个函数, 接收三个参数, 也就是我们在写React的时候最常见的 `props`, 'context', 以及一个我们一般不会用的 `updater`, 通过命名后面的代码可以看出, 这个参数作用就是一个任务队列, 当我们使用 `setState`/`forceUpdate` 的时候, 就会把更新操作放到这个队列里.

正因为这里有一个队列来操作, 所以当我们调用这两个方法的时候, 并不会立即在 `this` 指针上得到修改之后的效果, 我们只能在回调里, 或者 `componentDidUpdate` 方法里来获得更新之后的值.

