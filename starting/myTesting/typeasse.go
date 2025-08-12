package test

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s)
	}
	return w.Write([]byte(s))
}

func TypeassePlayground() {
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File) // success: ok, f == os.Stdout
	fmt.Println(f, ok)
	fmt.Println(reflect.TypeOf(w))

	writeString(w, "Hello World")
}
