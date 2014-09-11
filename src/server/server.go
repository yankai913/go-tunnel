package server

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

var (
	serverPort = flag.String("port", "6789", "server listen port")
	serverHost = flag.String("host", "0.0.0.0", "server host")
	file       = flag.String("file", "", "transfer file")
)

func Startup() {
	flag.Parse()
	port := fmt.Sprintf(":%s", *serverPort)
	ip := fmt.Sprintf(":%s", *serverHost)
	filePath := fmt.Sprintf(":%s", *file)

	host := ip + port
	fmt.Println(host)
	fmt.Println(filePath)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	for {
		conn, err2 := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleRequest(con, filePath)
	}
}

func handleRequest(con net.Conn, filePath string) {
	file, err = os.Open(filePath)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024*2)

	for {
		reader := bufio.NewReader(file)
		line, err2 := reader.ReadString('\n')
		if err2 == nil {
			fmt.Println("reader error")
			break
		}

		data := []byte(line)
		header := []byte{len(data)}
		con.Write(header)
		con.Write(byte(line))
	}
	con.Close()
}
