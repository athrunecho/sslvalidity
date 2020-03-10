package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"
)

func main() {

	var serverAddress string
	fmt.Printf("Please enter server address: ")
	fmt.Scanln(&serverAddress)
	checkCertValidity(serverAddress)

}

func checkCertValidity(address string) {

	address += ":443"

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", address, conf)
	if err != nil {
		log.Println(err)
		return
	}

	err = conn.Handshake()
	if err != nil {
		log.Println(err)
		return
	}

	x509 := conn.ConnectionState().PeerCertificates[0]
	fmt.Printf("expire date: %v", x509.NotAfter)

	defer conn.Close()

	expireDate := x509.NotAfter

	fifteenDaysBefore := expireDate.AddDate(0, 0, -15)
	thirtyDaysBefore := expireDate.AddDate(0, 0, -30)

	now := time.Now()

	if now.After(fifteenDaysBefore) {
		fmt.Printf("\033[01;31m CRITICAL \033[0m\n")
		return
	}

	if now.After(thirtyDaysBefore) {
		fmt.Printf("\033[01;33m WARNING \033[0m\n")
		return
	}

	fmt.Printf("\033[01;32m INFO \033[0m\n")
}
