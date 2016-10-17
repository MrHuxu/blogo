# Learn bind, call and apply in Javascript

学习这三个方法之前, 首先我们需要明确这三个方法的作用是什么, 排开冗长的定义不谈, 这三个方法在实际中的作用主要是为了改变JS代码中的```this```变量.

我们先看看官方关于this变量的一个定义:

> In most cases, the value of this is determined by how a function is called. It can't be set by assignment during execution, and it may be different each time the function is called.

也就是说, 一个函数里的this变量也就是调用函数的这个对象, 而这三个方法, 就是后面一句话所说的, 可以在运行时动态改变this

---

### bind

这个是定义在```Function```原型的一个方法, 作用就是将参数中的对象作为一个函数的调用对象, 并且返回一个函数.

当然在学习了JS闭包的概念之后, 我们可以换一个方式来理解这个过程, 也就是把调用bind的那个闭包里的this变量换成参数中的对象.

    var obja = {
      txt: 'A',
      func: function () {
        console.log(this.txt);
      }
    };
    var objb = { txt: 'B' };

    obja.func();   // => A
    obja.func.bind(objb)();   // => B

    console.log((function () {
      return `test ${this.txt}`
    }).bind(obja)());   // => test A

---

### call/apply

这两个函数改变this的效果和bind类似, 参数列表中第一个作为this变量, 后面填充进参数列表, 绑定完后立即执行.
    
    var obja = {
      txt: 'A',
      func: function (str) {
        console.log(str + ' ' + this.txt);
      }
    };
    var objb = { txt: 'B' };

    obja.func.call(obja, 'test');   // => test A
    obja.func.apply(objb, ['test']);   // => tet B

    console.log((function () {
      return `test ${this.txt}`
    }).call(obja));   // => test A

不过如果第一个变量是```null/undefined```的话, 就不会改变this了.

    var printThis = function () {
      console.log(this);
    };

    printThis.call(1);   // => [Number: 1]
    printThis.call({});   // => {}
    printThis.call(null);   // => Global env
    printThis.apply(undefined);   // => Global env

另一个比较有意思的是, 同样是接受参数, call接收的是一个参数序列, 而apply接收的是一个参数数组.

    var printTxt = function (str1, str2) {
      console.log(str1 + ' ' + str2);
    };

    printTxt.call(null, 'AA', 'BB');   // => AA BB
    printTxt.apply(null, ['AA', 'BB']);   // => AA BB

利用apply这个特性还有一个特别的用处, 将在下面说明.

---

### Usage

#### 柯里化

上面我们已经说过, bind函数的本质其实改变一个闭包里的内部变量, 那么既然参数也是一个闭包的内部变量, 我们也可以通过bind来改变, 而且当参数数量不够的时候, 返回的函数执行时带的参数将填到不够的位置上, 这种特性非常适合用来给一个参数列表很长的函数进行解构, 拆分成具有特定功能参数较少的函数, 实现函数的柯里华.

    var greet = function (hometown, name) {
      console.log(`Welcome ${name} from ${hometown}!`);
    };

    var greetFromMiami = greet.bind(null, 'Miami');
    greetFromMiami('Mike');   // => Welcome Mike from Miami!
    var greetFromNY = greet.bind(null, 'NY');
    greetFromNY('Jack');   // => Welcome Jack from NY!

额外说一点题外话, 如果不这么做的话, 其实JS里一种比较常见的柯里化是这样的:

    var greet = function (hometown) {
      return function (name) {
        console.log(`Welcome ${name} from ${hometown}!`);
      };
    };

    var greetFromMiami = greet('Miami');
    greetFromMiami('Mike');   // => Welcome Mike from Miami!

#### 借用原生方法

既然这三个函数可以改变函数的调用者, 那么我们实现一些很有意思的黑魔法, 比如我们知道JS中的函数自带```arguments```, 这个变量是一个有着和数组类似的数据结构, 但是却不支持大部分数组原生方法的一个对象, 那么我们可以通过```apply/call```来借用原生的数组方法.

比如把参数列表转换成一个真正的数组:

    var transition = function () {
      var args = Array.prototype.slice.call(arguments);
      console.log(args);
    };

    transition('AA', 'BB', 'CC');   // => [ 'AA', 'BB', 'CC' ]

还有一个常用的黑魔法就是, Ruby的数组有很多有用的方法比如```max```或```min```这些, JS里```Math```库也有这样的方法, 但是这个方法接收的是一个参数序列, 如果我们要用JS给一个数组求最大/最小值怎么办呢?

这时就可以利用apply接收参数数组的特性了:

    arr = [1, 0, 5, 1, 3];

    Math.max(1, 0, 5, 1, 3);   // => 5
    Math.max.apply(null, arr);   // => 5
    Math.min.apply(null, arr);   // => 0