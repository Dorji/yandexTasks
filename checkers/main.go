package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var NM string
	var w int

	var b int

	var color string

	fmt.Scan(&NM) //NM
	tmpNM := strings.Split(NM, " ")
	fmt.Println(NM)

	N, _ := strconv.Atoi(tmpNM[0])
	M, _ := strconv.Atoi(tmpNM[1])

	wArr := make([][]bool, N)
	tmpColorArr := make([]bool, M)

	for k := range wArr {
		wArr[k] = tmpColorArr
	}

	bArr := make([][]bool, N)
	for k := range wArr {
		bArr[k] = tmpColorArr
	}

	fmt.Scanln(&w) //w

	var tmpArr []string

	var coord string

	for i := 0; i < w; i++ {
		fmt.Scanln(&coord)
		tmpArr = strings.Split(coord, " ")
		fmt.Println(tmpArr)
		a, _ := strconv.Atoi(tmpArr[0])
		b, _ := strconv.Atoi(tmpArr[1])
		wArr[a][b] = true
	}

	fmt.Scanln(&b) //b
	for i := 0; i < b; i++ {
		fmt.Scanln(&coord)
		tmpArr = strings.Split(coord, " ")
		a, _ := strconv.Atoi(tmpArr[0])
		b, _ := strconv.Atoi(tmpArr[1])
		bArr[a][b] = true
	}

	fmt.Scanln(&color) //color

	for _, v := range wArr {
		fmt.Println(v)
	}
	for _, v := range bArr {
		fmt.Println(v)
	}

}
