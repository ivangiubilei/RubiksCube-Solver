package main

import "fmt"

const (
	red = iota
	green
	orange
	blue
	white
	yellow
)

type Cube struct {
	F []int
	R []int
	B []int
	L []int
	U []int
	D []int
}

func (c *Cube) New() {
	for range 9 {
		c.F = append(c.F, red)
		c.R = append(c.R, green)
		c.B = append(c.B, orange)
		c.L = append(c.L, blue)
		c.U = append(c.U, white)
		c.D = append(c.D, yellow)
	}
}

func Transpose(vec []int) []int {
	transposed := make([]int, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			transposed[j*3+i] = vec[i*3+j]
		}
	}
	return transposed
}

func ReverseRows(vec []int) []int {
	return []int{vec[2], vec[1], vec[0], vec[5], vec[4], vec[3], vec[8], vec[7], vec[6]}
}

func (c *Cube) TurnF() {
	Uface := c.U
	Dface := c.D
	Rface := c.R
	Lface := c.L
	c.F = ReverseRows(Transpose(c.F))
	c.R = []int{Uface[0], Rface[1], Rface[2], Uface[1], Rface[4], Rface[5], Uface[2], Rface[7], Rface[8]}
	c.D = []int{Rface[2], Rface[1], Rface[0], Dface[3], Dface[4], Dface[5], Dface[6], Dface[7], Dface[8]}
	c.L = []int{Lface[0], Lface[1], Dface[0], Lface[3], Lface[4], Dface[1], Lface[5], Lface[6], Dface[2]}
	c.U = []int{Uface[0], Uface[1], Uface[2], Uface[3], Uface[4], Uface[5], Lface[8], Lface[5], Lface[2]}
}

func (c *Cube) TurnF2() {
	c.TurnF()
	c.TurnF()
}

func main() {
	cube := Cube{}
	cube.New()
	cube.TurnF()
	fmt.Println(cube)
}
