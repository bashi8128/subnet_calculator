package main

import (
  "log"
  "subnet_calculator/utils"

  "github.com/atotto/clipboard"

)

func main() {

  readText, err := clipboard.ReadAll()

  if err != nil {
    log.Print(err)
  }

  mySubnet := utils.CalcSubnet(readText)
  utils.CreateCalculator(mySubnet)
}
