package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type Field struct {
	field [][]bool
	w, h  int
}

func (f *Field) Set(x, y int, b bool) {
	f.field[x][y] = b
}

func (f *Field) Alive(x, y int) bool {
	if x < 0 {
		x += f.h
	}
	if x >= f.h {
		x -= f.h
	}
	if y < 0 {
		y += f.w
	}
	if y >= f.w {
		y -= f.w
	}

	return f.field[x][y]
}

func (f *Field) Next(x, y int) bool {
	alive := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.Alive(x+i,y+j) {
				alive++
			}
		}
	}

	return alive == 3 || alive == 2 && f.Alive(x, y)
}

func NewField(w, h int) *Field {
	s := make([][]bool, h)

	for i := range s {
		s[i] = make([]bool, w)
	}

	return &Field{s, w, h}
}

type Life struct {
	a, b *Field
	w, h int
}

func (l *Life) Step() {
	for x := 0; x < l.h; x++ {
		for y := 0; y < l.w; y++  {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}

	l.a, l.b = l.b, l.a
}

func (l *Life) String() string {
	var buf bytes.Buffer

	for x := 0; x < l.h; x++ {
		for y := 0; y < l.w; y++  {
			b := byte(' ')
			if l.a.Alive(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}

	return buf.String()
}

func NewLife(w, h int) *Life {
	a := NewField(w, h)
	for i := 0; i < (w * h / 4); i++ {
		a.Set(rand.Intn(h), rand.Intn(w), true)
	}
	return &Life{a, NewField(w, h), w, h }
}


func main() {
	rand.Seed(time.Now().Unix())
	l := NewLife(40, 15)
	for i := 0; i < 300; i++ {
		l.Step()
		fmt.Print(l)
		//time.Sleep(time.Second / 30)
	}
}
