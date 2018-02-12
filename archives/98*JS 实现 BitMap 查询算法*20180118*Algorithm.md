[完整代码 on GitHub Gist](https://gist.github.com/MrHuxu/343fa6b1c35500d8be1c4fc066caefc6)


## 概念

BitMap 主要用来解决大量数据存储的问题, 在数据为亿级的情况下, 如果直接散列的话, 将会造成很大的内存浪费, 那么我们可以申请一块儿内存, 直接将数字 n 存在从低到高第 n + 1 上.

> 0 0 1 0 1 1 0 1  
_ _ 5 _ 3 2 _ 0

比如一个数组 `[0, 2, 3, 5]` 可以用上面的方式, 存到一个字节里. 如果需要存储 10 亿的数字, 使用 int 存储的话, 每个 int 正数是 32 位, 总共需要 1000000000 / 32 = 31250000 个正数, 只需要 30MB 左右的空间就可以满足.

---

## 实现

现在我们用 JS 来实现一个 BitMap, 下面是构造函数:

    const BitMap = function () {
      this.data = [];
    };

然后是两个基础函数, 用来计算一个数应该存在 `data` 数组里的索引, 以及在整数里的具体位置.

    BitMap.prototype.getIdx = num => parseInt(num / 32);
    BitMap.prototype.getPos = num => num % 32;

然后接下来就是添加操作, 就是找到具体的正数用 `|=` 操作符将相应位数置 1 即可:

    BitMap.prototype.add = function (num) {
      const index = this.getIdx(num);
      const pos = this.getPos(num);

      if (this.data[index] === undefined) this.data[index] = 0;
      this.data[index] |= Math.pow(2, pos);
    };

判断是否存在也很简单, 找到位置做按位与操作就可以得到结果:

    BitMap.prototype.exist = function (num) {
      const index = this.getIdx(num);
      const pos = this.getPos(num);
          
      return !!(this.data[index] && (this.data[index] & Math.pow(2, pos)));
    };

---

## 应用

1. 数字查找  
    其实就是上面的例子, 实现大数字的快速查找.

2. 状态存储  
    我们可以知道, n 个二进制位其实可以表示 2^n 个数字, 也就是可以看作是 2^n 个状态, 假设一个 id 对应 3 种状态, 那么我们可以用两位来表示一个 id, 1&2 位表示第一个 id = 0 的状态, 3&4 位可以表示 id = 1 的状态, 依次类推, 实现方法和上面的例子类似.