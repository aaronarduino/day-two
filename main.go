package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	UP          = "U"
	DOWN        = "D"
	RIGHT       = "R"
	LEFT        = "L"
	LIMIT_UP    = 0
	LIMIT_DOWN  = 2
	LIMIT_RIGHT = 2
	LIMIT_LEFT  = 0
)

type keypad struct {
	keys [][]int
	x    int
	y    int
}

func (k *keypad) move(way string) {
	switch way {
	case UP:
		if (k.y - 1) >= LIMIT_UP {
			k.y--
		}
	case DOWN:
		if (k.y + 1) <= LIMIT_DOWN {
			k.y++
		}
	case LEFT:
		if (k.x - 1) >= LIMIT_LEFT {
			k.x--
		}
	case RIGHT:
		if (k.x + 1) <= LIMIT_RIGHT {
			k.x++
		}
	}
}

var (
	out string
	in  []string
)

func main() {
	k := keypad{keys: [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, x: 1, y: 1}

	b := bufio.NewReader(os.Stdin)
	for {
		ln, isPrefix, err := b.ReadLine()
		if err != nil {
			break
		} else if isPrefix {
			log.Fatal(isPrefix)
		}
		in = append(in, string(ln))
	}
	fmt.Println("Lines:", len(in))

	for _, line := range in {
		for i := range line {
			k.move(string(line[i]))
		}
		out += fmt.Sprint(k.keys[k.y][k.x])
	}
	fmt.Println(out)
}
