package main

import (
  "log"
  "subnet_calculator/calc"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/widget"
  "github.com/atotto/clipboard"

)

func main() {
  readText, err := clipboard.ReadAll()

  if err != nil {
    log.Print(err)
  }

  IPAddr, NWAddr, BCAddr := calc.CalcSubnet(readText)

  myApp := app.New()
  myWindow := myApp.NewWindow("Subnet Calculator")

  myWindow.SetContent(container.NewVBox(
    widget.NewLabel("IP Address: " + IPAddr.String()),
    widget.NewLabel("Network Address: " + NWAddr.String()),
    widget.NewLabel("Broadcast Address: " + BCAddr.String()),
  ))

  myWindow.Resize(fyne.NewSize(500, 100))
  myWindow.ShowAndRun()
}
