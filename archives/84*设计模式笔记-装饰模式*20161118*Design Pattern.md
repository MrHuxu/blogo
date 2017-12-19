这次我们来学习一个新的设计模式, 话说在刚开始阅读我们公司代码的时候, 发现有不少文件是以`deco`结尾的, 当时还以为是很高深的用法, 后来问过之后才知道, 这个是单词`decorator`的缩写, 也就是所谓的`装饰器`, 而且这种文件里的代码, 一般是给一个已经存在的类做一些扩展用的, 倒也是名副其实, 而今天要学习的, 就是这些代码使用的, 一个很常见的设计模式, 装饰模式.

照样从一个需求开始:

假设我们在玩一个类似于QQ里的角色装扮的功能, 我需要创建一个人物, 并且给这个角色在客户端中打印出不同搭配的装扮.

首先我们自然的想到要创建一个类来表示人物, 并且根据服装的多样, 在类里定义多个方法, 每个方法都会打印出相应的服装.

`Person`类的定义如下:

    class Person {
      constructor(name) {
        this.name = name;
      }

      wearTShirts() {
        console.log('大T恤 ');
      }

      wearBigTrouser() {
        console.log('垮裤 ');
      }

      wearSneakers() {
        console.log('破球鞋 ');
      }

      wearSuit() {
        console.log('西装 ');
      }

      wearTie() {
        console.log('领带 ');
      }

      wearLeatherShoes() {
        console.log('皮鞋 ');
      }

      show() {
        console.log(`装扮的${this.name}`);
      }
    }
    
同时客户端代码为:


    var person = new Person('xhu');

    console.log('第一种装扮:');
    person.wearTShirts();
    person.wearBigTrouser();
    person.wearSneakers();
    person.show();

    console.log();

    console.log('第二种装扮:');
    person.wearSuit();
    person.wearTie();
    person.wearLeatherShoes();
    person.show();

这样会打印出两套装扮结果:

> 第一种装扮:
大T恤
垮裤
破球鞋
装扮的xhu
第二种装扮:
西装
领带
皮鞋
装扮的xhu

这个实现非常简单, 但是缺点也是显而易见的, 如果要增加一种服装, 就要给Person类增加一个方法来实现, 这显然让代码的可维护性变的很差, 所以仍然有改进的空间.

---

那么首先想到的改进是, 我们可以给服装同意创建一个类, 然后各种服装都是继承这个类, 并且在子类中实现具体内容.

**注: 这里最好使用接口或者抽象类来做, 因为这里的服装其实本身是不需要实现具体功能的, 但是因为这两者都不是ES6的内容, 所以我在这里仍然使用类的继承来实现.**

    class Person {
      constructor (name) {
        this.name = name;
      }

      show () {
        console.log(`装扮得${this.name}`);
      }
    }

    class Finery {
      show () {}
    }

    class TShirts extends Finery {
      show () {
        console.log('大T恤');
      }
    }

    class BigTrouser extends Finery {
      show () {
        console.log('垮裤');
      }
    }

    class Sneakers extends Finery {
      show () {
        console.log('破球鞋');
      }
    }

    class Suit extends Finery {
      show () {
        console.log('西装');
      }
    }

    class Tie extends Finery {
      show () {
        console.log('领带');
      }
    }

    class LeatherShoes extends Finery {
      show () {
        console.log('皮鞋');
      }
    }
    
这时客户端代码就可以分别的去实例化相应的服装然后显示了:

    var person = new Person('xhu');

    console.log('第一种装扮:');
    var tShirts = new TShirts();
    var bigTrouser = new BigTrouser();
    var sneakers = new Sneakers();
    tShirts.show();
    bigTrouser.show();
    sneakers.show();
    person.show()

    console.log();

    console.log('第二种装扮:');
    var suit = new Suit();
    var tie = new Tie();
    var leatherShoes = new LeatherShoes();
    suit.show();
    tie.show();
    leatherShoes.show();
    person.show();
    
但是这个代码还是有一个问题, 就是整个穿衣服的过程都是暴露在外部的, 重复的调用`show`方法显的并不美观, 而且穿衣服的搭配和顺序都有有很多种, 也很难在类定义的时候就把穿什么衣服的规则写好, 所以我们一个迫切的需求就是**把需要的功能按正确的顺序从外部串联起来进行控制**.

---

这就说到我们今天要学习的`装饰模式`了:

> 装饰模式(Decorator), 动态地给一个对象添加一些额外的职责, 就增加功能来说, 装饰模式比生成子类更为灵活.

接下来我们通过一个例子来讲解装饰模式的具体实现, 这里我们一般会声明一个组件`Component`类来作为父类:

    class Component {
      operation () {}
    }

然后定义一个装饰器类, 在这个类里, 我们通过`setComponent`方法来设定需要装饰的组件, 同时也复写了父类中具体的操作代码, 但是不提供实际功能.

    class Decorator extends Component {
      setComponent (component) {
        this.component = component;
      }

      operation () {
        if (this.component) {
          this.component.operation();
        }
      }
    }

最后当具体要实现各个装饰器的时候, 我们只需要复写`operation`方法, 一般在这个方法里, 我们会调用被装饰对象的`operation`方法, 并且加入每个装饰器独有的特定逻辑, 来实现对一个对象进行装饰的目的. 

    class ConcreteDecoratorA extends Decorator {
      operation () {
        super.operation();
        this.addedState = 'New State';   // custom variable, different with decorator B
        console.log('具体装饰对象A的操作');
      }
    }

    class ConcreteDecoratorB extends Decorator {
      operation() {
        super.operation();
        this.AddBehavior();
        console.log('具体装饰对象B的操作');
      }

      AddBehavior () {
        // custom method, different with decorator A
      }
    }

客户端的代码如下:

    var c = new Decorator();
    var dA = new ConcreteDecoratorA();
    var dB = new ConcreteDecoratorB();

    dA.setComponent(c);
    dB.setComponent(dA);
    dB.operation();
    
输出为:

> 具体装饰对象A的操作
具体装饰对象B的操作

---

那么从这个思路出发, 我们也可以对之前的代码进行改写了, 首先是`Person`类和`Finery`类, 我们在接下来的代码中, 把Finery作为装饰器的父类, 而Person就是我们需要装饰的组件:

    class Person {
      constructor (name) {
        this.name = name;
      }

      show () {
        console.log(`装扮的${this.name}`);
      }
    }

    class Finery extends Person {
      decorate (component) {
        this.component = component;
      }

      show () {
        if (this.component) {
          this.component.show();
        }
      }
    }

然后给不同的服装创建具体的装饰器类:

    class TShirts extends Finery {
      show () {
        console.log('大T恤');
        super.show();
      }
    }

    class BigTrouser extends Finery {
      show () {
        console.log('垮裤');
        super.show();
      }
    }

    class Sneakers extends Finery {
      show () {
        console.log('破球鞋');
        super.show();
      }
    }

    class Suit extends Finery {
      show () {
        console.log('西装');
        super.show();
      }
    }

    class Tie extends Finery {
      show () {
        console.log('领带');
        super.show();
      }
    }

    class LeatherShoes extends Finery {
      show () {
        console.log('皮鞋');
        super.show();
      }
    }

而在最后使用的时候, 只需要实例化装饰器来意思装饰`person`对象即可.

    var person = new Person('xhu');

    console.log('第一种装扮:');
    var sn = new Sneakers();
    var bt = new BigTrouser();
    var ts = new TShirts();
    sn.decorate(person);
    bt.decorate(sn);
    ts.decorate(bt);
    ts.show();

    console.log('第二种装扮:');
    var ls = new LeatherShoes();
    var ti = new Tie();
    var su = new Suit();
    ls.decorate(person);
    ti.decorate(ls);
    su.decorate(ti);
    su.show();
    
这时如果我们想换一种搭配, 只要更换装饰器或者改变顺序就可以了.


    console.log('第三种装扮:');
    bt.decorate(person);
    su.decorate(bt);
    su.show();
    
这样就输出西装垮裤的混搭风格了:

> 第三种装扮:
西装
垮裤
装扮的xhu

---

那么现在我们可以来总结一下装饰模式的目的:

> 把类的装饰功能和核心职责区分开, 为已有功能动态地添加更多功能

而且对于具体的装饰过程, 好的实现需要支持这两种特性:

1. 有选择地
2. 顺序地

这里我也给我们公司的`*_deco.rb`文件做一个反思, 虽然这些文件是用deco结尾, 但是却都是用`class.exec`或者`include`这样mixin的方式来实现对现有模型的扩展, 而且在RoR里文件加载的顺序不甚明晰的情况下, 这样做的结果是导致扩展的装饰功能其实都一定程度上污染了被装饰对象本身, 并没有很好的满足这两个特性, 所以才经常有deco文件里方法和原有方法冲突的情况出现.

但是真正如上文中的装饰模式在实际使用中也是有问题的, 就是需要创建的对象太多, 特别是对象上属性太多的时候, 每个属性都要实例化一个装饰器并调用装饰方法, 函数调用栈太深太复杂, 这些都可能会成为实际运行中的性能瓶颈, 所以我们公司对于装饰器的实现, 算是一个比较折中的方案了吧.