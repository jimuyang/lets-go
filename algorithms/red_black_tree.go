package main

import "fmt"

const (
	// RED 红色
	RED bool = true
	// BLACK 黑色
	BLACK bool = false
)

// RBTree 红黑树
// 1. 红黑原则: 每个节点是红色或者黑色
// 2. 黑根原则: 根节点(root)和叶子(LEAF)为黑色
// 3. 连红原则: 如果一个节点是红色，则它的子节点都是黑色
// 4. 黑高原则: 对每个节点，从该节点到其所有后代叶节点的简单路径上，包含相同数目的黑色节点
type RBTree struct {
	root *RBNode
	LEAF *RBNode
}

// Init 初始化一颗红黑树
func (tree *RBTree) Init() *RBTree {
	tree.LEAF = &RBNode{0, BLACK, nil, nil, nil}
	tree.root = tree.LEAF
	return tree
}

// NewTree 得到一棵初始好的红黑树
func NewTree() *RBTree {
	return new(RBTree).Init()
}

// Check 检查红黑树是否违反原则
func (tree *RBTree) Check() error {
	if tree.root.color == RED {
		return fmt.Errorf("违反黑根原则")
	}
	_, err := tree.check(tree.root)
	return err
}

// 返回以node为根的子树的黑高
func (tree *RBTree) check(node *RBNode) (int, error) {
	if node == tree.LEAF {
		return 1, nil
	}
	if node.color == RED {
		if node.left.color == RED || node.right.color == RED {
			return 0, fmt.Errorf("违反连红原则")
		}
	}
	// 左右子树的黑高
	lc, err := tree.check(node.left)
	if err != nil {
		return 0, err
	}
	rc, err := tree.check(node.right)
	if err != nil {
		return 0, err
	}
	if lc != rc {
		return 0, fmt.Errorf("违反黑高原则")
	}
	// 自己的黑高
	bh := lc
	if node.color == BLACK {
		bh++
	}
	return bh, nil
}

// Find 从红黑树中找到一个值
func (tree *RBTree) Find(value int64) *RBNode {
	node := tree.root
	for node != tree.LEAF && node.value != value {
		if value < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

// Min 红黑树的最小值
func (tree *RBTree) Min() (int64, error) {
	minNode := tree.minNode(tree.root)
	if minNode == tree.LEAF {
		return 0, fmt.Errorf("空树")
	}
	return minNode.value, nil
}

func (tree *RBTree) minNode(start *RBNode) *RBNode {
	if start == tree.LEAF {
		return tree.LEAF
	}
	for start.left != tree.LEAF {
		start = start.left
	}
	return start
}

// Max 红黑树的最大值
func (tree *RBTree) Max() (int64, error) {
	maxNode := tree.maxNode(tree.root)
	if maxNode == tree.LEAF {
		return 0, fmt.Errorf("空树")
	}
	return maxNode.value, nil
}

func (tree *RBTree) maxNode(start *RBNode) *RBNode {
	if start == tree.LEAF {
		return tree.LEAF
	}
	for start.right != tree.LEAF {
		start = start.right
	}
	return start
}

// Insert 红黑树中插入一个新值
func (tree *RBTree) Insert(value int64) {
	node := tree.root
	var parent *RBNode = tree.LEAF
	// 找到合适的插入位置
	for node != tree.LEAF {
		parent = node
		if value < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	// 放入一个红色节点
	newNode := &RBNode{value, RED, tree.LEAF, tree.LEAF, parent}
	if parent == tree.LEAF {
		// 空树
		tree.root = newNode
	} else if value < parent.value {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	// 调整红黑树 使之依然保持红黑
	tree.insertFixup(newNode)
}

// 旋转和变色 使之保持红黑。其中的关键是一直保持黑高原则不被打破
// 被打破的只有：
// 1. 黑根原则 => 根节点变色即可
// 2. 连红原则 此时分情况提供解决方案：
// 如果父节点和叔节点都是红色 => 双变色：父节点和叔节点都变黑 爷节点变红
// 如果父节点红、叔节点黑    => 统一战线干掉爷节点：保证我、父、爷成一条直线，然后旋转挤下爷节点
func (tree *RBTree) insertFixup(newNode *RBNode) {
	redNode := newNode
	for redNode.parent.color == RED {
		parent := redNode.parent
		uncle, grandPa := redNode.getUncle(), redNode.getGrandPa()
		if uncle.color == RED {
			// 如果叔节点是红色 => 双变色
			parent.color = BLACK
			uncle.color = BLACK
			grandPa.color = RED
			redNode = grandPa
		} else {
			// 如果叔节点是黑色 => 统一战线干掉爷节点
			if redNode == parent.left && parent == grandPa.right {
				parent, _ = tree.rightRotate(parent)
			} else if redNode == parent.right && parent == grandPa.left {
				parent, _ = tree.leftRotate(parent)
			}
			parent.color = BLACK
			grandPa.color = RED
			if parent == grandPa.left {
				tree.rightRotate(grandPa)
			} else {
				tree.leftRotate(grandPa)
			}
			break
		}
	}
	tree.root.color = BLACK
}

// 旋转
func (tree *RBTree) rotate(node *RBNode, left bool) (*RBNode, error) {
	if left {
		return tree.leftRotate(node)
	}
	return tree.rightRotate(node)
}

// 左旋 将node的右孩子旋转到自己的位置
func (tree *RBTree) leftRotate(node *RBNode) (*RBNode, error) {
	rc := node.right
	if rc == tree.LEAF {
		return nil, fmt.Errorf("右孩子不存在, 无法左旋")
	}
	// 右孩子的左孩子“过继”给自己当右孩子
	node.right = rc.left
	// 这里不希望随便改变LEAF的parent
	if rc.left != tree.LEAF {
		rc.left.parent = node
	}
	// 右孩子上位
	rc.parent = node.parent
	if node.parent == tree.LEAF {
		// 本来是根节点
		tree.root = rc
	} else if node.parent.left == node {
		node.parent.left = rc
	} else {
		node.parent.right = rc
	}
	// 自己成为左孩子
	rc.left = node
	node.parent = rc
	return rc, nil
}

// 右旋 将node的左孩子旋转到自己的位置
func (tree *RBTree) rightRotate(node *RBNode) (*RBNode, error) {
	lc := node.left
	if lc == tree.LEAF {
		return nil, fmt.Errorf("左孩子不存在，无法右旋")
	}
	// 左孩子的右孩子过继给自己
	node.left = lc.right
	// 这里不希望随便改变LEAF的parent
	if lc.right != tree.LEAF {
		lc.right.parent = node
	}
	// 左孩子上位
	lc.parent = node.parent
	if node.parent == tree.LEAF {
		tree.root = lc
	} else if node.parent.left == node {
		node.parent.left = lc
	} else {
		node.parent.right = lc
	}
	// 自己成为右孩子
	lc.right = node
	node.parent = lc
	return lc, nil
}

// Delete 从红黑树中删除一个值
func (tree *RBTree) Delete(value int64) error {
	node := tree.Find(value)
	if node == tree.LEAF {
		return fmt.Errorf("树中没有该值 无法删除")
	}
	return tree.deleteNode(node)
}

// 从红黑树中删除一个节点
func (tree *RBTree) deleteNode(node *RBNode) error {
	// 实际上移的节点
	var move *RBNode
	// 实际少了的颜色
	color := node.color
	if node.left == tree.LEAF {
		// 右子树上移
		move = node.right
		tree.transplant(node, move)
	} else if node.right == tree.LEAF {
		// 左子树上移
		move = node.left
		tree.transplant(node, move)
	} else {
		// 目标节点有2个子节点，此时找到目标节点右子树的最小值（即后继节点）
		// 将后继节点赋值给目标节点后，此时实际要删除的是后继节点
		// 后继节点的左子树一定为LEAF 将后继节点的右子树上移
		successor := tree.minNode(node)
		node.value = successor.value
		color = successor.color
		move = successor.right
		tree.transplant(successor, move)
	}
	if color == BLACK {
		// 删除了一个黑色节点 此时需要fixup
		// 黑高原则还是不容打破 此时我们认为上移的节点move拥有额外的一个黑色
		// 此时move的颜色为红黑(RED)或黑黑(BLACK) 此时违反了 红黑原则
		// move此时可能为LEAF 但在transplant中也赋值了parent
		tree.deleteFixup(move)
	}
	return nil
}

// fixup的目的是将这额外的黑色沿着树上移 直到：
// 1. node为一个红黑节点 => 直接染成黑色即可
// 2. node指向了根节点 => 将额外的黑色移除
// 3. 适当的旋转和着色从而退出循环
func (tree *RBTree) deleteFixup(moveNode *RBNode) {
	// node节点自带一层黑色 这里的node可能为LEAF 但带了parent
	node := moveNode
	// 是黑黑节点且不是根节点 此时需要处理额外的黑色 此时兄弟节点一定不为LEAF
	for node.color == BLACK && node != tree.root {
		// 分类讨论: 下面只看我是左孩子 兄弟是右孩子的情况
		// 1. 如果兄弟节点是红色的 左旋让兄弟是黑的
		// 2. 如果兄弟节点是黑的
		//    2.1 兄弟的子女都是黑的 就可以和兄弟一起脱黑 额外的黑转移给父节点
		//    2.2 兄弟的右孩子是红的 左旋即可 （拉父节点下水 让父节点承担额外的黑 解决问题
		//    2.3 兄弟的左孩子是红的 兄弟右旋让右孩子变红 回到2.2
		brother := node.getBrother()
		amLeft := node == node.parent.left
		if brother.color == RED {
			// 此时parent是黑的 brother的2个子节点都不为LEAF且都是黑的
			node.parent.color = RED
			brother.color = BLACK
			tree.rotate(node.parent, amLeft)
			// 此时兄弟已经换了 新兄弟一定不为LEAF 且一定是黑色（看图参考黑高原则）
			brother = node.getBrother()
		}
		// 此时兄弟是黑的
		if brother.left.color == BLACK && brother.right.color == BLACK {
			// 可以和兄弟一起脱黑
			brother.color = RED
			// 额外的黑色转移到父节点
			node = node.parent
		} else {
			// 不能一起脱黑 那就旋转变色 解决问题
			if amLeft {
				if brother.left.color == RED {
					brother.left.color = BLACK
					brother.color = RED
					tree.rightRotate(brother)
				}
				brother = node.parent.right
				// 此时brother的右孩子一定是红色 整体左旋
				brother.color = node.parent.color
				node.parent.color = BLACK // 背额外的黑
				brother.right.color = BLACK
				tree.leftRotate(node.parent)
			} else {
				if brother.right.color == RED {
					brother.right.color = BLACK
					brother.color = RED
					tree.leftRotate(brother)
				}
				brother = node.parent.left
				// 此时brother的左孩子一定是红色 整体右旋
				brother.color = node.parent.color
				node.parent.color = BLACK // 背额外的黑
				brother.left.color = BLACK
				tree.rightRotate(node.parent)
			}
			break
		}
	}
	// 是红黑节点或者是根节点 简单的染成黑色即可
	node.color = BLACK
}

// 移植 替代原来的节点 只有指针的赋值
func (tree *RBTree) transplant(deleted *RBNode, move *RBNode) {
	if deleted.parent == tree.LEAF {
		tree.root = move
	} else if deleted == deleted.parent.left {
		deleted.parent.left = move
	} else {
		deleted.parent.right = move
	}
	// 这里move可能是LEAF 也需要修改parent来记录自己的位置
	move.parent = deleted.parent
}

// RBNode 红黑树的节点
type RBNode struct {
	value               int64
	color               bool
	left, right, parent *RBNode
}

func (node *RBNode) getGrandPa() *RBNode {
	return node.parent.parent
}

func (node *RBNode) getUncle() *RBNode {
	grandPa := node.getGrandPa()
	if grandPa.left == node.parent {
		return grandPa.right
	}
	return grandPa.left
}

func (node *RBNode) getBrother() *RBNode {
	if node == node.parent.left {
		return node.parent.right
	}
	return node.parent.left
}
