package main

import (
	"fmt"
	"github.com/ebitengine/purego"
)

func main() {
	lib := "liba.so"
	dl, err := purego.Dlopen(lib, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = purego.Dlclose(dl)
		if err != nil {
			panic(err)
		}
	}()
	var add func(a, b int) int
	purego.RegisterLibFunc(&add, dl, "add")
	i := add(1, 2)
	fmt.Println(i)
}
