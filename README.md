# sslvalidity
[![Build Status](https://travis-ci.org/athrunecho/sslvalidity.svg?branch=master)](https://travis-ci.org/athrunecho/sslvalidity)

This is a small golang program that check the expiration date of SSL certificate on server.

### How to run the program

The program should run on Linux system. Please install Go on your environment. [Install the Go tools](https://golang.org/doc/install)

Run the code with command:

 	$ go run ssl.go

If you want to build a executable file:

 	$ go build ssl.go

### User input

User can enter the domain name or IP address to check certificate expiration date.

Example:
	"8.8.8.8"
	"google.com"
	"localhost"

### License
[MIT License](./LICENSE)
