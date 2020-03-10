package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"math"
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

	printValidityInfo(x509.NotAfter)

	defer conn.Close()
}

func printValidityInfo(expirationDate time.Time) {

	fifteenDaysBefore := expirationDate.AddDate(0, 0, -15)
	thirtyDaysBefore := expirationDate.AddDate(0, 0, -30)

	now := time.Now()
	timeLeft := expirationDate.Sub(now)
	days := math.Ceil(timeLeft.Hours() / 24)

	if now.After(expirationDate) {
		fmt.Printf("\033[01;31mCRITICAL: \033[0m")
		fmt.Printf("Your server's SSL certificate has already expired on %v\n", expirationDate)
		return
	}

	if now.After(fifteenDaysBefore) {
		fmt.Printf("\033[01;31mCRITICAL: \033[0m")
		fmt.Printf("Your server's SSL certificate is valid until %v. It will expire in %v days.\n", expirationDate, days)
		return
	}

	if now.After(thirtyDaysBefore) {
		fmt.Printf("\033[01;33mWARNING: \033[0m")
		fmt.Printf("Your server's SSL certificate is valid until %v. It will expire in %v days.\n", expirationDate, days)
		return
	}

	fmt.Printf("\033[01;32mINFO: \033[0m")
	fmt.Printf("Your server's SSL certificate is valid until %v. It will expire in %v days.\n", expirationDate, days)
}
