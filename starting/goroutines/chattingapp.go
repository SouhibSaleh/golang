package goroutines

import (
	"fmt"
	"net"
)

type UserMessage struct {
	Name string
	Conn net.Conn
}

func handleInput(conn net.Conn, ch chan string) {
	defer func() {

		conn.Close()
	}()
	fmt.Println(conn)
	for {
		buffer := make([]byte, 1024)
		byts, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		message := string(buffer[:byts])
		ch <- message
	}
}
func handleOutput(conn chan net.Conn, ch chan string) {
	defer func() {
		fmt.Println("Done")
	}()
	ls := make([]net.Conn, 0)
	for {
		select {
		case newCon := <-conn:
			ls = append(ls, newCon)
		case message := <-ch:
			fmt.Print("message:", message, "size:", len(ls), "\n")

			for i, l := range ls {

				_, err := fmt.Fprint(l, message)
				if err != nil {
					fmt.Println(err)
					ls = append(ls[:i], ls[i+1:]...)
					continue
				}
			}
		}
	}

}

func ChattingApp() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	connCh := make(chan net.Conn)
	ch := make(chan string)
	go handleOutput(connCh, ch)

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		connCh <- conn
		go handleInput(conn, ch)
	}

}
