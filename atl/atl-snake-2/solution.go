package atl_snake_2

import "fmt"

/**
head at end
tail at beginning
[(0,0), (0,1), (0,2)]



Rules:

Every time moveSnake() is called, the snake moves up, down, left or right
The snake’s initial size is 3 and grows by 1 every 5 moves
The game ends when the snake hits itself

*/

type Position [2]int

// Further down the line we can figure out how fast ArraySnake vs LinkedList snake is.
type Snake interface {
	Pop() Position
	Append(Position)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

type ArraySnake struct {
	snake []Position
}

const initialSize = 3

func (s ArraySnake) Pop() Position {
	a := s.snake[0]
	s.snake = s.snake[1:]
	return a
}

func (s ArraySnake) Append(pos Position) {
	s.snake = append(s.snake, pos)
}

type SnakeGame struct {
	snake     []Position
	moves     int
	growEvery int
	width     int
	height    int
}

func NewSnakeGame(width, height, growEvery int) *SnakeGame {
	game := &SnakeGame{
		snake:     make([]Position, initialSize, width*height),
		growEvery: growEvery,
		width:     width,
		height:    height,
	}
	game.snake[0] = Position{0, 0}
	game.snake[1] = Position{0, 1}
	game.snake[2] = Position{0, 2}
	return game
}

// IsGameOver checks whether the head is hitting other part of snake
func (s *SnakeGame) IsGameOver() bool {
	head := s.snake[len(s.snake)-1]
	for _, part := range s.snake[:len(s.snake)-1] {
		if part[0] == head[0] && part[1] == head[1] {
			return true
		}
	}
	return false
}

func (s *SnakeGame) MoveSnake(direction string) {
	s.moves++
	if s.moves%s.growEvery != 0 {
		s.snake = s.snake[1:] // chops off tail
	}
	oldHead := s.snake[len(s.snake)-1]
	switch direction {
	case "up":
		s.snake = append(s.snake, Position{mod(oldHead[0]+1, s.height), oldHead[1]})
	case "down":
		s.snake = append(s.snake, Position{mod(oldHead[0]-1, s.height), oldHead[1]})
	case "left":
		s.snake = append(s.snake, Position{oldHead[0], mod(oldHead[1]-1, s.width)})
	case "right":
		s.snake = append(s.snake, Position{oldHead[0], mod(oldHead[1]+1, s.width)})
	default:
		panic(fmt.Errorf("invalid direction"))
	}
}
