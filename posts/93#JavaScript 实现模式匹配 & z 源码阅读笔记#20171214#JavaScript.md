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

那我们接着来简单阅读一下 [z](https://github.com/z-pattern-matching/z) 的源码, 我们从基本的 `matches` 函数入手, 将代码 clone 到本地后, `package.json` 可以看到入口文件是 `src/z.js`, 先看这个文件的内容:


    /*
     *   src/z.js
     */
    const getMatchDetails = require('./getMatchDetails')
    const matchArray = require('./matchArray')
    ...

    const resolveMatchFunctions = (subjectToMatch, functions) => {
      for (let i = 0; i < functions.length; i++) {
        const currentMatch = getMatchDetails(functions[i])

        ...

        const matchHasMultipleArguments = currentMatch.args.length > 1
        if (matchHasMultipleArguments && Array.isArray(subjectToMatch)) {
          const multipleItemResolve = matchArray(currentMatch, subjectToMatch)
          if (
            multipleItemResolve.hasValue &&
            Array.isArray(multipleItemResolve.value)
          ) {
            return currentMatch.func.apply(null, multipleItemResolve.value)
          }

          if (multipleItemResolve.hasValue) {
            return currentMatch.func(multipleItemResolve.value)
          }
        }

        ...
      }
    }

    const matches = (subjectToMatch) => (...functions) =>
      resolveMatchFunctions(subjectToMatch, functions)

    module.exports = { matches }

这里我们只保留了匹配数组的相关代码, 通过最底下部分我们可以看到, `matches` 函数的执行结果返回的仍然是一个函数, 这个函数接收一个函数数列作为参数也就是需要匹配的模式, 然后通过 `resolveMatchFunctions` 进行匹配.

这个 `resolveMatchFunctions` 的功能有:

1. 遍历参数中的函数, 并且通过 `getMatchDetails` 函数转换成真正的匹配模式 `currentMatch`;
2. 根据模式的参数数量和类型来判断使用的match方法, 这里当参数数量大于1的时候, 调用的是处理数组的 `matchArray` 方法;
3. 根据match之后的返回值, 如果结果是数组的话, 就通过 `apply` 方法结构数组作为 `currentMatch.func` 的参数执行, 否则就直接调用该函数处理结果.

这一段代码的逻辑其实很简单, 那么我们来看一下其中两个比较关键的方法, 首先是 `getMatchDetails`:

    /*
     *   src/getMatchDetails.js
     */
    const functionReflector = require('js-function-reflector')

    module.exports = (matchFunction) => {
      const reflectedFunction = functionReflector(matchFunction)

      return {
        args: reflectedFunction.args,
        func: matchFunction
      }
    }

这个函数其实很简单, 就是调用一个 JS 反射库 `js-function-reflector` 把模式函数的参数提取了出来, 放到返回对象的 `args` 字段里, 而 `func` 对象其实就是模式函数本身.

那么这个库的执行结果是什么样的呢? 我们可以用下面这一段示例来看:

    const functionReflector = require('js-function-reflector')

    const matchFunction1 = (head, tail) => {console.log(head); console.log(tail)};
    console.log(functionReflector(matchFunction1).args)   //   => [ 'head', 'tail' ]

    const matchFunction2 = (head, tail = []) => {console.log(head); console.log(tail)};
    console.log(functionReflector(matchFunction2).args)   //   => [ 'head', [ 'tail', [] ] ]

也就是说, 这个函数会把参数中的参数名称和默认值以数组的方式返回.

那么我们再看 `matchArray` 函数:

    /*
     *   src/matchArray.js
     */
    const option = require('./option')
    const match = require('./match')

    module.exports = (currentMatch, subjectToMatch) => {
      const matchArgs = currentMatch.args.map(
        (x, index) =>
          Array.isArray(x) ? { key: x[0], value: x[1], index } : { key: x, index }
      )

      ...

      const heads = Array.from(
        Array(matchArgs.length - 1),
        (x, y) => subjectToMatch[y]
      )
      const tail = subjectToMatch.slice(matchArgs.length - 1)

      ...

      const headsWithArgs = matchArgs.filter(x => x.value)
      for (let i = 0; i < headsWithArgs.length; i++) {
        const matchObject = {
          args: [[headsWithArgs[i].key, headsWithArgs[i].value]]
        }

        const matchResult = match(matchObject, heads[headsWithArgs[i].index])
        if (matchResult === option.None) {
          return option.None
        }
      }

      return option.Some(heads.concat([tail]))
    }

当然这个函数也不太复杂, 主要的功能有:

1. 首先处理模式函数的参数数组, 当某个参数为数组时, 根据 `js-function-reflector` 的例子中可以知道这是因为这个参数有默认值, 那么就把这个默认值保存起来;
2. 根据模式函数参数的长度 `matchArgs.length`, 把前 `matchArgs.length - 1` 个元素对应到参数的 head 部分, 而剩余的元素放到 tail 尾数组里;
3. 对 `matchArgs` 进行遍历, 如果一个参数的 `value` 字段有值即有默认值的时候, 就调用基本的 `match` 进行匹配, 如果有任何一个元素和默认值匹配补上, 直接返回空结果;
4. 当所有元素都匹配的时候, 把所有 head 元素和 tail 尾数组放到一个数组里返回.

那么通过这段代码我们可以分析出, 当一个数组匹配上一个模式的时候, 结果中的 `head.concat([tail])` 会被 `apply` 结构, 正好对应上模式函数的参数, 然后就可以执行函数体了.

不过我们也不难发现, 这个库的使用语法部分是完全基于 ES6 的, 因为函数参数的默认值是在 ES6 才被添加进标准, 而且在看完文档后我也对源码中使用的 `Array.from` 有了更深入的了解, 这个项目用非常简炼的代码实现了函数式编程的一个常见功能, 也算是函数式在 JS 中的一次成功实践吧.