package main

import (
	"fmt"
	"testing"
)

var game1 = [][]string{
	{"X", "", "", ""},
	{"X", "", "", ""},
	{"X", "", "", ""},
	{"X", "", "", ""},
}

var game2 = [][]string{
	{"X", "", "O", "O"},
	{"X", "X", "X", "X"},
	{"X", "", "", ""},
	{"O", "", "", ""},
}

var game3 = [][]string{
	{"X", "X", "", ""},
	{"O", "X", "", ""},
	{"X", "X", "", ""},
	{"X", "X", "", ""},
}

var game4 = [][]string{
	{"X", "X", "", ""},
	{"O", "X", "", ""},
	{"", "X", "X", ""},
	{"X", "", "", "X"},
}
var game5 = [][]string{
	{"X", "", "", "X"},
	{"O", "X", "X", ""},
	{"", "X", "X", ""},
	{"X", "", "", ""},
}

func TestLeftColumn(t *testing.T) {
	if getWinner(game1) != "X" {
		t.Fail()
	} else {
		fmt.Println("success")
	}
}

func TestMidColumn(t *testing.T) {
	if getWinner(game3) != "X" {
		t.Fail()
	}
}

func TestRows(t *testing.T) {
	if getWinner(game2) != "X" {
		t.Fail()
	}
}

func TestDiagonals(t *testing.T) {
	if getWinner(game4) != "X" {
		t.Fail()
	}
	if getWinner(game5) != "X" {
		t.Fail()
	}
}
