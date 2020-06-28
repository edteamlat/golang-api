package funciones

import (
	"fmt"
	"time"
)

// MyFunction .
type MyFunction func(string)

// MiddlewareLog .
func MiddlewareLog(f MyFunction) MyFunction {
	return func(name string) {
		fmt.Println("inicio:", time.Now().Format("2006-01-02 15:04:05"))
		f(name)
		fmt.Println("fin:", time.Now().Format("2006-01-02 15:04:05"))
	}
}
