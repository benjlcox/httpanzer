package main

import (
  "fmt"
  "net/http"
  "os"
  )

func main() {
  response, err := http.Get("http://benify.myshopify.com/collections/all/products.json?limit=250")
  if err != nil {
      fmt.Printf("%s", err)
      os.Exit(1)
  } else {
      defer response.Body.Close()
      contents := response.Status
      if err != nil {
          fmt.Printf("%s", err)
          os.Exit(1)
      }
      fmt.Printf("%s\n", string(contents))
  }
}
