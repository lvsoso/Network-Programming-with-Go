package main


import (
	"fmt"
	"net"
	"os"
)


const ECHO_REPLY_HEAD_LEN = 20


func main()  {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
        os.Exit(1)
	}

	addr, err := net.ResolveIPAddr("ip", os.Args[1])
	if err != nil {
		fmt.Println("Resolution error", err.Error())
        os.Exit(1)
	}

	conn, err := net.DialIP("ip4:icmp", addr, addr)
	checkError(err)

	var msg [512]byte
	
	msg[0] = 8	// echo
	msg[1] = 0	// code 0
	msg[2] = 0	// checksum
	msg[3] = 0	// checksum
	msg[4] = 0	// arbitrary identifier[0]
	msg[5] = 13 // arbitrary identifier[1]
	msg[6] = 0 	// arbitrary sequence[0]
	msg[7] = 37 // arbitrary sequence[1]
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8) // get higher 8 bits
	msg[3] = byte(check & 255) // get lower 8 bits

	_, err = conn.Write(msg[0:len])
    checkError(err)



    _, err = conn.Read(msg[0:])
    checkError(err)



    fmt.Println("Got response")
    if msg[ECHO_REPLY_HEAD_LEN+5] == 13 {
        fmt.Println("identifier matches")
    }
    if msg[ECHO_REPLY_HEAD_LEN+7] == 37 {
        fmt.Println("Sequence matches")
    }

	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0

	// assume even for now
    for n := 1; n < len(msg)-1; n += 2 {
        sum += int(msg[n])*256 + int(msg[n+1])
    }

	sum = (sum >> 16) + (sum & 0xffff)
    sum += (sum >> 16)
    var answer uint16 = uint16(^sum)
    return answer
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
        os.Exit(1)
    }
}
