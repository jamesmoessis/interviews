package atl_snake_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnakeGame(t *testing.T) {
	game := NewSnakeGame(10, 10, 2)
	assert.False(t, game.IsGameOver())
	assert.Equal(t, []Position{{0, 0}, {0, 1}, {0, 2}}, game.snake) // snake starts bottom left facing right
	game.MoveSnake("right")
	assert.False(t, game.IsGameOver())
	assert.Equal(t, []Position{{0, 1}, {0, 2}, {0, 3}}, game.snake)
	game.MoveSnake("down")
	assert.Equal(t, []Position{{0, 1}, {0, 2}, {0, 3}, {9, 3}}, game.snake) // passes through bottom, and grows
	game.MoveSnake("left")
	assert.False(t, game.IsGameOver())
	game.MoveSnake("up")
	assert.True(t, game.IsGameOver())
}
