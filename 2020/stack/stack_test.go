package stack_test

import (
	"testing"

	"github.com/arthurbailao/aoc/2020/stack"
	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	stack := stack.New()

	t.Run("when stack is empty", func(t *testing.T) {
		assert.Equal(t, 0, stack.Len(), "Len() should return 0")
	})

	t.Run("when stack is not empty", func(t *testing.T) {
		stack.Push(1)
		stack.Push(1)
		assert.Equal(t, 2, stack.Len(), "Len() should return the current length")
	})
}

func TestPeek(t *testing.T) {
	stack := stack.New()

	t.Run("when stack is empty", func(t *testing.T) {
		assert.Nil(t, stack.Peek(), "Peek() should return nil")
	})

	t.Run("when stack is not empty", func(t *testing.T) {
		stack.Push(1)
		stack.Push(3)
		assert.Equal(t, 3, stack.Peek(), "Peek() should return top element, without removing it")
		assert.Equal(t, 2, stack.Len(), "Len() should return the actual length")
	})
}

func TestPop(t *testing.T) {
	stack := stack.New()

	t.Run("when stack is empty", func(t *testing.T) {
		assert.Nil(t, stack.Pop(), "Pop() should return nil")
	})

	t.Run("when stack is not empty", func(t *testing.T) {
		stack.Push(1)
		stack.Push(3)
		assert.Equal(t, 3, stack.Pop(), "Pop() should return top element, removing it from stack")
		assert.Equal(t, 1, stack.Len(), "Len() should return the actual length")
	})
}

func TestPush(t *testing.T) {
	stack := stack.New()

	stack.Push(1)
	stack.Push(3)

	assert.Equal(t, 3, stack.Pop(), "Pop() should return top element, removing it from stack")
	assert.Equal(t, 1, stack.Pop(), "Pop() should return top element, removing it from stack")
}
