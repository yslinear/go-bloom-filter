package main

import "fmt"

func main() {
	f := filter{}
	f.Set("hello")
	fmt.Printf("%#v\n", f.bitfield)
	fmt.Println(f.Get("hello"))
}
