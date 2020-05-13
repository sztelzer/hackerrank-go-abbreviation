package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	hk "github.com/sztelzer/hackerrank"
)

const testsPath string = "tests"

func main() {
	tests := hk.NewTests(testsPath)
	tests.Run(Solution)
}

func Solution(t hk.Test) {
	fmt.Printf("Test: %s\n", t.Name)
	line := t.In.NextLine()
	n, err := strconv.Atoi(line)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < n; i++ {
		a := t.In.NextLine()
		b := t.In.NextLine()
		y := abbreviation(a, b)
		x := t.Out.NextLine()

		u := x == y
		o := ""
		if !u {
			o = "X"
		}
		fmt.Printf("%v: %v\t%v\t%v\n", i+1, x, y, o)
	}

}

func abbreviation(as, bs string) string {

}

// func compare(a, b rune) (same bool, A bool, B bool) {
// 	same = low(a) == low(b)
// 	A = a < 95
// 	B = b < 95
// 	return
// }
//
// func low(r rune) rune {
// 	if r < 95 {
// 		return r + 32
// 	}
// 	return r
// }
//
// func up(r rune) rune {
// 	if r > 95 {
// 		return r - 32
// 	}
// 	return r
// }
//
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
//
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
//
// func arrayMaxes(X [][]int) (err error, imax []int, jmax []int) {
// 	lines := len(X)
// 	rows := 0
// 	if lines > 0 {
// 		rows = len(X[0])
// 	}
//
// 	if lines == 0 || rows == 0 {
// 		return errors.New("zero length array"), nil, nil
// 	}
//
// 	resi := make([]int, rows, rows)
// 	resj := make([]int, rows, rows)
// 	for i := range resi {
// 		resi[i] = 0
// 	}
// 	for j := range resj {
// 		resj[j] = 0
// 	}
//
// 	for i := range X {
// 		for j := range X[i] {
// 			resi[i] = max(resi[i], X[i][j])
// 			resj[j] = max(resj[j], X[i][j])
// 		}
// 	}
//
// 	fmt.Println(resi, resj)
//
// 	return nil, resi, resj
// }
//
// func arrayMax(x []int) (error, int) {
// 	if len(x) == 0 {
// 		return errors.New("zero length array"), 0
// 	}
//
// 	m := -999
// 	for i := range x {
// 		m = max(m, x[i])
// 	}
// 	return nil, m
// }
//
// func arrayMin(x []int) (error, int) {
// 	if len(x) == 0 {
// 		return errors.New("zero length array"), 0
// 	}
//
// 	m := 999
// 	for i := range x {
// 		m = min(m, x[i])
// 	}
// 	return nil, m
// }
