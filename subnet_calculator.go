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

  mySubnet := calc.CalcSubnet(readText)

  myApp := app.New()
  myWindow := myApp.NewWindow("Subnet Calculator")

  IPAddrBound := binding.NewString()
  if mySubnet.Addr != nil {
    IPAddrBound.Set(mySubnet.Addr.String())
    IPAddrEntry = widget.NewEntryWithData(IPAddrBound)
  } else {
    IPAddrEntry = widget.NewEntry()
  }
  IPAddrLabel := widget.NewLabel("IP Address")

  NWAddrBound := binding.NewString()
  if mySubnet.Net.IP != nil {
    NWAddrBound.Set(mySubnet.Net.IP.String())
    NWAddrEntry = widget.NewEntryWithData(NWAddrBound)
  } else {
    NWAddrEntry = widget.NewEntry()
  }
  NWAddrLabel := widget.NewLabel("Network Address")

  BCAddrBound := binding.NewString()
  if mySubnet.BCAddr != nil {
    BCAddrBound.Set(mySubnet.BCAddr.String())
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
