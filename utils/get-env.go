package utils

import (
  "os"
  "fmt"
)

func GetEnv(variable string) string{
  // Set Environment Variables
  if os.Getenv(variable)==""{
    fmt.Println("Environment variable not found: "+variable)
  }
  return os.Getenv(variable)
}
