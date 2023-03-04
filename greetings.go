package greetings

import (
  "fmt"
  "errors"
  "math/rand"
  "time"
)

//Hellos retursn a message for the person calling it
func Hello(name string) (string, error) {

  if name == "" {
    return "", errors.New("empty name")
  }

  var message = fmt.Sprintf(randomFormat(), name)

  return message, nil
}

// init sets initial values for variables used in the function
func init(){
  rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
  var formats = []string{
    "Hi, %v, Welcome!",
    "Great to see you, %v",
    "Hail, %v! Well met!",
    "F*ck, %v is that you?",
  }
  return formats[rand.Intn(len(formats))]
}
