package main

import (
	"log"

	"github.com/stianeikeland/go-rpio"
)

var (
//pin goes here
)

//check if gpios are enable
func init() {
	log.Println("opening gpios")
	err := rpio.Open()
	if err != nil {
		log.Panic("Unable to open gpio")
	}
	defer rpio.Close()
}

//code goes here
func main() {

}
