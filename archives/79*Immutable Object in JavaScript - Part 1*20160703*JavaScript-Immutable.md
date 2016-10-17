# Immutable Object in JavaScript - Part 1

最近因为写了不少JS代码, 并且开始学习Scala, 所以很多时候会忍不住去对比这两门语言, 说老实话, 初学函数式语言真的是有很大的难度, 不过抛开Scala语言繁杂的语法, 函数式语言的通用特性倒是很简单

1. 函数是```First-class Object```, 也就是说函数是作为基本对象存在的
2. 闭包, 也就是能创造出保存上下文的执行环境, 这个一般都是通过函数来实现
3. ```Immutable Object```, 不可变对象

这三个特性里, 前两个是JavaScript原生支持的, 但是第三个却没有很好的支持, 这个特性的含义是, 对于一个对象的任何操作, 都应该产生一个新对象, 而不是在原有的对象上进行修改, JS里的Array和Object显然都不符合这样的要求.

其实目前除了函数式编程语言之外, 大多数编程语言里的对象都是Mutable的, 我的一个猜想是, 在计算机硬件还没有像现在这么廉价的时候, 创建对象在代码里是很昂贵的操作, 因为一定需要额外的内存空间. 然而在现在硬件相对比较便宜的情况下, 如果能够获得极大的开发便利和效率提升, 牺牲一点空间也是允许的.

### Why Immutable

那么为什么要选择不可变对象呢? 简单一个概念好像的确找不到在JS里的应用场景, 在jQuery的时代, 前端应用好像也的确用不到这么高大上的东西, 直到这几年一系列新的前端框架和理念的出现, 其中有两个非常重要:

1. 组件化(React, Angular#Directive)
2. 状态管理(Redux, Angular#ViewModel)

现在一般一个页面都会被分成很多组件, 每个组件都可以进行局部渲染, 避免大规模的重绘, 而这个状态管理就是用来通知组件重新渲染, 在Angular里, dom和数据是动态绑定的, 数据更新通过内置的脏检测机制来告知dom重绘, 这些我们这里不做更深的探讨~~, 因为我已经忘的差不多了~~; 而当我们使用Redux管理状态的时候, 我们需要手动告诉组件状态是否改变, 如果状态只是一些基本类型还好说, 直接使用```===```比较就可以了, 如果是对象呢, 即使改变一个键的值, 因为引用没有改变, 比较的结果仍然是```true```, 这样是无法得知状态是否更新了的.

所以, 人们自然而然想到使用Immutable Object来表示程序状态了.

事实上Redux也的确是这么做的, 我们可以通过[这行代码](https://github.com/reactjs/redux/blob/master/src/combineReducers.js#L146)看出, Redux的确只是简单的比较对象引用是否相等来判断状态是否更新.

### How Immutable - Copy Object

既然JS里原生不提供不可变对象, 而不可变对象有那么多好处, 那么我们自己让一个对象不可变不就好了.

产生不可变对象, 最简单的方式就是改变对象的时候不是在原对象上进行操作, 而是复制出一个新对象进行修改, 这样就获得了对一个新对象的引用, 庆幸的是, 在ES6之后, 我们拥有了```Object#assign```方法, 通过这个方法, 我们就可以很简单的复制一个对象.

### Shallow Copy

```Object#assign```方法的语法很简单, 其实就是和```jQuery#extend```类似, 这个方法接受一个对象序列, 把第二个以及往后的参数里对象依次merge到第一个参数中, 具体的用法是:

    obj1 = { a: 1 };
    obj2 = Object.assign({}, obj1);
    
    obj1.a = 2;
    console.log(obj2['a']);   // 1
    console.log(obj1 === obj2);   // false
    
只要我们在每次处理对象的时候都进行一次assign操作, 就可以把对象看成是immutable的了, 这个方案看上去很完美了, 但是它真的完美吗? 我们看看这段代码

    obj1 = { a: { b: 1 } }
    obj2 = Object.assign({}, obj1)
    
    obj1.a.b = 2
    console.log(obj2.a)   // { b: 2}
    console.log(obj1.a === obj2.a)   // true
    
最后一个表达式的结果并不是我们预想的```false```, 为什么呢? 因为assign这个操作是**浅复制 Shallow Copy**.

什么是浅复制呢? 我们知道, 和Scala不同的是, JS的数据类型并不全是引用类型, 真正的引用类型只有使用```typeof```操作符结果为```object```的数据, 基本类型在JS中是值类型的, 而在使用assign方法的时候, 值类型会被复制, 而引用类型, 却不会复制内存, 仍然只是复制之前对象的引用, 这样一来, 复制之后的结果是没法使用相等来判断内部对象是否更新的.

值得一提的是, 函数类型在复制的时候行为和其他类型都有不同, 看下面这段代码:

    obj1 = { a: x => x };
    obj2 = Object.assign({}, obj1);
    console.log(obj1.a === obj2.a);   // true
    
    obj1.a = x => x * x;
    console.log(obj1.a(3));   // 9
    console.log(obj2.a(3));   // 3
    console.log(obj1.a === obj2.a);   // false
    
    obj1.a = x => x;
    console.log(obj1.a === obj2.a);   // false
    
如果把函数当成引用类型, 那么第一个判断式我们会认为函数a只复制了引用, 但是当改变obj1.a的时候, obj2.a并没有跟着改变, 这时的判断相等也变成了false, 看上去很像是值类型, 而当把obj1.a改变回原值的时候, 判断式仍然为false, 这又像是引用类型的结果.

具体的原因就不深究了, 复制JS对象中, 函数类型是一种特殊的存在.

### Deep Copy

既然有浅拷贝, 那么必然就有**深拷贝 Deep Copy**了, 也就是在复制一个对象的时候, 连对象内部引用的内存也复制一遍.

如果是需要复制的内存只有基本类型的话, 那么有一个非常简单的方法:

    obj1 = { a: { b: 1 } };
    obj2 = JSON.parse(JSON.stringify(obj1));
    console.log(obj1.a === obj2.a);   // false
    
原理就不用说了, 这个方法简单粗暴, 但是缺点也是很明显:

1. 存在性能问题, ```stringify```和```parse```都是耗时操作;
2. 只支持JSON里的数据类型, 不支持函数.

所以这个方案只是一个妥协的结果, 业界已经有了很好的深复制方案, 比如[jQuery#extend](http://api.jquery.com/jQuery.extend/)在扩展对象的时候, 有一个可选的```deep```参数, 当这个参数为true的时候,就是深复制.

那么我们来看一下具体的代码实现吧:

    jQuery.extend = jQuery.fn.extend = function() {
      var options, name, src, copy, copyIsArray, clone,
        target = arguments[ 0 ] || {},
        i = 1,
        length = arguments.length,
        deep = false;

      // 如果第一个参数是bool类型, 那么使用深复制模式
      if ( typeof target === "boolean" ) {
        deep = target;

        // 使用第二个参数作为复制结果
        target = arguments[ i ] || {};
        i++;
      }

      // 如果目标对象不是object类型, 则强制使用一个空对象来放置复制结果
      if ( typeof target !== "object" && !jQuery.isFunction( target ) ) {
        target = {};
      }

      // 如果只有一个参数, 那么把后面的参数扩展到jQuery本身
      if ( i === length ) {
        target = this;
        i--;
      }

      for ( ; i < length; i++ ) {

        // 只复制非空对象
        if ( ( options = arguments[ i ] ) != null ) {

          // 扩展目标对象
          for ( name in options ) {
            src = target[ name ];
            copy = options[ name ];

            // 避免复制过程中出现的环
            if ( target === copy ) {
              continue;
            }

            // 对复制对象中的值进行递归调用, 复制下面所有的对象
            if ( deep && copy && ( jQuery.isPlainObject( copy ) ||
              ( copyIsArray = jQuery.isArray( copy ) ) ) ) {

              if ( copyIsArray ) {
                copyIsArray = false;
                clone = src && jQuery.isArray( src ) ? src : [];

              } else {
                clone = src && jQuery.isPlainObject( src ) ? src : {};
              }

              // 如果是引用类型, 就递归的复制内部的内容
              target[ name ] = jQuery.extend( deep, clone, copy );

            } else if ( copy !== undefined ) {
              target[ name ] = copy;
            }
          }
        }
      }

      // 返回的结果是目标对象
      return target;
    };

这段代码主要就是用递归的方法来把引用类型也做了一遍复制. 而且也检测了复制时是否有环的出现, 实际使用的时候还是相当好用的.

那么还有没有更好用的方法呢? 答案是肯定的, 那就是大名鼎鼎的~~React全家桶里的~~[Immutable.js](https://facebook.github.io/immutable-js/)了, 这个库让产生和操作不可变对象变得非常简单, 具体的用法, 就请看 Part 2 吧~

#### refs

- [jQuery API Documentation](http://api.jquery.com/)
- [知乎 - JavaScript 如何完整实现深度Clone对象？](https://www.zhihu.com/question/47746441)

