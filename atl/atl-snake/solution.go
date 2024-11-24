package main

import "fmt"

/**
head at end
tail at beginning
[(0,0), (0,1), (0,2)]
*/

// Further down the line we can figure out how fast ArraySnake vs LinkedList snake is.
type Snake interface {
	Pop() [2]int
	Append([2]int)
}

type ArraySnake struct {
	snake [][2]int
}

func (s ArraySnake) Pop() [2]int {
	a := s.snake[0]
	s.snake = s.snake[1:]
	return a
}

func (s ArraySnake) Append(pos [2]int) {
	s.snake = append(s.snake, pos)
}

type SnakeGame struct {
	snake     [][2]int
	moves     int
	growEvery int
	width     int
	height    int
}

func NewSnakeGame(width, height, growEvery int) *SnakeGame {
	return &SnakeGame{
		snake:     make([][2]int, 0, width*height),
		growEvery: growEvery,
		width:     width,
		height:    height,
	}
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
		s.snake = append(s.snake, [2]int{(oldHead[0] + 1) % s.height, oldHead[1]})
	case "down":
		s.snake = append(s.snake, [2]int{(oldHead[0] - 1) % s.height, oldHead[1]})
	case "left":
		s.snake = append(s.snake, [2]int{oldHead[0], (oldHead[1] - 1) % s.width})
	case "right":
		s.snake = append(s.snake, [2]int{oldHead[0], (oldHead[1] + 1) % s.width})
	default:
		panic(fmt.Errorf("invalid direction"))
	}
}
