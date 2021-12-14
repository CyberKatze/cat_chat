package client

import "fmt"

var Port int

func init() {

}

func Start() {
	fmt.Printf("CLient startet and connected to %v", Port)
}

func Set(port int) {
	Port = port
}
