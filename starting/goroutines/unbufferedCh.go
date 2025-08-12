package goroutines

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func UnbufferedCh() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan string)

	go func() {
		fmt.Println("go routine")
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- ""
	}()
	fmt.Fprintf(conn, "Hello World!\n")
	fmt.Fprintf(conn, "Hello World!\n")
	fmt.Fprintf(conn, "Hello World!\n")
	fmt.Fprintf(conn, "Hello World!\n")

	conn.Close()
	<-done

}
