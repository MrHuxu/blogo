# Immutable Object in JavaScript - Immutable.js

上篇文章里说到了[Immutable.js](https://facebook.github.io/immutable-js/)这个强大的第三方库, 这个库加上Redux, 基本上就是目前React圈子里的三剑客了, 基本上用了各种解决方案后都会回到这上面来, Redux做状态管理, React做组件渲染, Immutable.js则提供了简单高效的不可变对象来使React在订阅Redux状态之后对变化的响应更加快速.

### Why Immutable.js

首先Immutable.js对我们最经常使用的两个Mutable的数据类型即Object和Array提供了非常好的支持, 也就是Immutable.js里的List和Map, 而且还有OrderedSet, Stack和Range这样方便的内置类型, 而且操作方式的改动基本就是把JS里赋值取值符号换成了set/get操作符, 而且每次操作都是产生一个新的对象而不是在原有对象上进行修改.

当然支持完善使用简单是一方面原因, 作为一个要经常读和取操作的库, 性能当然也是非常重要的, 这里我们分别使用JSON转换, Object#assign, 以及Immutable.js对一个对象进行1000000万次的复制以及取值操作:

    var a = { a: 1 };
    var d1 = new Date();
    for (i = 1; i < 1000000; ++i) {
      var b = JSON.parse(JSON.stringify(a));
      b.a = 2
      var c = b.a;
    }
    var d2 = new Date();
    console.log(d2 - d1);    // 606
    
    var a = { a: 1 };
    var d1 = new Date();
    for (i = 1; i < 1000000; ++i) {
      var b = Object.assign({}, a);
      b.a = 2;
      var c = b.a;
    }
    var d2 = new Date();
    console.log(d2 - d1);   // 200
    
    var Immutable = require('immutable');
    var a = Immutable.Map({ a: 1});
    var d1 = new Date();
    for (var i = 1; i < 1000000; ++i) {
      var b = a.set('a', 2);
      var c = b.get('a');
    }
    var d2 = new Date();
    console.log(d2 - d1);   // 164

通过这组数据可以看出, Immutable.js速度比JSON转换快了太多, 甚至比ES6之后原生的浅复制Object#assign还快.

既然这样, 那么使用Immutable.js看来就是水到渠成了, 那接下来我就对这个库做一个简单的讲解.

### Immutable.js - fromJS(), toJS(), is()

一般当我们使用AJAX的方式从后段获取数据的时候, 获得的都是JSON, 然后parse成JSON对象, 但是我们需要将其转换成Immutable对象供Redux和React使用, 这时我们就可以使用```fromJS```这个方法将JS对象转换成Immutable.js对象:

    > var i = require('immutable')
    undefined
    > i.fromJS({a : [1, 2, 3]})
    Map { "a": List [ 1, 2, 3 ] }
    
我们可以看到, 这个方法的行为是深转换, 不仅最外层被转换成了一个Immutable Map, 内部的数组也被转换成了Immutable List.

当然, 如果要把前端的Immutable.js数据转成JS对象传给后段, 也有一个对应的```toJS```方法:

    > var i = require('immutable')
    undefined
    > var map = i.Map({
    ... list: i.List.of(1, 2, 3)
    ... })
    undefined
    > map.toJS()
    { list: [ 1, 2, 3 ] }
    
这个转换同样是深转换, 对内部的Immutable.js对象仍然有效.

有趣的是, 和JS中原生的```Object#is```方法类似, Immutable.js还提供了一个```is```方法, 不过行为却和原生的方法相反, 是把两个Immutable.js对象进行Mutable的比较:

    > var i = require('immutable')
    undefined
    > Object.is({a: 1}, {a: 1})
    false
    > Object.is([1, 2], [1, 2])
    false
    > i.is(i.Map({a: 1}), i.Map({a: 1}))
    true
    > i.is(i.List.of(1, 2), i.List.of(1, 2))
    true

我们可以看到, 原生JS对对象和数组的操作是Mutable的, 但是```Object#is```操作符的比较却是Immutable的, 而Immutable.js中对Map和List的操作是Immutable的, 但是```Immutable#is```操作符却是Mutable的.

### Immutable.js - List

```List```对应的是JS原生的```Array```, 原生的数组操作都可以在List对象中找到对应的方法, 我们可以通过如下的方式初始化一个List:

    var i = require('immutable')

    var list1 = i.List.of(1, 2, 3)
    var list2 = i.List([1, 2, 3])
    var list3 = i.List.of(...[1, 2, 3])

上面分别用```of```方法, List构造函数, 以及ES6生成数组iterator的方式够早了```List [ 1, 2, 3 ]```, 这就是一个基本的List对象了.

对于这个对象, 我们可以做一些和普通数组一样的操作:

    list.size   // 3
    list.set(0, 0)               // List [ 0, 2, 3 ]
    list.delete(0)               // List [ 2, 3 ]
    list.push(4)                 // List [ 1, 2, 3, 4 ]
    list.pop()                   // List [ 1, 2 ]
    list.unshift(0)              // List [ 0, 1, 2, 3 ]
    list.shift()                 // List [ 2, 3 ]
    list.update(1, i => i * i)   // List [ 1, 4, 3 ]
    list.insert(1, 4)            // List [ 1, 4, 2, 3 ]
    list.clear()                 // List []

从上面的例子我们可以看出, 每次在list上执行方法的时候, 返回的都是一个新的List对象, 而且初始的对象并没有被改变.

而且List对象还有一些以```In```结尾的方法, 这些方法可以对List对象里面做深层次的修改

    > list = i.fromJS([[1, 2, 3], [4, 5]])
    List [ List [ 1, 2, 3 ], List [ 4, 5 ] ]
    > list.setIn([0, 1], 7)
    List [ List [ 1, 7, 3 ], List [ 4, 5 ] ]
    > list.deleteIn([0, 2])
    List [ List [ 1, 2 ], List [ 4, 5 ] ]
    > list.updateIn([1, 0], i => i * i)
    List [ List [ 1, 2, 3 ], List [ 16, 5 ] ]
    
这样我们不用手动的去get内部的对象再操作, 直接使用带In的方法, 这样内部和外部的对象都会复制出一个新对象.

当然List的操作远不止这么些, 更多的方法请看[这里](https://facebook.github.io/immutable-js/docs/#/List).


### Immutable.js - Map

```Map```对应原生JS里的```Object```类型, 也就是键值对, 不过和原生JS不同的是, Map对键的要求比原生宽泛很多, 比如:

    > a = i.Map({})
    Map {}
    > a.set([1], 1)
    Map { [1]: 1 }

根据官方的说法, 认识只要是```值```的对象都可以作为Map里的键, 数组当然也是一种值了, 但是这种写法还是不推荐的.

首先是初始化Map, 这里的方法就比较单一了, 就是使用Map对象的构造方法, 但是这里一个有趣的点是, 在传入成对出现的数组时, 会将pair自动转成键值对:

    var i = require('immtable');
    
    var map1 = Map({key: "value"});
    var map2 = Map([["key", "value"]]);   // map2: Map { "key": "value" }
    
然后是一些常规操作:

    var map = i.Map({a: 1, b: 2, c: 3})
    map.size                          // 3
    map.set('a', 4)                   // Map { "a": 4, "b": 2, "c": 3 }
    map.delete('b')                   // Map { "a": 1, "c": 3 }
    map.update('c', i => i * i)       // Map { "a": 1, "b": 2, "c": 9 }
    map.merge(i.Map({c: 4, d: 5}))    // Map { "a": 1, "b": 2, "c": 4, "d": 5 }
    map.clear()                       // Map {}


当然也缺不了用```In```结尾的方法:

    > var map = i.Map({a: i.Map({b: 1, c: 2}), d: i.Map({e: 3})})
    undefined
    > map.setIn(['a', 'c'], 4)
    Map { "a": Map { "b": 1, "c": 4 }, "d": Map { "e": 3 } }
    > map.deleteIn(['a', 'b'])
    Map { "a": Map { "c": 2 }, "d": Map { "e": 3 } }
    > map.updateIn(['d', 'e'], i => i * i)
    Map { "a": Map { "b": 1, "c": 2 }, "d": Map { "e": 9 } }
    
更多的操作可以看[这里](https://facebook.github.io/immutable-js/docs/#/Map)

### Immutable.js in Use

在日常使用中, Immutable.js还有一个特性是我非常喜欢的, 就是每个操作的返回值都非常唯一, 而不是像原生JS那样随意, Immutable.js的除了取值之外的操作基本上都是返回新生成的对象, 这样可以方便我们写出非常好看的链式调用:

    > var map = i.Map({a: i.Map({b: 1, c: 2}), d: i.Map({e: 3})})
    undefined
    > map.setIn(['d', 'e'], 4).deleteIn(['a', 'b'], 2).updateIn(['a', 'c'], i => --i)
    Map { "a": Map { "c": 1 }, "d": Map { "e": 4 } }

当然, 如果真要在项目中使用Immutable.js的话, 还需要进行PropTypes验证, 这时我们可以使用```react-immutable-proptypes```这个库:

    import ImmutablePropTypes from 'react-immutable-proptypes';

    /**
     * props = {
     *   ids: List [ 1, 2 ]
     *   infos: Map {
     *     1: {
     *       name: 'test1'
     *       infos: Map {1: 4, 2: 2, 3: 1, 4: 2}
     *     },
     *     2: {
     *       name: 'test2'
     *       infos: Map {1: 4, 2: 2, 3: 1, 4: 2}
     *     }
     *   }
     * }
    **/
    class Dashboard extends Component {
      static propTypes = {
        ids   : ImmutablePropTypes.listOf(React.PropTypes.number).isRequired,
        infos : ImmutablePropTypes.mapOf(ImmutablePropTypes.contains({
          name  : React.PropTypes.string.isRequired,
          infos : ImmutablePropTypes.mapOf(React.PropTypes.number).isRequired
        })).isRequired
      };
      ...
    }

到这儿关于```Immutable.js```这个库的讲解就算完成了, 当然这里的内容还是很浅的, 如果想更进一步的了解这个强大的库, 我推荐在看完这篇文章之后, 继续深入学习[官方文档](https://facebook.github.io/immutable-js/docs/#/).




