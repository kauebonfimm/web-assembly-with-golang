package main

import (
	"fmt"
	"syscall/js"

	generator "github.com/kauebonfimm/web-assembly-with-golang/internal/core/password_engine"
)

func GeneratePassword(this js.Value, p []js.Value) interface{} {
	length, hasLetter, hasNumber, hasSpecialChar, removeCaracters := uint16(p[0].Int()), p[1].Bool(), p[2].Bool(), p[3].Bool(), p[4].String()

	password, err := generator.GeneratePassword(length, hasLetter, hasNumber, hasSpecialChar, removeCaracters)
	if err != nil {
		return err.Error()
	}

	return password
}

func main() {
	fmt.Println("Loading WebAssembly...")

	c := make(chan struct{}, 0)

	js.Global().Set("generate", js.FuncOf(GeneratePassword))

	<-c
}
