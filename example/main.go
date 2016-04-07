package main

import "fmt"
import "github.com/mrcsparker/go-julia"

func main() {
	j := julia.New()

	//o, e := j.Eval("import JSON; JSON.json([1, 2])")
	o, e := j.Eval("using JSON; JSON.json([1,2])")
	fmt.Println(e)
	fmt.Println(o.(string))

	defer j.Free()
}
