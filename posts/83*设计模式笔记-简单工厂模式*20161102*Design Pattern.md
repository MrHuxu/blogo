这篇文章是我开始学习设计模式的第一篇, 虽然设计模式(Design Pattern)这个概念在别的语言里远不如在Java里那么举足轻重, 但是稍微看了一点之后, 却让我感觉这些概念其实并不局限于Java或者C#, 而是一些在计算机行业里通行的设计理念, 于是我决定简单的通读一下23种设计模式, 这篇文章便是一个开始.

基本的学习都是围绕着<大话设计模式>这本书来进行的, 不过语言换成了我比较熟悉JavaScript, 并且使用ES6中最新的基于类的面向对象语法.

开始学习之前, 我们先看一下这次需要解决的问题: 编写一个简单的计算器, 能够从终端中读取数字和操作符, 并且输出计算结果.

相信大部分程序员都能很快的写出一个简单的实现:

    const readline = require('readline');

    const rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout
    });

    rl.question('Input number A: ', (A) => {
      rl.question('Input operator(+, -, *, /): ', (operator) => {
        rl.question('Input number B:', (B) => {
          var result;

          if (operator === '+') result = parseInt(A) + parseInt(B);
          if (operator === '-') result = parseInt(A) - parseInt(B);
          if (operator === '*') result = parseInt(A) * parseInt(B);
          if (operator === '/') result = parseInt(A) / parseInt(B);

          console.log('The result is: ', result);
          rl.close();
        });
      });
    });

当然, 这份代码基本上可以满足基本的四则运算了, 但是代码本身可以看到, 还是有很多地方需要改进的:

1. 变量`A`, `B`命名很不规范
2. 判断分支, 每次都有三次无用判断
3. 没有处理被除数是0的意外情况

---

那么我们现在着手来把上面的问题解决一下, 代码如下:

    const readline = require('readline');

    const rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout
    });

    rl.question('Input number A: ', (strNumA) => {
      rl.question('Input operator(+, -, *, /): ', (operator) => {
        rl.question('Input number B: ', (strNumB) => {
          var result;

          try {
            switch (operator) {
              case '+':
                result = parseInt(strNumA) + parseInt(strNumB)
                break;
              case '-':
                result = parseInt(strNumA) - parseInt(strNumB)
                break;
              case '*':
                result = parseInt(strNumA) * parseInt(strNumB)
                break;
              case '/':
                if (strNumB === '0')
                  throw new Error('divided by 0.');
                else
                  result = parseInt(strNumA) / parseInt(strNumB);
                break;
              default:
                break;
            }
          } catch (e) {
            console.log('Error happens:', e.message);
          }

          console.log('The result is:', result);
          rl.close();
        });
      });
    });
    
这样看来, 代码的逻辑是比第一段清晰了不少, 但是这段代码有一个问题, 就是逻辑内聚严重, 比如计算的过程本来和输入输出是独立的, 在这里的输入输出是终端, 如果需要用web端或者其他途径来作为输入输出呢, 这段代码就无法复用了, 所以仍然有改进的空间.

这里我们选择的改进方式就是使用面向对象的方式对代码进行重构.

首先我们复习一下面向对象的三个重要特征, 也就是`封装`, `继承`和`多态`, 我们需要做的就是就是基于这三个概念把程序的耦合度降低, 并且达到如下的目标:

1. 可维护
2. 可复用
3. 可扩展
4. 灵活性好

---

那么首先我们可以做的一个改进就是, 把运算过程从输入输出中独立出来形成一个类:

    // Operatiron 运算类
    class Operation {
      static getResult (numA, numB, operator) {
        var result;
        switch (operator) {
          case '+':
            result = numA + numB;
            break;
          case '-':
            result = numA - numB;
            break;
          case '*':
            result = numA * numB;
            break;
          case '/':
            if (strNumB === '0')
              throw new Error('divided by 0.');
            else
              result = parseInt(strNumA) / parseInt(strNumB);
            break;
        }
        return result;
      }
    }

    // 客户端代码
    const readline = require('readline');

    const rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout
    });

    rl.question('Input number A: ', (strNumA) => {
      rl.question('Input operator(+, -, *, /): ', (operator) => {
        rl.question('Input number B: ', (strNumB) => {
          var result;

          try {
            console.log('The result is:', Operation.getResult(parseInt(strNumA), parseInt(strNumB), operator));
          } catch (e) {
            console.log('Error happens:', e.message);
          }

          rl.close();
        });
      });
    });
    
这样一来, 我们就把运算过程独立出来了, 这样已经是使用了面向对象中的`封装`特性, 看上去已经好不少了, 但是问题仍然是有的, 比如如果我们需要增加一个运算符呢, 我们就无可避免的要去修改`getResult`这个方法, 这样显然不满足易扩展的原则, 所以仍然不算是一个好的解决方案.

---

这时我们就可以使用`继承`和`多态`来实现可扩展的代码, 一下就是改进之后的运算类:

    class Operation {
      constructor(operator) {
        this.operator = operator;
        this.numA = this.numB = 0;
      }

      getResult () {
        return 0;
      }
    }

    class OperationAdd extends Operation {
      getResult () {
        return this.numA + this.numB;
      }
    }

    class OperationSub extends Operation {
      getResult () {
        return this.numA - this.numB;
      }
    }

    class OperationMul extends Operation {
      getResult () {
        return this.numA * this.numB;
      }
    }

    class OperationDiv extends Operation {
      getResult () {
        if (this.numB === 0)
          throw new Error('divided by 0.');
        else
          return parseInt(this.numA) / parseInt(this.numB);
      }
    }
    
这样一来, 当我们需要添加新的运算方式的时候, 只需要添加一个新的类来继承`Operation`类就可以了, 使用的时候, 将相应的类实例化, 设定好数据调用`getResult`方法即可.

但是即使定义好了类, 我们还有一个任务需要完成, 那么就是需要一个东西来决定什么时候去实例化相应的类, 那么这就说到今天要学习的设计模式了, 我们需要给代码添加一个工厂, 而工厂的作用, 简而言之就是:

> 根据输入条件, 决定需要实例化的类.

那么现在就来撸这个工厂吧:

    class OperationFactory {
      static createOperation (operator) {
        var operation;
        switch (operator) {
          case '+':
            operation = new OperationAdd();
            break;
          case '-':
            operation = new OperationSub();
            break;
          case '*':
            operation = new OperationMul();
            break;
          case '/':
            operation = new OperationDiv();
            break;
        }
        return operation;
      }
    }
    
调用的客户端部分也需要做相应改进:

    const readline = require('readline');

    const rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout
    });

    rl.question('Input number A: ', (strNumA) => {
      rl.question('Input operator(+, -, *, /): ', (operator) => {
        rl.question('Input number B: ', (strNumB) => {
          var result;

          try {
            var operation = OperationFactory.createOperation(operator);
            operation.numA = parseInt(strNumA);
            operation.numB = parseInt(strNumB);
            console.log('The result is:', operation.getResult());
          } catch (e) {
            console.log('Error happens:', e.message);
          }

          rl.close();
        });
      });
    });
    
到这里我们就完成了对这个简单计算器的改进.

---

最后我们再总结一下这次的学习, 简单工厂模式其实是设计模式中非常基础的一个, 而且使用的也都是各位程序员非常熟悉的概念, 但是通过对代码的规范编写, 我们完美实现了之前所设定的任务:

1. 可维护, 各种运算方式在独立的类内部, 修改任意一个对别的没有影响
2. 可复用, 运算逻辑和输入输出分离, 运算逻辑暴露统一接口很容易被外部使用
3. 可扩展, 添加运算符只需要建立新的类继承`Operation`并且给工厂添加条件即可
4. 灵活性好, 各种运算符可以在独立的类里面进行自定义的操作, 互不影响

那么到这里这次的学习就完成了, 这里我想再次把书上的一句话重复一边, 与大家共勉:

> 编程是一门技术, 更加是一门艺术.