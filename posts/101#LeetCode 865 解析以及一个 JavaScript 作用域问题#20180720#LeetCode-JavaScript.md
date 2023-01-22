[代码及测试 case](https://github.com/MrHuxu/leetcode/tree/master/problems/865_smallest-subtree-with-all-the-deepest-nodes)

这道题应该算是 medium 难度里的一道水题了, 就是求一棵二叉树中包含所有最深节点的最小子树

![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/01/sketch1.png)

如上图, 7 和 4 就是这棵树里的最深结点, 所以返回 [2, 7, 4], 这样的一个子树就包含了所有的最深节点, 其实这道题换一个说法就是求所有最深节点的最深公共父节点.

既然是二叉树, 那么果断还是递归搞起, 对于一个给定的节点, 从参数可以拿到当前深度, 然后向下遍历:

1. 如果左子节点为空, 那么左子树最大深度为当前节点深度, 不为空就遍历左子节点;
2. 如果右子节点为空, 那么右子树最大深度为当前节点深度, 不为空就遍历右子节点;
3. 如果两个子树最大深度相等, 那么返回当前节点, 否则就返回深度更大的子树.

算法想明白了代码就好说了, 递归搞定, 两次 ac:

    /**
    * Problem: https://leetcode.com/problems/smalle  st-subtree-with-all-the-deepest-nodes/description/
    */

    const smallestSubtreeWithAllTheDeepestNodes = root => {
      const traverse = (node, depth) => {
        let subtreeL, subtreeR;
        let depthL = depth;
        let depthR = depth;

        if (node.left) [subtreeL, depthL] = traverse(node.left, depth + 1);
        if (node.right) [subtreeR, depthR] = traverse(node.right, depth + 1);

        if (depthL === depthR)
          return [node, depthL];
        else
          return depthL > depthR ? [subtreeL, depthL] : [subtreeR, depthR];
      };

      return traverse(root, 0)[0];
    };

---

看上去明明是很简单的题, 为什么我是两次才 ac 的呢, 这就是分割线之后要说的内容.

第一次提交的代码是这样的:

    const smallestSubtreeWithAllTheDeepestNodes = root => {
      const traverse = (node, depth) => {
        let subtreeL, subtreeR;
        let depthL = depthR = depth;

        if (node.left) [subtreeL, depthL] = traverse(node.left, depth + 1);
        if (node.right) [subtreeR, depthR] = traverse(node.right, depth + 1);

        if (depthL === depthR)
          return [node, depthL];
        else
          return depthL > depthR ? [subtreeL, depthL] : [subtreeR, depthR];
      };

      return traverse(root, 0)[0];
    };

但是提交并不能 ac, 然后看了一下 `depthR` 这个变量的值似乎不太对, 想了想, 原来又是中了 JS 作用域的圈套.

在 JS 里, 赋值操作是一个从右到左的过程, 那么下面这么一个赋值语句:

    let depthL = depthR = depth;

会被拆成:

    depthR = depth;
    let depthL = depthR;

这下问题就很明显了, `depthR` 成了一个全局变量, 值肯定会有问题. 所以把这个赋值语句拆开, 轻松 ac. 这一点上, Go 做的就比较鸡贼, 直接就禁止 `var a = b = c` 这种语句了.