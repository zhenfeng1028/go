package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Node 代表跳表中的一个节点
type Node struct {
	Value int     // 节点存储的值
	Level int     // 节点的高度/层级
	Next  []*Node // Next[i] 指向第i层的下一个节点
	Prev  *Node   // 指向上一个节点（用于有序遍历）
}

// SkipList 跳表的主结构体
type SkipList struct {
	head     *Node   // 虚拟头节点
	tail     *Node   // 虚拟尾节点
	level    int     // 当前跳表的最高层级
	size     int     // 跳表中元素的个数
	maxLevel int     // 跳表允许的最高层级
	p        float64 // 概率参数，用于生成随机层级（通常为0.5）
}

// NewSkipList 创建一个新的跳表
func NewSkipList() *SkipList {
	maxLevel := 32
	sl := &SkipList{
		head:     &Node{Value: math.MinInt64, Level: maxLevel, Next: make([]*Node, maxLevel)},
		tail:     &Node{Value: math.MaxInt64, Level: maxLevel, Next: make([]*Node, maxLevel)},
		level:    0,
		size:     0,
		maxLevel: maxLevel,
		p:        0.5,
	}

	// 初始化head的Next指针都指向tail
	for i := 0; i < maxLevel; i++ {
		sl.head.Next[i] = sl.tail
	}

	return sl
}

// randomLevel 生成随机的层级
func (sl *SkipList) randomLevel() int {
	level := 1
	for rand.Float64() < sl.p && level < sl.maxLevel {
		level++
	}
	return level
}

// Search 查找单个元素是否存在
func (sl *SkipList) Search(value int) bool {
	current := sl.head

	// 从最高层开始向下搜索
	for i := sl.level - 1; i >= 0; i-- {
		for current.Next[i] != sl.tail && current.Next[i].Value < value {
			current = current.Next[i]
		}
	}

	// 移动到第0层，检查下一个节点
	current = current.Next[0]
	return current != sl.tail && current.Value == value
}

// Insert 向跳表中插入元素
func (sl *SkipList) Insert(value int) {
	// 如果元素已存在，不插入
	if sl.Search(value) {
		return
	}

	// 找到所有需要更新的节点
	update := make([]*Node, sl.maxLevel)
	current := sl.head

	// 从最高层开始向下搜索
	for i := sl.level - 1; i >= 0; i-- {
		for current.Next[i] != sl.tail && current.Next[i].Value < value {
			current = current.Next[i]
		}
		update[i] = current
	}

	// 生成新节点的随机层级
	newLevel := sl.randomLevel()

	// 如果新节点的层级大于当前跳表的层级，需要更新head的指针
	if newLevel > sl.level {
		for i := sl.level; i < newLevel; i++ {
			update[i] = sl.head
		}
		sl.level = newLevel
	}

	// 创建新节点
	newNode := &Node{
		Value: value,
		Level: newLevel,
		Next:  make([]*Node, newLevel),
	}

	// 在第0层找到前一个节点用于设置Prev指针
	prevNode := update[0]
	newNode.Prev = prevNode
	if prevNode.Next[0] != sl.tail {
		prevNode.Next[0].Prev = newNode
	}

	// 在各层插入新节点
	for i := 0; i < newLevel; i++ {
		newNode.Next[i] = update[i].Next[i]
		update[i].Next[i] = newNode
	}

	sl.size++
}

// Delete 从跳表中删除元素
func (sl *SkipList) Delete(value int) bool {
	// 首先检查元素是否存在
	if !sl.Search(value) {
		return false
	}

	// 找到所有需要更新的节点
	update := make([]*Node, sl.maxLevel)
	current := sl.head

	// 从最高层开始向下搜索
	for i := sl.level - 1; i >= 0; i-- {
		for current.Next[i] != sl.tail && current.Next[i].Value < value {
			current = current.Next[i]
		}
		update[i] = current
	}

	// 获取要删除的节点
	nodeToDelete := update[0].Next[0]

	// 删除节点：更新所有层的指针
	for i := 0; i < nodeToDelete.Level; i++ {
		update[i].Next[i] = nodeToDelete.Next[i]
	}

	// 更新Prev指针
	if nodeToDelete.Next[0] != sl.tail {
		nodeToDelete.Next[0].Prev = nodeToDelete.Prev
	}

	// 如果删除节点后，最高层没有节点了，降低层级
	for sl.level > 0 && sl.head.Next[sl.level-1] == sl.tail {
		sl.level--
	}

	sl.size--
	return true
}

// GetAll 按顺序返回所有元素
func (sl *SkipList) GetAll() []int {
	result := make([]int, 0, sl.size)

	// 从第0层的第一个节点开始遍历
	current := sl.head.Next[0]
	for current != sl.tail {
		result = append(result, current.Value)
		current = current.Next[0]
	}

	return result
}

// RangeSearch 查找区间 [min, max] 内的所有元素
func (sl *SkipList) RangeSearch(min, max int) []int {
	result := make([]int, 0)

	// 找到第一个大于等于min的节点
	current := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for current.Next[i] != sl.tail && current.Next[i].Value < min {
			current = current.Next[i]
		}
	}

	// 移动到第0层
	current = current.Next[0]

	// 从该节点开始收集所有在 [min, max] 范围内的元素
	for current != sl.tail && current.Value <= max {
		result = append(result, current.Value)
		current = current.Next[0]
	}

	return result
}

// Size 返回跳表中元素的个数
func (sl *SkipList) Size() int {
	return sl.size
}

// IsEmpty 检查跳表是否为空
func (sl *SkipList) IsEmpty() bool {
	return sl.size == 0
}

// Display 打印跳表的结构（用于调试）
func (sl *SkipList) Display() {
	fmt.Println("=== Skip List Structure ===")
	for level := sl.level - 1; level >= 0; level-- {
		fmt.Printf("Level %d: ", level)
		current := sl.head.Next[level]
		for current != sl.tail {
			fmt.Printf("%d -> ", current.Value)
			current = current.Next[level]
		}
		fmt.Println("nil")
	}
	fmt.Println()
}

// Min 返回最小元素
func (sl *SkipList) Min() (int, bool) {
	if sl.IsEmpty() {
		return 0, false
	}
	return sl.head.Next[0].Value, true
}

// Max 返回最大元素
func (sl *SkipList) Max() (int, bool) {
	if sl.IsEmpty() {
		return 0, false
	}

	current := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for current.Next[i] != sl.tail {
			current = current.Next[i]
		}
	}

	return current.Value, true
}

// CountRange 统计区间内元素的个数
func (sl *SkipList) CountRange(min, max int) int {
	return len(sl.RangeSearch(min, max))
}

func main() {
	fmt.Println("\n========== 跳表演示 ==========")
	sl := NewSkipList()

	// 插入数据
	data := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 45}
	fmt.Println("1. 插入元素:", data)
	for _, v := range data {
		sl.Insert(v)
	}

	// 显示跳表结构
	sl.Display()

	// 有序输出所有元素
	fmt.Println("2. 按顺序输出所有元素:")
	fmt.Println("   ", sl.GetAll())
	fmt.Printf("   元素个数: %d\n\n", sl.Size())

	// 查找单个元素
	fmt.Println("3. 查找单个元素:")
	fmt.Printf("   查找 25: %v\n", sl.Search(25))
	fmt.Printf("   查找 99: %v\n\n", sl.Search(99))

	// 范围查询
	fmt.Println("4. 范围查询:")
	fmt.Println("   [30, 80] 范围内的元素:", sl.RangeSearch(30, 80))
	fmt.Println("   [10, 50] 范围内的元素:", sl.RangeSearch(10, 50))
	fmt.Println("   [0, 100] 范围内的元素:", sl.RangeSearch(0, 100))
	fmt.Printf("   [30, 80] 范围内元素个数: %d\n\n", sl.CountRange(30, 80))

	// 最小最大值
	fmt.Println("5. 最小/最大值:")
	min, _ := sl.Min()
	max, _ := sl.Max()
	fmt.Printf("   最小值: %d, 最大值: %d\n\n", min, max)

	// 删除元素
	fmt.Println("6. 删除元素 [25, 50, 88]:")
	fmt.Printf("   删除 25: %v\n", sl.Delete(25))
	fmt.Printf("   删除 50: %v\n", sl.Delete(50))
	fmt.Printf("   删除 88: %v\n", sl.Delete(88))
	fmt.Printf("   删除不存在的元素 99: %v\n\n", sl.Delete(99))

	fmt.Println("7. 删除后的所有元素:")
	fmt.Println("   ", sl.GetAll())
	fmt.Printf("   元素个数: %d\n\n", sl.Size())

	sl.Display()

	// 进行更多的插入和删除操作
	fmt.Println("8. 进行更多操作:")
	fmt.Println("   插入 [100, 5, 40, 55, 20]")
	for _, v := range []int{100, 5, 40, 55, 20} {
		sl.Insert(v)
	}

	fmt.Println("   所有元素:", sl.GetAll())
	fmt.Println("   [40, 90] 范围内的元素:", sl.RangeSearch(40, 90))
	fmt.Printf("   [40, 90] 范围内元素个数: %d\n", sl.CountRange(40, 90))
}
