package utils

import . "container/list"

type MyStack struct {
    ele *List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{ele: New()}
}

/** Push element x onto stack. */
func (m *MyStack) Push(x int) {
    m.ele.PushFront(x)
}

/** Removes the element on top of the stack and returns that element. */
func (m *MyStack) Pop() int {
    v := m.ele.Front()
    if v == nil {
        return -1
    }
    return m.ele.Remove(v).(int)
}

/** Get the top element. */
func (m *MyStack) Top() int {
    v := m.ele.Front()
    if v == nil {
        return -1
    }
    return v.Value.(int)
}

/** Returns whether the stack is empty. */
func (m *MyStack) Empty() bool {
    return m.ele.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
