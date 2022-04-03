package main

import (
  "log"
  "subnet_calculator/calc"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/data/binding"
  "fyne.io/fyne/v2/widget"
  "github.com/atotto/clipboard"

)

func main() {
  var IPAddrEntry, NWAddrEntry, BCAddrEntry *widget.Entry

  readText, err := clipboard.ReadAll()

  if err != nil {
    log.Print(err)
  }

  IPAddr, NWAddr, BCAddr := calc.CalcSubnet(readText)

  myApp := app.New()
  myWindow := myApp.NewWindow("Subnet Calculator")

  IPAddrBound := binding.NewString()
  if IPAddr != nil {
    IPAddrBound.Set(IPAddr.String())
    IPAddrEntry = widget.NewEntryWithData(IPAddrBound)
  } else {
    IPAddrEntry = widget.NewEntry()
  }
  IPAddrContainer := container.NewHBox(
    widget.NewLabel("IP Address"),
    IPAddrEntry,
  )

  NWAddrBound := binding.NewString()
  if NWAddr != nil {
    NWAddrBound.Set(NWAddr.String())
    NWAddrEntry = widget.NewEntryWithData(NWAddrBound)
  } else {
    NWAddrEntry = widget.NewEntry()
  }
  NWAddrContainer := container.NewHBox(
    widget.NewLabel("Network Address"),
    NWAddrEntry,
  )

  BCAddrBound := binding.NewString()
  if BCAddr != nil {
    BCAddrBound.Set(BCAddr.String())
    BCAddrEntry = widget.NewEntryWithData(BCAddrBound)
  } else {
    BCAddrEntry = widget.NewEntry()
  }
  BCAddrContainer := container.NewHBox(
    widget.NewLabel("Network Address"),
    BCAddrEntry,
  )

  myWindow.SetContent(container.NewVBox(
    IPAddrContainer,
    NWAddrContainer,
    BCAddrContainer,
  ))

  myWindow.Resize(fyne.NewSize(500, 100))
  myWindow.ShowAndRun()
}
