package greetings

import (
  "fmt"
  "errors"
)

//Hellos retursn a message for the person calling it
func Hello(name string) (string, error) {

  if name == "" {
    return "", errors.New("empty name")
  }

  var message = fmt.Sprintf("Hi, %v. Welcome!", name)
  return message, nil
}
