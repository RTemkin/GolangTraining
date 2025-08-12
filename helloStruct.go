package main

import (
	"fmt"
	"math"
)

type Arifm interface {
	Ploshd() float32
	Perimetr() float32
}

type Pryamoug struct {
	a float32
	b float32
}

type Krug struct {
	radius float32
}

func (pr Pryamoug) Ploshd() float32 {
	if pr.a > 0 && pr.b > 0 {
		return pr.a * pr.b
	}
	return -1
}

func (pr Pryamoug) Perimetr() float32 {
	return 2 * (pr.a + pr.b)
}

func (kr Krug) Ploshd() float32 {
	return 2 * math.Pi * (kr.radius * kr.radius)
}

func (kr Krug) Perimetr() float32 {
	return math.Pi * kr.radius
}

func GetPloshad(ar Arifm) float32 {
	return ar.Ploshd()
}

func GetPerim(ar Arifm) float32 {
	return ar.Perimetr()
}

func Fibon(n int) int {
	sliceFib := []int{0, 1}

	for i := 2; i < n; i++ {
		sliceFib = append(sliceFib, sliceFib[i-2]+sliceFib[i-1])
	}
	return sliceFib[len(sliceFib)-1]
}

func Fibon2(n int) int {
	a := 0
	b := 1
	var val int
	for i := 0; i < n-2; i++ {
		val = a + b
		a = b
		b = val
	}
	return val
}

func main() {

	pryam1 := Pryamoug{
		a: 4,
		b: 5,
	}

	krug1 := Krug{
		8,
	}

	sliseFigur := []Arifm{pryam1, krug1}

	for _, val := range sliseFigur {
		fmt.Println(GetPloshad(val))
		fmt.Println(GetPerim(val))

	}

	// sPr := GetPloshad(pryam1)
	// sKr := GetPloshad(krug1)

	// perPr := GetPerim(pryam1)
	// perKr := GetPerim(krug1)

	// fmt.Println(sPr, sKr)
	// fmt.Println(perPr, perKr)

	fmt.Println(Fibon(10))
	fmt.Println(Fibon2(9))
}
