package main

import (
  "fmt"
  "net/http"
  "time"
  "strings"
  "math"
  "github.com/ansi"
)

var completed int = 0
var start_time time.Time
var response_codes = make(map[string]int)


func do_connect(url string, n int, time_tracker chan int64){
  for i := 0; i < n; i++ {
    routine_time := time.Now().UnixNano()
    response, err := http.Get(url)
    if err != nil {
      fmt.Println(err)
    } else {
      defer response.Body.Close()
      time_tracker <- (time.Now().UnixNano() - routine_time) / 1000000
      status := string(response.Status)
      if response_codes[status] == 0{
        response_codes[status] = 1
      } else {
        response_codes[status] = response_codes[status] + 1
      }
    }
  }

  completed++
  fmt.Print("+")

}

func get_url() string{
  test_url := "http://random-responder.herokuapp.com/random"

  var url string
  fmt.Println("Enter URL (http://host.com/path):")
  _, err := fmt.Scanln(&url)
  if err != nil && err.Error() != "unexpected newline" {
    fmt.Println(err)
  }
  if url == ""{
    url = test_url
  }
  if !strings.HasPrefix(url, "http://"){
    if !strings.HasPrefix(url, "https://"){
      url = "http://" + url
    }
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
  for _, element := range slice{
    sum += element
  }
  return sum / int64((len(slice) - 1))
}

func standard_deviation(slice []int64, average int64) (float64) {
  length := len(slice) - 1
  var sum int64
  var variance float64
  var ssd float64

  for _, element := range slice{
    delta := element - average
    sum += delta * delta
  }
  variance = float64(sum) / float64(length)
  ssd = math.Sqrt(variance)

  return ssd
}

func analyze_times(time_slice []int64) (int64, int64, int64, float64){
  short := time_slice[0]
  long := time_slice[len(time_slice)-1]
  average :=  average_time(time_slice)
  stddev := standard_deviation(time_slice, average)

  return short, long, average, stddev
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
    go do_connect(url, number, time_tracker)
  }

  all_times := <- final_times
  bubble_sort(all_times)
  shortest, longest, average, stddev := analyze_times(all_times)
  elapsed := time.Since(start_time)

  green := ansi.ColorCode("green")
  yellow := ansi.ColorCode("yellow")
  red := ansi.ColorCode("red")
  reset := ansi.ColorCode("reset")

  fmt.Println("")
  fmt.Println("========================")
  fmt.Println("")
  fmt.Println("Run time: ", elapsed)
  fmt.Println("Response Codes:")
  for key,value := range response_codes {
    if (strings.HasPrefix(key, "2")){
      fmt.Println("  *", green, strings.TrimSpace(key), reset, "-> ", value, "of", total_calls)
    } else if (strings.HasPrefix(key, "3") || strings.HasPrefix(key, "4")){
      fmt.Println("  *", yellow, strings.TrimSpace(key), reset, "-> ", value, "of", total_calls)
    } else if (strings.HasPrefix(key, "5")){
      fmt.Println("  *", red, strings.TrimSpace(key), reset, "-> ", value, "of", total_calls)
    }else{
    fmt.Println("  *", strings.TrimSpace(key), "-> ", value, "of", total_calls)
    }
  }

  fmt.Println("Shortest Response: ", shortest, "ms")
  fmt.Println("Longest Response: ", longest, "ms")
  fmt.Println("Average Response: ", average, "ms")
  fmt.Println("Standard Deviation: ", stddev)
}

