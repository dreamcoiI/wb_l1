package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := new(big.Int)
	y := new(big.Int)

	_, ok := x.SetString("100000000000000000", 10)
	if !ok {
		fmt.Println("Wrong data")
	}

	_, ok = y.SetString("15", 10)
	if !ok {
		fmt.Println("Wrong data")
	}

	ans := new(big.Int)
	fmt.Println(ans)

	ans.Add(x, y)
	fmt.Println(ans)

	ans.Sub(x, y)
	fmt.Println(ans)

	ans.Mod(x, y)
	fmt.Println(ans)

	if y.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("DIVISION BY ZERO")
	} else {
		ans.Div(x, y)
		fmt.Println(ans)
	}

}
