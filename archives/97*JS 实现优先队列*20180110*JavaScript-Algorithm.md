今天看了《算法(第四版)》的优先队列一章, 原文是用 Java 编写的, 我们现在来用 JavaScript 实现一下, 也顺便复习一下这章节的内容.

首先我们看一下维基上对于优先队列的定义:

> 优先队列是计算机科学中的一类抽象数据类型。优先队列中的每个元素都有各自的优先级，优先级最高的元素最先得到服务；优先级相同的元素按照其在优先队列中的顺序得到服务。

翻译一下就是, 优先队列和普通队列不同, 队列里的元素是有优先级的, 元素入队之后, 优先级高的先出队.

为了简化代码, 并且假设内部使用数组来存储数据, 其中元素都是整数, 整数的大小就是其优先级, 我们首先来声明构造函数 `PriorityQueue`:

    const PriorityQueue = function () {
      this.data = [];
    };

    PriorityQueue.prototype.enqueue = function (...ele){};
    PriorityQueue.prototype.dequeue = function (){};

接下来就是重点的入队和出队方法.

我们先来思考一下队列内部数组的元素结构, 当这个数组始终保持有序的时候, 我们每次出队可以迅速找到最大元素, 但是每次入队, 为了保持有序的特性, 其实就是插入排序中的一次插入; 反过来, 当这个数组不强制有序的时候, 我们可以直接加到队列末尾, 但是出队的时候需要遍历整个数组找到最大元素. 两种方式不论哪种都存在一个性能瓶颈.

那么我们换个思路, 我们并不是每次都需要队列整体有序, 只需要队列最前面的元素是最大的元素即可, 这正好就是堆的特性, 不论是入队还是出队, 我们都可以给数组建一次堆, 让最大元素始终处在最前的位置, 时间复杂度对比如下:

| 数据结构 | 入队 | 出队 |
| ------- | --- | --- |
| 有序数组 | O(N)    | 1       |
| 无序数组 | 1       | O(N)    |
| 堆      | O(logN) | O(logN) |

可以看出, 当 N 足够大的时候, 使用堆的综合效率是要远高于前两者的.

首先我们来实现几个基础方法:

    // 判断队列是否为空
    PriorityQueue.prototype.isEmpty = function() {
      return 0 === this.data.length;
    };

    // 对比队列中优先级大小
    PriorityQueue.prototype.less = function(i, j) {
      return this.data[i] < this.data[j];
    };

    // 交换队列中元素位置
    PriorityQueue.prototype.swap = function (i, j) {
      let tmp = this.data[i];
      this.data[i] = this.data[j];
      this.data[j] = tmp;
    };

那么首先来完成 `enqueue` 方法, 照旧, 选择最简单的二叉堆来使用, 假设我们把内部数组看做一个二叉树的话, 那么在每次在树的最后插入元素后, 我们需要 `由下至上` 的将新加入的元素根据优先级提升到本应该在的位置:

    PriorityQueue.prototype.enqueue = function (...ele) {
      this.data.push(...ele);
      this.data.unshift(null); // 加入头元素, 方便从1开始遍历二叉堆
      for (let i = parseInt((this.data.length - 1) / 2); i >= 1; i--) {
        if (this.less(i, i * 2)) this.swap(i, i * 2);
        if (this.data[i * 2 + 1] !== undefined && this.less(i, i * 2 + 1)) this.swap(i, i * 2 + 1);
      }
      this.data.shift();
    };

然后是 `dequeue` 方法, 最简单的做法就是拿出数组第一个元素之后, 重新建堆, 那有没有更快的方法呢?《算法(第四版)》中提供了这样一个思路:

在最大元素出队之后, 如果我们从第二个元素开始建堆, 那么整个二叉树的形状都发生改变了, 可能实际的交换次数会很多, 那么我们换一个方式, 把整个队列的最后一个元素放到二叉堆根的位置, 再 `由上至下` 地将父元素和子元素对比, 如果比子元素优先级小, 则下沉父元素, 再用同样方式处理替换过的子元素, 直到堆根下沉到应该在的位置, 代码如下:

    PriorityQueue.prototype.dequeue = function () {
      if (this.isEmpty()) return undefined;

      const result = this.data.shift();
      this.data = [null, this.data[this.data.length - 1], ...this.data.slice(0, this.data.length - 1)];
      for (let i = 1; i <= parseInt((this.data.length - 1) / 2);) {
        let j = 2 * i;
        if (j < this.data.length - 1 && this.less(j, j + 1)) j++;
        if (!this.less(i, j)) break;
        this.swap(i, j);
        i = j;
      }
      this.data.shift();
      return result;
    };


测试代码:

    const pq = new PriorityQueue();
    pq.enqueue(1, 13, 3);
    pq.enqueue(31, 3);
    pq.enqueue(22);
    console.log(pq.dequeue()); // 31
    console.log(pq.dequeue()); // 22
    console.log(pq.dequeue()); // 13
    console.log(pq.dequeue()); // 3
    console.log(pq.dequeue()); // 3
    console.log(pq.dequeue()); // 1

[完整代码 on GitHub Gist](https://gist.github.com/MrHuxu/06a673b093dd02c621d1d36f38bc825b)

### refs:

- [《算法(第四版)》](https://book.douban.com/subject/19952400/)
- [优先队列](https://zh.wikipedia.org/zh-hans/%E5%84%AA%E5%85%88%E4%BD%87%E5%88%97)
