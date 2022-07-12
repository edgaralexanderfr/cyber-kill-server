package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println(`Running`, js.Global().Get("document").Get("title"), `...`)
	fmt.Println(`Executed at:`, js.Global().Get("Date").New().Call("toDateString"))
}
