package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

type cube struct {
	square
}

func (c cube) volume() float64 {
	return math.Pow(c.length, 3)
}

type shape interface {
	area() float64
}

type object interface {
	shape
	volume() float64
}

func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func areaVolumeSum(shapes ...object) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

func main() {
	s := square{length: 4}
	c := cube{s}

	fmt.Println(areaSum(s, c))
	fmt.Println(areaVolumeSum(c))
}

