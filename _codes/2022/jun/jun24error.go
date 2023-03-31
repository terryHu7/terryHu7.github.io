package main

import "fmt"

func main() {
	err := returnerrorf()
	if err.Error() == "EcsOverTcp ems restart not ready" {
		fmt.Println("pass")
	}
	fmt.Println("end")
}

func returnerrorf() error {
	return fmt.Errorf("EcsOverTcp ems restart not ready")
}
