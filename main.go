package main

import (
	"fmt"
	"gitlab.com/akbar1482027/golangconverter/Convert"
)


func main() {
	res :=convert.PrsToInt("2")
    res1 := convert.ToString(2)
	fmt.Print(res , res1)
}