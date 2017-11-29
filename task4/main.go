package main

import "strings"
import "unicode"
import "fmt"


func PowerGenerator(toPower int) (func() int) {
  power := 1
  return func() int {
    power = power * toPower
    return power
  }
}


func RemoveEven(array []int) []int {
  var answer []int
  for _, curr := range array {
    if curr % 2 == 0 {
	continue
    }
    answer = append(answer, curr)
  }
  return answer
}

func DifferentWordsCount(str string) int {
  answer := make(map[string]int)
  currrent := ""
  for _, val := range str {
    if unicode.IsLetter(val) {
      currrent += strings.ToUpper(string(val))
    } else if len(currrent) != 0 {
      answer[currrent] = 1
      currrent = ""
    }
  }
  if len(currrent) != 0 {
    answer[currrent] = 1
  } 
  return len(answer)
}
