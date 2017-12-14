今天我们来看一个 JavaScript 项目[z](https://github.com/z-pattern-matching/z), 一个 JavaScript 的模式匹配库.

模式匹配是在函数式编程中一个很常见的概念, 首先我们来看 wiki 上的定义:

> Pattern Matching: In computer science, pattern matching is the act of checking a given sequence of tokens for the presence of the constituents of some pattern.

模式匹配的模式是一种预定义好的元素成分, 而模式匹配就是判断一个元素组合是否满足这个模式.

这么说可能还是比较抽象, 那么我们可以用一段代码了看看怎么使用模式匹配来循环一个数组的:

    const { matches } = require('z');

    const traverse = (arr, handler) => {
      matches(arr)(
        (head, tail) => {
          handler(head)
          traverse(tail, handler);
        },
        (head, tail = []) => handler(head)
      );
    };

    traverse([1, 2, 3, 4, 5], ele => console.log(ele));

我们从[z](https://github.com/z-pattern-matching/z)中引入了 `matches` 这个函数, 这个函数的用法是, 首先接受一个被匹配的对象作为参数生成另一个函数 `matches(arr)`, 然后用匹配的模式作为新函数的参数来进行具体的匹配, 这里我们使用了两个模式:

1. 被匹配的数组有头元素 `head` 和尾数组 `tail`, 当符合这个模式时, 我们就使用 `handler` 处理头元素, 并且将尾数组作为参数进行下一次遍历, 当头元素为 `1`, `2`, `3`, `4` 的时候符合这个模式;
2. 被匹配的数组只有头元素, 尾数组为空, 当被匹配数组为 `[5]` 符合这个模式, 同样使用 `handler` 处理头元素.

而例子中我们传入的参数 `handler` 是简单的打印函数, 这样一来我们就实现了把数组中的元素都打印出来.

这样编写的代码虽然比 `for`/`each` 这样的循环略显复杂, 但是在实际的使用过程中, `traverse` 函数是可以复用的, 我们也可以使用不同的 `handler` 函数对头元素进行不同的处理, 这样的复用更加优雅, 并且也符合函数式语言中 `通过函数组成函数` 的思想.

那我们接着来简单阅读一下 z 的源码, 我们从基本的 `matches` 函数入手: