package main

import (
	//"bufio"
	// "bufio"
	"fmt"
	"net"
	"time"
)

func main() {

	time.Sleep(2 * time.Second)
	// Send a Command
	send("command", "192.168.10.1:8889")
	time.Sleep(2 * time.Second)
	send("takeoff", "192.168.10.1:8889")
	time.Sleep(5 * time.Second)
	send("cw 90", "192.168.10.1:8889")
	time.Sleep(5 * time.Second)
	send("land", "192.168.10.1:8889")

	// Start UDP listner
	// listner()

}

func send(cmd string, addr string) {
	fmt.Printf("Sending CMD [%s] ...\n", cmd)
	// p := make([]byte, 2048)
	conn, err := net.Dial("udp", addr)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	_, err = fmt.Fprintf(conn, cmd)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
	}
	// _, err = bufio.NewReader(conn).Read(p)
	// if err == nil {
	// 	fmt.Printf("%s\n", p)
	// } else {
	// 	fmt.Printf("Some error %v\n", err)
	// }
	conn.Close()

	fmt.Printf("cmd [%s] sent\n", cmd)
}

func listner() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 8890,
		IP:   net.ParseIP("0.0.0.0"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	fmt.Println("UDP Server Started ....")
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
	}
}
