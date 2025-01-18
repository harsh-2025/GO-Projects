package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func sayHello(s string) {
	s = "Hello, World"
}
func sayHelloPointer(s *string) {
	*s = "Hello, World"
}
func main() {
	fmt.Println("Hello", "Harsh")
	var agency string = "Fast track"
	fmt.Println(agency)
	var two int = 2
	fmt.Println("two", two)
	startingPrice := 29.54
	fmt.Printf("our price is at %v", startingPrice)
	fmt.Println("starting price is", startingPrice)
	insuranceInclude := false
	fmt.Println("insurance is not", insuranceInclude)
	// string
	str1 := "abcd"
	str2 := "Abcd"
	fmt.Println(strings.EqualFold(str1, str2)) // it measures only words not capital/small
	cIndex := strings.Index(str1, "c")
	// it compares capital/small and is 0 based
	fmt.Println(cIndex)
	str3 := "Go learning Index"
	str3 = "Go index"
	fmt.Println(strings.Contains(str3, "index"))
	fmt.Println(strings.Contains(str3, "Index"))
	temperatureC := 32
	temperatureK := 0.0
	temperatureK = float64(temperatureC) + 273.15
	fmt.Printf("temperature in kelvin is %v\n", temperatureK)
	var temperatureF float64
	temperatureF = float64(temperatureC)*9/5 + 32
	fmt.Println("temperature in farenhit", temperatureF)
	var tempc float64 = 32.34
	fmt.Println("temperature is", tempc)
	fmt.Println("temp round", math.Round(tempc))
	fmt.Println("temp ceil", math.Ceil((tempc)))
	fmt.Println("temp floor", math.Floor(tempc))
	fmt.Println(math.Pow(3, 2))
	fmt.Printf("in to float %.2f\n", 43.43)
	var str4 string = "321"
	var inttostr, _ = strconv.Atoi(str4)
	fmt.Println("string to int is", inttostr)
	var myfloatstr string = "32.432"
	var myfloatfromstring, _ = strconv.ParseFloat(myfloatstr, 43)
	fmt.Println("float", myfloatstr)
	fmt.Printf("%T\n", myfloatfromstring)
	newbool, _ := strconv.ParseBool("T")
	fmt.Println("new is", newbool)
	const Agency string = "fesfsd"
	// Agency="fewfs" cannot be reassigned
	fmt.Println("agency new is ", Agency)
	const (
		_ = iota
		Economy
		Compact
	)
	fmt.Println("eco", Economy)
	// fmt.Println("What is name")
	// var name string
	// fmt.Scanln(&name)
	// fmt.Println("name is ", name)
	// pointers read;
	var pointerstring *string
	var mystringname string
	mystringname = "hero lal"
	pointerstring = &mystringname
	fmt.Println("my name is ,", mystringname)
	fmt.Println("pointer is ", pointerstring)
	// pointer ;var gr
	var greetings string = "Hello main"
	sayHello(greetings)
	fmt.Println("after calling sayhello", greetings)
	sayHelloPointer(&greetings)
	fmt.Println("greeting after sayhelloptr", greetings)
	var stringptr *string
	var intPtr *int
	fmt.Println("printing the pointers", stringptr, intPtr)
}
