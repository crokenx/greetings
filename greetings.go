package greetings

import "fmt"

func Hello(name string) string {
  var message = fmt.Sprintf("Hi, %v. Welcome!", name)
  return message
}
