
package main

import (
	"fmt"
	"time"
)

func main () {
	fmt.Println("Vítej na dyzgortu")
	time.Sleep(2 * time.Second)
	fmt.Println("Voláš kamarádovi")
	time.Sleep(4 * time.Second)
	fmt.Println("Kamarád ti to nevzal")
}

