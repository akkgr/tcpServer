package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	connHost = "localhost"
	connPort = "3333"
	connType = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + connHost + ":" + connPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

	defer func() {
		conn.Close()
	}()

	// // Make a buffer to hold incoming data.
	// buf := make([]byte, 1024)
	// // Read the incoming connection into the buffer.
	// _, err := conn.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {

	timeoutDuration := 30 * time.Second
	bufReader := bufio.NewReader(conn)
	for {
		// Set a deadline for reading. Read operation will fail if no data
		// is received after deadline.
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		// Read tokens delimited by newline
		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s", bytes)
	}

	// uuid, err := newUUID()
	// file, err := os.Create("./messages/" + uuid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close() // make sure to close the file even if we panic.
	// n, err := io.Copy(file, conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(n, "bytes sent")


	// }

	// Send a response back to person contacting us.
	//conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	//conn.Close()
}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
