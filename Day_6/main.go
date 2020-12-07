package main

import (
  "fmt"
  "strings"
  "io/ioutil"
)

func removeDuplicate(answer string) int {
  s := ""
  for _, v := range answer {
    if strings.Contains(s, string(v)) {
      
    } else {
      s = s+string(v)
    }
  }
  return len(s)
}

func returnString(answer string) string {
  s := ""
  for _, v := range answer {
    if strings.Contains(s, string(v)) {
      
    } else {
      s = s+string(v)
    }
  }
  return s
}

func remove(answer string) {
  complete := ""
  fixed := returnString(answer)
  complete = complete + fixed
}

func secondPart(groups []string) {
  total := 0

  for _,v := range groups {
    group := strings.Split(v, "\n")
   
   // for single element containg slice
    if len(group) == 1 {
      for _, y := range group {
      sum := removeDuplicate(y)
      total = total + sum
      }
    } else {
      complete := ""
      for _, y := range group {
        fixed := returnString(y)
        complete = complete + fixed
      }
      original := ""
      for _, v := range complete{
        count := strings.Count(complete, string(v))
        if count == len(group){
          original = original + string(v)
        }
      }
      answer := removeDuplicate(original)
      total = total + answer
    }
  }
  fmt.Println("Second Part answer is := ", total)
}

func main() {
  content, err := ioutil.ReadFile("input.txt")
  if err != nil {
    fmt.Println(err)
  }

  var total int

  groups := strings.Split(string(content), "\n\n")
  for _,v := range groups {
    group := strings.Split(v, "\n")
    answer := strings.Join(group, "")

    individual := removeDuplicate(answer)
    total = total + individual
  }
  fmt.Println("First part answer is := ",total)

  secondPart(groups)
}