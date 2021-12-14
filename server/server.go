package server

import "fmt"

var Port int

func Init() {

}

func Start() {
	fmt.Printf("Server is listen to Port %v", Port)
}

func Set(port int) {
	Port = port
}
