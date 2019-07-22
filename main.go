package main

import (
	//"bufio"
	"fmt"
	"net"
	"time"
)

// State represents the state of the Tello, it's expected to be in the following format e.g.
// "pitch:0;roll:0;yaw:0;vgx:0;vgy:0;vgz:0;templ:89;temph:91;tof:10;h:0;bat:3;baro:48.26;time:0;agx:0.00;agy:-6.00;agz:-1003.00;"
//
// Source: file:///Users/danish/Downloads/Tello%20SDK%20Documentation%20EN_1.3_1122.pdf
type State struct {
	Pitch              int
	Roll               int
	Yaw                int
	VGX                int
	VGY                int
	VGZ                int
	LowestTemperature  int
	HighestTemperature int
	TOF                int
	Height             int
	Battery            int
	Barometer          float64
	MotorsOnTime       int
	AGX                float64
	AGY                float64
	AGZ                float64
}

func main() {

	time.Sleep(2 * time.Second)
	// Send a Command
	send("command", "192.168.10.1:8889")
	send("remark2", "192.168.10.1:8890")

	// Start UDP listner
	listner()

}

func send(cmd string, addr string) {
	fmt.Printf("Sending CMD [%s] ...", cmd)
	//p := make([]byte, 2048)
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

	fmt.Printf("cmd [%s] sent", cmd)
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
