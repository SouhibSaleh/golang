package goroutines

import (
	"fmt"
	"io"
	"net"
	"time"
)

func handleConnection(c net.Conn) {

	defer c.Close()
	for {
		buffer := make([]byte, 1024)
		byts, err := c.Read(buffer)
		message := string(buffer[:byts])
		fmt.Print("Reading:", message)
		if err != nil {
			fmt.Println("read error:", err)
			return
		}
		_, err = io.WriteString(c, "alive:"+time.Now().String()+": "+message)
		if err != nil {
			return
		}
		time.Sleep(time.Second * 1)
	}
}

func Clockserver() {

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer listen.Close()
	i := 0
	for {

		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println(conn.RemoteAddr(), i)
		i++
		go handleConnection(conn)
	}
}
