package main

import (
  "fmt"
  "net/http"
  "time"
)

var completed int = 0
var ok int = 0
var start_time time.Time


func do_connect(url string, routines int, n int, time_tracker chan int64){

  for i := 0; i < n; i++ {
    routine_time := time.Now().UnixNano()
    response, err := http.Get(url)
    if err != nil {
      fmt.Println(err)
    } else {
      defer response.Body.Close()
      status := string(response.Status)
      time_tracker <- (time.Now().UnixNano() - routine_time) / 1000000
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
    elapsed := time.Since(start_time)
    fmt.Println("")
    fmt.Println("========================")
    fmt.Println("")
    fmt.Println("Number of 200 OK: ", ok)
    fmt.Println("Run time: ", elapsed)
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
  fmt.Println("Enter number of routines: ")
  _, err := fmt.Scanln(&routines)
  handle_error(err)
  return routines
}

func get_number() int{
  var number int
  fmt.Println("Enter connections per routine: ")
  _, err := fmt.Scanln(&number)
  handle_error(err)
  return number
}

func handle_error(err error){
  if err != nil {
    fmt.Println(err)
  }
}

func gather_times(number_calls int, time_tracker chan int64, final_times chan []int64){
  times := make([]int64, 0)
  for {
    received_time := <- time_tracker
    times = append(times, received_time)
    if len(times) == number_calls{
      final_times <- times
    }
  }
}

func bubble_sort(slice []int64){
  for sliceCount := len(slice) -1; ; sliceCount-- {
    changed := false
    for i := 0; i < sliceCount; i++ {
      if slice[i] > slice[i + 1]{
        slice[i], slice[i+1] = slice[i+1], slice[i]
        changed = true
      }
    }
    if changed == false {
      break
    }
  }
}

func average_time(slice []int64) int64 {
  var sum int64 = 0
  for _,element := range slice{
    sum += element
  }
  return sum / int64((len(slice) - 1))
}

func analyze_times(time_slice []int64) (int64, int64, []int64){
  short := time_slice[0]
  long := time_slice[len(time_slice)-1]
  all := time_slice

  return short, long, all
}

func main() {
  time_tracker := make(chan int64)
  final_times := make(chan []int64)

  url := get_url()
  routines := get_routines()
  number := get_number()
  total_calls := routines * number

  fmt.Println("Running...")
  go gather_times(total_calls, time_tracker, final_times)
  start_time = time.Now()
  for i := 0; i < routines; i++ {
    go do_connect(url, routines, number, time_tracker)
  }

  all_times := <- final_times
  bubble_sort(all_times)
  shortest, longest, all := analyze_times(all_times)

  fmt.Println("Shortest Response: ", shortest, "ms")
  fmt.Println("Longest Response: ", longest, "ms")
  fmt.Println("Average Response: ", average_time(all), "ms")
}

