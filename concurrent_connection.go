package main

import (
  "fmt"
  "net/http"
  "os"
  )

var completed int = 0

func do_connect(url string, n int){
  for i := 0; i < n; i++ {
    response, err := http.Get(url)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    } else {
      defer response.Body.Close()
      status := string(response.Status)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      if status != "200 OK" {
        fmt.Println(status)
      }
    }
  }

  completed = completed + 1

  fmt.Println("Completed Routine: ", completed)
  if completed == 10 {
    fmt.Println("ALL DONE :D")
  }
}

func main() {

  test_url := "http://benify.myshopify.com/collections/all/products.json?limit=250"

  var url string
  fmt.Println("Enter URL:")
  fmt.Scanln(&url)
  if url == ""{
    url = test_url
  }

  var number int
  fmt.Println("Enter connections per routine (ten routines): ")
  fmt.Scanln(&number)

  for i := 0; i < 10; i++ {
    go do_connect(test_url, number)
  }
  var input string
  fmt.Scanln(&input)
  fmt.Println(input)
}

