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
		
		u := x==y
		o := ""
		if !u {
			o = "X"
		}
		fmt.Printf("%v: %v\t%v\t%v\n", i+1, x, y, o)
	}

}

func abbreviation(as, bs string) string {
	debug := false
	del := len(as)-len(bs)
	fmt.Println(del)
	X := make([][]int, len(bs), len(bs))
	for i := range X {
		X[i] = make([]int, len(as), len(as))
	}

	a := []rune(as)
	b := []rune(bs)
	
	if debug {
		fmt.Print(" ")
		for i := range a {
			fmt.Printf("  %s", string(a[i]))
		}
		fmt.Print("\n")
	}
	
	for bi := range X {
		if debug {
			fmt.Printf("%s", string(b[bi]))
		}
		for ai := 0; ai < len(X[bi]); ai++ {
			S, A, B := compare(a[ai], b[bi])
			if !S && A || S && A && !B {
				X[bi][ai] = 0	// different and capital, can't do anything "XY" "Xy"  || cannot decapitalize or delete "Xy"
			} else
			if !S && !A {
				X[bi][ai] = 1	// different but can delete "xy" "xY"
			} else
			if S && !A && B {
				X[bi][ai] = 2	// can capitalize "xX"
			} else
			if S && A == B {
				X[bi][ai] = 3	// perfect match "xx" "XX"
			}
		}
		
		if debug {
			err, m := arrayMax(X[bi])
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("% d[% d]\n", X[bi], m)
		}
		
	}
	
	err, maxis, maxjs := arrayMaxes(X)
	if err != nil {
		log.Fatalln(err)
	}
	if debug {
		fmt.Print(" ")
		fmt.Printf("% d\n", maxjs)
	}
	
	// avaliations
	// if lines or rows have 0s
	err, minj := arrayMin(maxjs)
	if err != nil {
		log.Fatalln(err)
	}
	err, mini := arrayMin(maxis)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(minj, mini)
	if minj <= 0 || mini <= 0 {
		return "NO"
	}

	// if maxjs have more 1 (deletions) than can
	// count 1s
	ones := 0
	for _, v := range maxjs {
		if v == 1 {
			ones++
		}
	}
	if ones > del {
		return "NO"
	}
	
	return "YES"
}

func compare(a, b rune) (same bool, A bool, B bool) {
	same = low(a) == low(b)
	A = a < 95
	B = b < 95
	return
}

func low(r rune) rune {
	if r < 95 {
		return r+32
	}
	return r
}

func up(r rune) rune {
	if r > 95 {
		return r-32
	}
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func arrayMaxes(X [][]int) (err error, imax []int, jmax []int) {
	lines := len(X)
	rows := 0
	if lines > 0 {
		rows = len(X[0])
	}
	
	if lines == 0 || rows == 0 {
		return errors.New("zero length array"), nil, nil
	}
	
	resi := make([]int, rows, rows)
	resj := make([]int, rows, rows)
	for i := range resi {
		resi[i] = 0
	}
	for j := range resj {
		resj[j] = 0
	}
	
	for i := range X {
		for j := range X[i] {
			resi[i] = max(resi[i], X[i][j])
			resj[j] = max(resj[j], X[i][j])
		}
	}
	
	fmt.Println(resi, resj)
	
	return nil, resi, resj
}

func arrayMax(x []int) (error, int) {
	if len(x) == 0 {
		return errors.New("zero length array"), 0
	}
	
	m := -999
	for i := range x {
		m = max(m, x[i])
	}
	return nil, m
}

func arrayMin(x []int) (error, int) {
	if len(x) == 0 {
		return errors.New("zero length array"), 0
	}
	
	m := 999
	for i := range x {
		m = min(m, x[i])
	}
	return nil, m
}