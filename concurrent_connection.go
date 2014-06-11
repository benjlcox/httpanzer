package main

import (
  "fmt"
  "net/http"
  )

var completed int = 0
var ok int = 0

func do_connect(url string, routines int, n int){
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
        ok++
      }
    }
  }

  completed++

  fmt.Println("Completed Routine: ", completed)
  if completed == routines {
    fmt.Println("ALL DONE :D")
    fmt.Println("Number of 200 OK: ", ok)
  }
}

func get_url() string{
  test_url := "http://benify.myshopify.com/collections/all/products.json?limit=250"

  var url string
  fmt.Println("Enter URL:")
  _, err := fmt.Scanln(&url)
  if err != nil && err.Error() != "unexpected newline" {
    fmt.Println(err)
  }
  if url == ""{
    url = test_url
  }
  return url
}

func get_routines() int{
  var routines int
  fmt.Println("Enter number of routines (5 reccomended): ")
  _, err := fmt.Scanln(&routines)
  if err != nil {
    fmt.Println(err)
  }
  return routines
}

func get_number() int{
  var number int
  fmt.Println("Enter connections per routine: ")
  _, err := fmt.Scanln(&number)
  if err != nil {
    fmt.Println(err)
  }
  return number
}

func main() {

  url := get_url()
  routines := get_routines()
  number := get_number()

  for i := 0; i < routines; i++ {
    go do_connect(url, routines, number)
  }
  var input string
  fmt.Scanln(&input)
  fmt.Println(input)
}

