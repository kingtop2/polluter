package main

import (
	"log"
	"os"
)

var data = [...]byte{2, 5, 5, 1, 2, 5, 12, 5, 2, 12, 25, 2, 1, 2, 5, 224, 211, 111, 134}

func main() {

	f, err := os.OpenFile("fat.tat", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer f.Close()

	for i := 0; i < len(data) * 0xfffffff; i++ {
		f.Write(data[:])
	}

	log.Println("Operation finished.")
}