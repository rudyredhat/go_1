package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intSum)

	f1, f2, f3 := 12.234, 45.423, 68.4234
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum: ", floatSum)

	//big numbers for the precision
	var b1, b2, b3, bigSum big.Float
	//SetFloat64 to set the values to the variables w/o using :=
	b1.SetFloat64(23.5)
	b1.SetFloat64(65.1)
	b1.SetFloat64(76.3)
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("BigSum = %.10g\n", &bigSum)
	//if w/p & the bigSum is printed something like this = BigSum = {%!g(uint32=0000000053) %!g(big.RoundingMode=0000000000)
	// %!g(big.Accuracy=0000000000) %!g(big.form=0000000001) %!g(bool=false) [%!g(big.Word=1099598
	// 8850187802624)] %!g(int32=0000000007)}

	circleRa := 15.5
	circleCu := circleRa * math.Pi
	//.2f = 2 characters after the decimal point
	fmt.Printf("Circumference: %.2f\n", circleCu)

}
