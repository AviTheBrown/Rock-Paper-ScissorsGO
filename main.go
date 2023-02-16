package main

import (
    "bufio"
    "fmt"
    "os"
)


func setUserName() string {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter your name: ")
    scanner.Scan()

    name := scanner.Text()
    fmt.Printf("Hello %s\n", name)
    return name
}

func createFileForPlayerRecords(){
    ratingFile, err := os.Create("Ratings.txt")
    if  err != nil{
        fmt.Println(err.Error())
        return
    }
    defer ratingFile.Close()

}

func main() {
    name := setUserName()
    names := []string{name}
    fmt.Println(names)
    fmt.Println(len(names))
    createFileForPlayerRecords()
}