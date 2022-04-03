package utils

import (
  "strconv"
  "log"
)

/*
Interpret strings and convert them into corresponding values with type byte
*/
func AtoByte(str string) byte {
  octet, err := strconv.Atoi(str)
  if err != nil {
    log.Fatal("Failed to convert string into byte.")
  }
  return byte(octet)
}
