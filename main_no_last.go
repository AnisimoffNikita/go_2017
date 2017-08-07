package main

import (  
    "fmt"
    "os"
    "io/ioutil"
    "strings"
)

func main() {  
    if len(os.Args) != 2 {
        fmt.Println("No input file")
        return
    } 

    file, err := ioutil.ReadFile(os.Args[1]) 
    if err != nil {
        fmt.Print(err)
    }

    s := string(file)
    n := strings.Count(s, "\n")

    fmt.Println(n)
}