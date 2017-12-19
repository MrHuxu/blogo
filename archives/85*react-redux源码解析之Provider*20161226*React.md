在现在React全家桶里, Redux可以说是扮演着无比重要的角色, 这个通过函数式思想实现的单一数据流的状态机, 给React应用提供了一个简洁强大的全局状态管理系统, React和Redux本身的思想都是很简单的, 那么这两个框架是怎么结合起来的呢? 这就不得不说到另一个框架[react-redux](https://github.com/reactjs/react-redux), 正是这个框架把React和Redux有机的结合了起来.


在官方repo里我们可以看到, 这个框架非常简单的只暴露了两个接口:

1. [Provider store](https://github.com/reactjs/react-redux/blob/master/docs/api.md#provider-store)
2. [connect([mapStateToProps], [mapDispatchToProps], [mergeProps], [options])](https://github.com/reactjs/react-redux/blob/master/docs/api.md#connectmapstatetoprops-mapdispatchtoprops-mergeprops-options)

这次我们先探索一下`Provider`.

---

首先看一下官方的例子:

    import { Provider } from 'react-redux';

    ReactDOM.render(
      <Provider store={store}>
        <MyRootComponent />
      </Provider>,
      rootEl
    );

也就是说, 在真正使用的时候, 我们应该用`Provider`这个组件把我们的应用包起来, 这个组件暴露一个props也就是我们使用Redux生成的store.

看上去应该不是很简单, 那么我们直接上源码吧.

---

首先把react-redux的源码clone下来:

    git clone https://github.com/reactjs/react-redux.git && cd react-redux
    git checkout ac01d706dd0b0542a0befd9cd5869c96cd2314dc
    
打开`src/index.js`就可以看到`Provider`的来源了:

    import Provider from './components/Provider'
    import connectAdvanced from './components/connectAdvanced'
    import connect from './connect/connect'

    export { Provider, connectAdvanced, connect }

那么继续打开`./components/Provider.js`这个文件, 找到默认的导出对象:


    export default class Provider extends Component {
      getChildContext() {
        return { store: this.store, storeSubscription: null }
      }

      constructor(props, context) {
        super(props, context)
        this.store = props.store
      }

      render() {
        return Children.only(this.props.children)
      }
    }
    
可以看到Provider就是一个React组件, 使用`only`函数保证只渲染一个子组件, 给子组件的context里加上了`store`和`storeSubsciption`两个字段, 而前者就是通过props传进来的store.

接下来是定义好contexts和propTypes里的数据类型:

    Provider.propTypes = {
      store: storeShape.isRequired,
      children: PropTypes.element.isRequired
    }
    Provider.childContextTypes = {
      store: storeShape.isRequired,
      storeSubscription: PropTypes.instanceOf(Subscription)
    }
    
`children`字段就是React元素这个好说, 而两个`store`字段都是通过一个引入的`storeShape`来定义的, 我们可以在`src/utils/storeShape.js`里找到:

    export default PropTypes.shape({
      subscribe: PropTypes.func.isRequired,
      dispatch: PropTypes.func.isRequired,
      getState: PropTypes.func.isRequired
    })

通过之前对Redux源码的学习, 我们可以看出, 其实这个shape定义的就是一个Redux生成的store对象.

而context里的另一个字段`storeSubscription`, 应该是`Subscription`这个类的实例, 我们可以在`src/utils/Subscription.js`里找到定义:


    export default class Subscription {
      constructor(store, parentSub) {
        this.store = store
        this.parentSub = parentSub
        this.unsubscribe = null
        this.listeners = nullListeners
      }

      addNestedSub(listener) {
        this.trySubscribe()
        return this.listeners.subscribe(listener)
      }

      notifyNestedSubs() {
        this.listeners.notify()
      }

      isSubscribed() {
        return Boolean(this.unsubscribe)
      }

      trySubscribe() {
        if (!this.unsubscribe) {
          // this.onStateChange is set by connectAdvanced.initSubscription()
          this.unsubscribe = this.parentSub
            ? this.parentSub.addNestedSub(this.onStateChange)
            : this.store.subscribe(this.onStateChange)
    
          this.listeners = createListenerCollection()
        }
      }

      tryUnsubscribe() {
        if (this.unsubscribe) {
          this.unsubscribe()
          this.unsubscribe = null
          this.listeners.clear()
          this.listeners = nullListeners
        }
      }
    }

可以看到, 这个类的构造函数接受两个参数: 一个Redux store和父组件的Subscription实例, 我们通过这个类的暴露的一些方法来实现一个观察者模式, 也就是在React组件中监听store的改变:


Method | Feature
---|---
trySubscribe() | 如果当前组件没有订阅一个store, 就判断参数里`parentSub`是否存在, 如果存在, 就把`listener`注册到父组件的监听器里, 如果不存在, 就直接订阅store. 订阅的同时获得相应的`unsubscribe`方法. 然后创建组件自己的监听器集合.
tryUnsubscribe() | 执行`trySubscribe`获得的`unsubcribe`方法取消对store的订阅, 清空当前组件的监听器集合.
addNestedSub(listener) | 首先将当前组件订阅到store上, 然后把参数中的`listener`添加到当前组件的监听器集合中
notifyNestedSubs() | 通知所有监听器, 触发监听器中所有的函数
isSubscribed() | 当前组件是否已经订阅store

这时我们来看看用来创建监听器集合的`createListenerCollection`方法:

    const CLEARED = null
    const nullListeners = { notify() {} }

    function createListenerCollection() {
      // the current/next pattern is copied from redux's createStore code.
      // TODO: refactor+expose that code to be reusable here?
      let current = []
      let next = []

      return {
        clear() {
          next = CLEARED
          current = CLEARED
        },

        notify() {
          const listeners = current = next
          for (let i = 0; i < listeners.length; i++) {
            listeners[i]()
          }
        },

        subscribe(listener) {
          let isSubscribed = true
          if (next === current) next = current.slice()
          next.push(listener)

          return function unsubscribe() {
            if (!isSubscribed || current === CLEARED) return
            isSubscribed = false

            if (next === current) next = current.slice()
            next.splice(next.indexOf(listener), 1)
          }
        }
      }
    }

每次执行这个方法之后, 就会返回一个对象, 这个对象包含几个简单的方法, 就像代码里注释所说的, 其实这段代码完全就是Redux源码的翻版, 读过Redux源码的话, 这段里的几个方法也是不难理解的


Method | Feature
---|---
clear() | 清空当前的监听器集合
notify() | 通知所有的监听器, 触发相应函数
subscribe(listener) | 将`listener`放到监听器集合里, 注意这里的给集合添加元素的是一个immutable操作, 同时返回相应的`unsubscribe`方法.

---

到这里, 我们就完成了react-redux源码中Provider的部分, 这个API的作用简单来说就是如下两点:

1. 把Redux store绑定到组件的context里
2. 同时给组件的context里添加`Subscription`类的实例, 而这个实例里就包含了对store的订阅操作.

文章到这里就该结束了, 扯个题外话, React的[官方文档](https://facebook.github.io/react/docs/context.html)里明确说明, 如果熟悉Redux的话, 最好使用Redux作为全局的数据管理而不是通过context传递, 但是react-redux作为官方binding库, 为了实现Redux和React的结合又用到了context. 换句话说, 官方叫你用Redux别用context, 但是用Redux又必须先用context, 这不得不说是`毅种循环`.