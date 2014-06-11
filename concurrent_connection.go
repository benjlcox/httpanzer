package main

import (
  "fmt"
  "net/http"
  "os"
  )


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
      if status == "200 OK" {
        fmt.Println(status)
      }
    }
  }
  fmt.Println("Completed.")
}

func main() {

  test_url := "http://benify.myshopify.com/collections/all/products.json?limit=250"
  number := 100

  for i := 0; i < 10; i++ {
    go do_connect(test_url, number)
  }
  var input string
  fmt.Scanln(&input)
}

