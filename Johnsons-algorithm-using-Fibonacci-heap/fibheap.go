package main

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"math"
)

type HeapInterface interface {
	Num() uint
	Insert(tag interface{}, key float64) error
	Minimum() (interface{}, float64)
	ExtractMin() (interface{}, float64)
	DecreaseKey(tag interface{}, key float64) error
}

// FibHeap представляє собою піраміду Фібоначчі.
type FibHeap struct {
	roots       *list.List             // Список коренів
	index       map[interface{}]*node  // Хеш-таблиця пар Значення - Вузол
	treeDegrees map[uint]*list.Element // Хеш-таблиця пар Степінь кореня - Кореневий вузол
	min         *node                  // Вузол із мінімальним ключем
	num         uint                   // Кількість вузлів у піраміді
}

// Структура вузла піраміди Фібоначчі
type node struct {
	self     *list.Element // Посилання на самого себе у списку дітей батька
	parent   *node         // посилання на батьківський вузол
	children *list.List    // посилання на список дочірніх вузлів
	degree   uint          // кількість дочірніх вузлів
	position uint          // ключ для ноди в хеш-таблиці treeDegrees
	tag      interface{}   // Значення яке ми зберігаємо у піраміді
	key      float64       // Ключ-пріорітет значення
	// Значення, яке вказує, чи були втрати вузлом дочірніх вузлів,
	// починаючи з моменту, коли данний став дочірнім вузлом якогось іншого вузла.
	marked bool
}

// Коструктор піраміди
func NewFibHeap() *FibHeap {
	heap := new(FibHeap)
	heap.roots = list.New()
	heap.index = make(map[interface{}]*node)
	heap.treeDegrees = make(map[uint]*list.Element)
	heap.num = 0
	heap.min = nil

	return heap
}

// Num повертає загальну кількість значень у піраміді.
func (heap *FibHeap) Num() uint {
	return heap.num
}

// Insert вставляє вхідне значення і ключ до піраміди.
// Спроба вставити дублікат значення призведе до повернення помилки.
// Допустимий діапазон ключа (-inf, +inf].
// Спроба вставити значення ключа -inf призведе до помилки.
// Спроба введення значення, що дорівнює nil викличе помилку невірної адреси.
func (heap *FibHeap) Insert(tag interface{}, key float64) error {
	if tag == nil {
		return errors.New("Input tag is nil ")
	}

	return heap.insert(tag, key)
}

// Minimum повертає значення з мінімальним ключем та сам ключ у піраміді
// Порожня піраміда поверне nil та -inf.
func (heap *FibHeap) Minimum() (interface{}, float64) {
	if heap.num == 0 {
		return nil, math.Inf(-1)
	}

	return heap.min.tag, heap.min.key
}

// Minimum повертає значення з мінімальним ключем та сам ключ у піраміді, а потім витягує їх з піраміди.
// Порожня піраміда поверне nil/-inf і нічого не витягне.
func (heap *FibHeap) ExtractMin() (interface{}, float64) {
	if heap.num == 0 {
		return nil, math.Inf(-1)
	}

	min := heap.extractMin()

	return min.tag, min.key
}

// DecreaseKey оновлює ключ заданого значення у піраміді на вхідний ключ.
// Якщо вхідний ключ має більший ключ або ключ -inf, буде повернуто помилку.
// Якщо вхідне значення не існує у піраміді, буде повернуто помилку.
func (heap *FibHeap) DecreaseKey(tag interface{}, key float64) error {
	if tag == nil {
		return errors.New("Input tag is nil ")
	}

	if math.IsInf(key, -1) {
		return errors.New("Negative infinity key is reserved for internal usage ")
	}

	if node, exists := heap.index[tag]; exists {
		return heap.decreaseKey(node, key)
	}

	return errors.New("Value is not found ")
}

// String надає деяку базову налагоджувальну інформацію про піраміду.
// Повертає загальну кількість, розмір коренів, розмір індексу та поточне мінімальне значення піраміди.
// Також повертає топологію дерев за результатами пошуку dfs.
func (heap *FibHeap) String() string {
	var buffer bytes.Buffer

	if heap.num != 0 {
		buffer.WriteString(fmt.Sprintf("Total number: %d, Root Size: %d, Index size: %d,\n", heap.num, heap.roots.Len(), len(heap.index)))
		buffer.WriteString(fmt.Sprintf("Current minimun: key(%f), tag(%v),\n", heap.min.key, heap.min.tag))
		buffer.WriteString(fmt.Sprintf("Heap detail:\n"))
		probeTree(&buffer, heap.roots)
		buffer.WriteString(fmt.Sprintf("\n"))
	} else {
		buffer.WriteString(fmt.Sprintf("Heap is empty.\n"))
	}

	return buffer.String()
}

func probeTree(buffer *bytes.Buffer, tree *list.List) {
	buffer.WriteString(fmt.Sprintf("< "))
	for e := tree.Front(); e != nil; e = e.Next() {
		buffer.WriteString(fmt.Sprintf("%f ", e.Value.(*node).key))
		if e.Value.(*node).children.Len() != 0 {
			probeTree(buffer, e.Value.(*node).children)
		}
	}
	buffer.WriteString(fmt.Sprintf("> "))
}

func (heap *FibHeap) decreaseKey(n *node, key float64) error {
	if key >= n.key {
		return errors.New("New key is not smaller than current key ")
	}

	n.key = key
	if n.parent != nil {
		parent := n.parent
		if n.key < n.parent.key {
			heap.cut(n)
			heap.cascadingCut(parent)
		}
	}

	if n.parent == nil && n.key < heap.min.key {
		heap.min = n
	}

	return nil
}

func (heap *FibHeap) cut(n *node) {
	n.parent.children.Remove(n.self)
	n.parent.degree--
	n.parent = nil
	n.marked = false
	n.self = heap.roots.PushBack(n)
}

func (heap *FibHeap) cascadingCut(n *node) {
	if n.parent != nil {
		if !n.marked {
			n.marked = true
		} else {
			parent := n.parent
			heap.cut(n)
			heap.cascadingCut(parent)
		}
	}
}

func (heap *FibHeap) consolidate() {
	for tree := heap.roots.Front(); tree != nil; tree = tree.Next() {
		heap.treeDegrees[tree.Value.(*node).position] = nil
	}

	for tree := heap.roots.Front(); tree != nil; {
		if heap.treeDegrees[tree.Value.(*node).degree] == nil {
			heap.treeDegrees[tree.Value.(*node).degree] = tree
			tree.Value.(*node).position = tree.Value.(*node).degree
			tree = tree.Next()
			continue
		}

		if heap.treeDegrees[tree.Value.(*node).degree] == tree {
			tree = tree.Next()
			continue
		}

		for heap.treeDegrees[tree.Value.(*node).degree] != nil {
			anotherTree := heap.treeDegrees[tree.Value.(*node).degree]
			heap.treeDegrees[tree.Value.(*node).degree] = nil
			if tree.Value.(*node).key <= anotherTree.Value.(*node).key {
				heap.roots.Remove(anotherTree)
				heap.link(tree.Value.(*node), anotherTree.Value.(*node))
			} else {
				heap.roots.Remove(tree)
				heap.link(anotherTree.Value.(*node), tree.Value.(*node))
				tree = anotherTree
			}
		}
		heap.treeDegrees[tree.Value.(*node).degree] = tree
		tree.Value.(*node).position = tree.Value.(*node).degree
	}

	heap.resetMin()
}

func (heap *FibHeap) insert(tag interface{}, key float64) error {
	if math.IsInf(key, -1) {
		return errors.New("Negative infinity key is reserved for internal usage ")
	}

	if _, exists := heap.index[tag]; exists {
		return errors.New("Duplicate tag is not allowed ")
	}

	node := new(node)
	node.children = list.New()
	node.tag = tag
	node.key = key

	node.self = heap.roots.PushBack(node)
	heap.index[node.tag] = node
	heap.num++

	if heap.min == nil || heap.min.key > node.key {
		heap.min = node
	}

	return nil
}

func (heap *FibHeap) extractMin() *node {
	min := heap.min

	children := heap.min.children
	if children != nil {
		for e := children.Front(); e != nil; e = e.Next() {
			e.Value.(*node).parent = nil
			e.Value.(*node).self = heap.roots.PushBack(e.Value.(*node))
		}
	}

	heap.roots.Remove(heap.min.self)
	heap.treeDegrees[min.position] = nil
	delete(heap.index, heap.min.tag)
	heap.num--

	if heap.num == 0 {
		heap.min = nil
	} else {
		heap.consolidate()
	}

	return min
}

func (heap *FibHeap) link(parent, child *node) {
	child.marked = false
	child.parent = parent
	child.self = parent.children.PushBack(child)
	parent.degree++
}

func (heap *FibHeap) resetMin() {
	heap.min = heap.roots.Front().Value.(*node)
	for tree := heap.min.self.Next(); tree != nil; tree = tree.Next() {
		if tree.Value.(*node).key < heap.min.key {
			heap.min = tree.Value.(*node)
		}
	}
}
