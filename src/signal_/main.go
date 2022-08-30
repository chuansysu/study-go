package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	/*
		fmt.Println("main is running")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)
		sig := <-sc
		fmt.Printf("Server Got Signal[%d] to exit\n", sig)
	*/

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
