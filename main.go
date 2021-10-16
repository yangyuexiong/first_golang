package main

import (
	A "first_golang/def"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

//go env -w GO111MODULE=on
//go env -w GOPROXY=https://goproxy.io,direct

func init() {
	fmt.Println("init.......")
}

func testDecimal() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875

}

func test_gjson() {

	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	println(value.String())

}
func main() {
	A.Def("yyx")
	//def.Def("yyx")

	fmt.Println(A.Okc)
	//fmt.Println(def.Okc)

	r := A.Def("yyx")
	fmt.Println(r)

	//def.def("yyx")
	//fmt.Println(def.okc)

	testDecimal()

	test_gjson()
}
