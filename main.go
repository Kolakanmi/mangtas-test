package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("GoLang_Test.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()
	buf := make([]byte, 16)
	var inputBuilder strings.Builder
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to read file: %v", err)
		}
		inputBuilder.Write(buf[:n])
	}
	jsonByte, err := Service(inputBuilder.String())
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(string(jsonByte))
}
