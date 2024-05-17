package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Type 'x' at any moment to exit")
    fmt.Println()
    for {
        fmt.Println("> Enter the city name : ")
        fmt.Print("> ")
        input, err := reader.ReadString('\n')
        if err != nil{
            os.Exit(1)
        }
        input = strings.TrimSpace(input)
        if input == "x"{
            fmt.Print("Exiting...")
            os.Exit(0)
        }
        fmt.Println(input)
    }
}