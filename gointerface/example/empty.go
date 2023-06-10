package example

import "fmt"

func EmptyRun() {
	var i interface{}
	Describe(i)

	i = 42
	Describe(i)

	i = "hello"
	Describe(i)
}

func Describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func EmptyWithMap() {
    product := make(map[string]interface{}, 0)

    product["name"] = "Iphone 13 Pro Max"
    product["price"] = 31000000
    product["quantity"] = 40

    fmt.Println(product)
}

func TypeAssertions() {
		var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}