// Реализовать собственную функцию sleep.

package main

import "time"

func sleep(s int) {
	<-time.After(time.Second * time.Duration(s))
}

func main() {
	sleep(3)
}
