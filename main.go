package main

func main() {
	LoadConfig()
	Init()
	<-make(chan struct{})
	return
}