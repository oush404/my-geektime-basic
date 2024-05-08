package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()

	var a [3]int
	fmt.Println(a[0])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	q := [...]int{1, 2, 3}
	fmt.Printf("%d\n", q[1])

	symbol := [...]string{USD: "$", EUR: "E", GBP: "G", RMB: "R"}
	fmt.Println(RMB, symbol[RMB])
	for i := range symbol {
		fmt.Println(symbol[i])
	}

	b := [...]int{0, 1, 2, 3, 4, 5}
	reverse(b[:3])
	fmt.Println(b)

	c := [...]string{"a", "b", "c", "d"}
	d := [...]string{"a", "b", "c", "d"}
	res := equal(c[:], d[:])
	fmt.Println(res)

	var runes []rune
	for i, r := range "abcdef" {
		fmt.Println(i)
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var ages map[string]int
	if age, ok := ages["bob"]; !ok {
		fmt.Println(age)
	}

	var any interface{}
	any = true
	any = map[string]int{"one": 1, "two": 2}
	fmt.Println(any)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func equalMap(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	return 0
}
