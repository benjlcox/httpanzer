package main

import (
  "fmt"
  "net/http"
  )

var completed int = 0
var ok int = 0

func do_connect(url string, n int){
  for i := 0; i < n; i++ {
    response, err := http.Get(url)
    if err != nil {
      fmt.Println(err)
    } else {
      defer response.Body.Close()
      status := string(response.Status)
      if err != nil {
        fmt.Println(err)
      }
      if status != "200 OK" {
        fmt.Println(status)
      } else {
        ok = ok + 1
      }
    }
  }

  completed = completed + 1

  fmt.Println("Completed Routine: ", completed)
  if completed == 10 {
    fmt.Println("ALL DONE :D")
    fmt.Println("Number of 200 OK: ", ok)
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

