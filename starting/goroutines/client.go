package goroutines

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	fmt.Println("connecting to localhost:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	for {
		go io.Copy(conn, os.Stdin)
		io.Copy(os.Stdout, conn)
	}
}
