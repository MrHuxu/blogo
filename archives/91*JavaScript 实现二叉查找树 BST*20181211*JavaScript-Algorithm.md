学过数据结构的都知道, 二叉树是一种非常重要的数据结构, 这种树是有两个子节点分别被成为左右子节点, 一般一个用来存储整数的二叉树节点的定义如下:

    /**  
    * @param {number} val  
    * @return {}  
    */  
    const TreeNode = function (val) {
      this.data = val;
      this.left = this.right = null;
    };

而二叉查找树是一种特殊的二叉树, 它的主要特征就是: 对于一棵没有重复元素的二叉查找树, 任何一个不为空的节点 `node` 一定满足 `node.left.data < node.data < node.right.data` 这个不等式.

那么根据这个二叉查找树最基本的特性, 我们很容易写出查找函数 `searchNode`:

1. 如果查找的值小于当前节点, 那么在当前节点的左子结点中继续查找, 如果大于, 则在右子节点中查找
2. 如果等于当前节点的值, 则查找成功, 第一个返回值是目标节点, 第二个是目标节点的父节点
3. 如果当前节点为空, 则查找失败, 第一个返回值是 null, 第二个是当前节点的父节点

至于为什么需要在这里返回父节点呢, 看了下面的 `insertNode` 就能知道了, 代码如下:


    /**  
     * @param {TreeNode} node  
     * @param {number} val  
     * @param {TreeNode} parentNode  
     * @return {TreeNode[]} node, parentNode  
     */  
    const searchNode = (node, val, parentNode) => {
      if (!node) return [null, parentNode];
      else {
        if (node.data === val) return [node, parentNode];
        else if (node.data > val) return searchNode(node.left, val, node);
        else if (node.data < val) return searchNode(node.right, val, node);
      }
    };

那么接下来再来实现向二叉查找树中插入元素的 `insertNode` 方法, 这个方法正好就用到了上面 `searchNode` 结果中的父节点:

1. 如果当前树是空的话, 新建一个节点返回
2. 如果当前树不为空, 那么通过 `searchNode` 找到目标节点和父节点, 如果目标节点为空, 就根据插入元素的大小, 响应创建父节点的左子结点或右子节点

代码如下:


    /**  
    * @param {TreeNode} root  
    * @param {number} val  
    * @return {TreeNode} root  
    */  
    const insertNode = (root, val) => {
      if (!root) return new TreeNode(val);

      const [node, parentNode] = searchNode(root, val, null);
      if (!node) {
        const temp = new TreeNode(val);
        if (val < parentNode.data) parentNode.left = temp;
        else if (val > parentNode.data) parentNode.right = temp;
      }
      return root;
    }

接下来就是难度比较大的删除操作, 先描述算法:

1. 当被删除节点只有左子结点时, 直接把其父节点往下的指针指向左子结点即可, 只有右子节点同理
2. 当被删除节点有两个子节点时, 需要找到其后继节点, 删除节点后将其后继节点移动到被删除节点的位置

算法说起来简单, 但是实现起来就有一定难度了, 我们解释一下代码的逻辑:

1. 如果被删除元素小于当前节点的值, 这个元素应该存在于这个节点的左子树中, 当前节点的左子结点应该是左子树删除元素之后的结果, 大于则应该用相同方式处理右子树
2. 如果等于, 那么根据前面描述的算法, 有下面三中情况
    - 如果左子节点为空, 直接返回当前节点的右子节点
    - 如果右子节点为空, 直接返回当前节点的左子结点
    - 如果左右子节点都不为空, 当前节点的后继节点应该右子树中的最左节点, 将节点值替换完成之后, 再递归删除右子树中后继节点, 并且返回替换完成的节点

代码如下:

    /**  
    * @param {TreeNode} node  
    * @param {number} val  
    * @return {}  
    */  
    const deleteNode = (node, val) => {
      if (!node) {
        return null;
      } else if (val < node.data) {
        node.left = deleteNode(node.left, val);
        return node;
      } else if (val > node.data) {
        node.right = deleteNode(node.right, val);
        return node;
      } else {
        if (!node.left) {
          return node.right;
        } else if (!node.right) {
          return node.left;
        } else {
          let next = node.right;
          while (next.left) next = next.left;

          node.data = next.data;
          node.right = deleteNode(node.right, next.data);
          return node;
        }
      }
    }

以上就是一个普通二叉查找树 BST 的三个关键算法了, 二叉查找树还有一个性质, 我们知道中序遍历二叉树的时候, 顺序是左子结点 -> 当前节点 -> 右子节点, 那么二叉查找树的中序遍历结果一定是一个从小到大的有序数列. 下面就是实际应用时的代码:

    /**  
    * @param {TreeNode} node  
    * @return {}  
    */  
    const inorderTraverse = node => {
      if (node) {
        inorderTraverse(node.left);
        console.log(node.data);
        inorderTraverse(node.right);
      }
    };

    let root = insertNode(null, 3);
    root = insertNode(root, 5);
    root = insertNode(root, 4);
    root = insertNode(root, 7);
    root = insertNode(root, 6);
    inorderTraverse(root);   // 3 \n 4 \n 5 \n 6 \n 7

    deleteNode(root, 5);
    console.log(JSON.stringify(root));
    /*  
     *          3  
     *            \  
     *              6  
     *             / \  
     *            4   7  
     */  
    inorderTraverse(root);   // 3 \n 4 \n 6 \n 7