先解释题目, 给定一棵没有重复元素的二叉树 `root` 和目标结点值 `k`, 找出这棵树中距离k最近的叶子结点并返回结点值, 如果有多个叶子结点距离相同且最小, 返回任意一个即可.

这个问题的难点是, 最短距离的叶子结点可能并不是目标结点的子节点, 比如下面这个例子:

>
    Input:
    root = [1,2,3,4,null,null,null,5,null,6], k = 2
    Diagram of binary tree:
                 1
                / \
               2   3
              /
             4
            /
           5
          /
         6

    Output: 3

在这个例子里, `2` 的子节点中有叶子结点 `6`, 但是6和2的距离是3, 而另一个叶子结点 `3` 通过 `1` 到达2的距离是2, 所以返回的结果应该是3.

通过上面这个例子, 可以发现在这道题里存在两种情况:
1. 当目标结点是叶子结点的祖先结点时, 两者距离是深度之差;
2. 当目标结点不是叶子结点的祖先结点时, 有一个回溯过程, 也就是要找到这个叶子结点和目标结点的最低公共父节点, 然后才能计算出和目标结点的距离.

那我决定使用暴力一点的解法, 即首先便利一遍树并且储存下回溯时需要的信息, 再对扩展之后的二叉树进行处理.

首先我们扩展一下树结点, 加上两个字段

1. `targetInChildren`: 这个字段表示目标结点是否存在于当前结点的子节点中;
2. `depth`: 这个字段表示当前结点在二叉树中的深度, 根节点为0, 往下依次递增

当前使用的二叉树结点可以表示为:

    const NewTreeNode = function(val) {
      this.val = val;
      this.left = this.right = null;
      this.targetInChildren = false;
      this.depth = 0;
    };

同时我们引入一个外部变量 `targetDepth` 来保存目标结点的深度, 这里可以用一个递归函数来遍历这棵树以得到这两个字段的值:

1. 直接更新 `depth` 字段;
2. 更新左子结点和右子节点的信息;
3. 判断当前结点值是否等于目标值:

    - 如果相等, 更新 `targetDepth` 的值, 并将 `targetInChildren` 置为 `true`;
    - 如果不相等, 则 `targetInChildren` 字段值由子节点的该字段值决定.

下面是这个函数的实现:

    let targetDepth;
    const expandTree = (node, depth) => {
      if (node) {
        node.depth = depth;
        node.left = expandTree(node.left, depth + 1);
        node.right = expandTree(node.right, depth + 1);
        if (node.val === k) {
          targetDepth = depth;
          node.targetInChildren = true;
        } else {
          node.targetInChildren = (node.left && node.left.targetInChildren) || (node.right && node.right.targetInChildren);
        }
        return node; 
      } else return null;
    };

得到我们需要的信息后, 我们就可以用另一个函数来递归求解了, 这个函数有三个参数:

1. `node`: 当前遍历的树节点;
2. `targetInParent`: 目标结点是否是当前结点的祖先结点;
3. `plusDepth`: 这个参数表示当目标结点不是当前结点的祖先结点时, 计算距离所需要的参数, 这个参数的计算方法将在下面说明.

以下是一次遍历过程:

1. 如果是叶子结点:
    - 如果目标结点是祖先结点或者等于当前结点, 直接用当前结点深度减去 `targetDepth` 获得距离, 根据距离更新结果;
    - 如果不是, 则用当前节点深度和参数 `plusDepth` 相加获得距离, 据距离更新结果.
2. 如果不是叶子结点, 那么计算新的参数递归遍历左右子节点:
    - `newTargetInParent`: 直接用当前参数或上 `node.val === k`;
    - `newPlusDepth`: 这个值只有在目标结点不是当前结点祖先结点, 但是是当前结点子节点的时候需要更新, 首先目标结点和当前结点的距离是 `targetDepth - node.depth`, 叶子结点和当前结点的距离是 `leafDepth - node.depth`, 那么距离只和应该是 `targetDepth + leafDepth - 2 * node.depth`, 在这个式子里除去 `leafDepth`, 可以得出新的参数计算方式为 `targetDepth - 2 * node.depth`.

代码实现如下:

    let min = Number.MAX_SAFE_INTEGER, result;
    const traverse = (node, targetInParent, plusDepth) => {
      if (node) {
        if (!node.left && !node.right) {
          const subDepth = (targetInParent || node.val === k) ? (node.depth - targetDepth) : node.depth + plusDepth;
          if (subDepth < min) {
            min = subDepth;
            result = node.val;
          }
        } else {
          const newTargetInParent = targetInParent || node.val === k;
          const newPlusDepth = newTargetInParent ? 0 : (node.targetInChildren ? targetDepth - node.depth * 2 : plusDepth);
          traverse(node.left, newTargetInParent, newPlusDepth);
          traverse(node.right, newTargetInParent, newPlusDepth);
        }
      }
    };

这样通过遍历两次这棵树, 我们可以在 O(2n) 的时间复杂度内获得结果.

LeetCode AC 的完整代码如下:

    var findClosestLeaf = function(root, k) {
      let targetDepth;
      const expandTree = (node, depth) => {
        if (node) {
          node.depth = depth;
          node.left = expandTree(node.left, depth + 1);
          node.right = expandTree(node.right, depth + 1);
          if (node.val === k) {
            targetDepth = depth;
            node.targetInChildren = true;
          } else {
            node.targetInChildren = (node.left && node.left.targetInChildren) || (node.right && node.right.targetInChildren);
          }
          return node; 
        } else return null;
      };
      expandTree(root, 0);

      let min = Number.MAX_SAFE_INTEGER, result;
      const traverse = (node, targetInParent, plusDepth) => {
        if (node) {
          if (!node.left && !node.right) {
            const subDepth = (targetInParent || node.val === k) ? (node.depth - targetDepth) : node.depth + plusDepth;
            if (subDepth < min) {
              min = subDepth;
              result = node.val;
            }
          } else {
            const newTargetInParent = targetInParent || node.val === k;
            const newPlusDepth = newTargetInParent ? 0 : (node.targetInChildren ? targetDepth - node.depth * 2 : plusDepth);
            traverse(node.left, newTargetInParent, newPlusDepth);
            traverse(node.right, newTargetInParent, newPlusDepth);
          }
        }
      };
      traverse(root, false, 0);

      return result;
    };