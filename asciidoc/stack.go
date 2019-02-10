package asciidoc

import "errors"

// Stack is an implementation of a stack of integers. That is all
// we need to keep track of heading levels.
type Stack struct {
  data []int
  size int
}

func NewStack() *Stack {
  stack := &Stack{}
  return stack
}

func (stack *Stack) Push(x int) {
  stack.data = append(stack.data, x)
  stack.size++
}

func (stack Stack) Empty() bool {
  return (stack.size == 0)
}

func (stack Stack) Top() (value int, err error) {
  if stack.size == 0 {
    return 0, errors.New("Empty stack")
  }
  return stack.data[stack.size - 1], nil
}

func (stack *Stack) Pop() (value int, err error) {
  if stack.size == 0 {
    return 0, errors.New("Empty stack")
  }
  res := stack.data[stack.size - 1]
  stack.size--
  return res, nil
}
