package main

import (
  "log"
  "subnet_calculator/calc"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/data/binding"
  "fyne.io/fyne/v2/layout"
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
  IPAddrLabel := widget.NewLabel("IP Address")

  NWAddrBound := binding.NewString()
  if NWAddr != nil {
    NWAddrBound.Set(NWAddr.String())
    NWAddrEntry = widget.NewEntryWithData(NWAddrBound)
  } else {
    NWAddrEntry = widget.NewEntry()
  }
  NWAddrLabel := widget.NewLabel("Network Address")

  BCAddrBound := binding.NewString()
  if BCAddr != nil {
    BCAddrBound.Set(BCAddr.String())
    BCAddrEntry = widget.NewEntryWithData(BCAddrBound)
  } else {
    BCAddrEntry = widget.NewEntry()
  }
  BCAddrLabel := widget.NewLabel("Broadcast Address")

  grid := container.New(layout.NewFormLayout(), 
                        IPAddrLabel, IPAddrEntry,
			NWAddrLabel, NWAddrEntry,
			BCAddrLabel, BCAddrEntry)
  myWindow.SetContent(grid)

  myWindow.Resize(fyne.NewSize(500, 100))
  myWindow.ShowAndRun()
}
