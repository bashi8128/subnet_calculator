package utils

import (
  "strconv"
  "log"
  "net"
  "strings"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/data/binding"
  "fyne.io/fyne/v2/layout"
  "fyne.io/fyne/v2/widget"
)

type Subnet struct {
  Addr net.IP
  Net *net.IPNet
  BCAddr net.IP
}

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

/*
Create window of subnet calculator with information of subnet
*/
func CreateCalculator(mySubnet Subnet) {
  var IPAddrEntry, NWAddrEntry, BCAddrEntry *widget.Entry
  var SubnetEntry  *widget.SelectEntry

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

  SubnetEntry = widget.NewSelectEntry(subnets)
  if mySubnet.Net.Mask != nil {
    SubnetBound := binding.NewString()
    ones, _ := mySubnet.Net.Mask.Size()
    SubnetBound.Set(subnets[ones - 1])
    SubnetEntry.Bind(SubnetBound)
  }
  SubnetLabel := widget.NewLabel("Subnet")

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
			SubnetLabel, SubnetEntry,
			NWAddrLabel, NWAddrEntry,
			BCAddrLabel, BCAddrEntry)

  CalcButton := widget.NewButton("Execute calculation", func(){log.Println("pressed")})

  button := container.New(layout.NewCenterLayout(), CalcButton)

  wrapperContainer := container.New(layout.NewVBoxLayout(),
                                    grid,
				    button)

  myWindow.SetContent(wrapperContainer)

  myWindow.Resize(fyne.NewSize(500, 400))
  myWindow.ShowAndRun()
}

/*
Calculate informations about given subnet
*/
func CalcSubnet(str string) Subnet {
  var mySubnet Subnet

  var err error

  str = strings.TrimRight(str, "\n")

  if strings.Contains(str, "/") {
    mySubnet.Addr, mySubnet.Net, err = net.ParseCIDR(str)
    if err != nil {
      log.Fatal(err)
    }
    mySubnet.BCAddr = CalcBCAddr(mySubnet)
  } else if strings.ContainsAny(str, " '\t'") {
    splitText := strings.Fields(str)
    mySubnet.Addr = net.ParseIP(splitText[0])
    if mySubnet.Addr == nil {
      log.Fatal("Wrong IP address notation")
    }
    strMask := strings.Split(splitText[1], ".")
    mask := net.IPv4Mask(AtoByte(strMask[0]), 
                         AtoByte(strMask[1]),
                         AtoByte(strMask[2]),
                         AtoByte(strMask[3]))
    mySubnet.Net = &net.IPNet{IP: mySubnet.Addr.Mask(mask),
                              Mask: mask}
    mySubnet.BCAddr = CalcBCAddr(mySubnet)
  } else {
    log.Fatal("Invalid format in clipboard")
  }

  return mySubnet
}

/*
Calculate a broadcast address from network address and its subnet mask
*/
func CalcBCAddr(mySubnet Subnet) net.IP {
  BCAddr := make([]byte, 4)
  for i, octet := range(mySubnet.Net.IP) {
    BCAddr[i] = octet | ^mySubnet.Net.Mask[i]
  }

  return BCAddr
}

//func parseSubnet(subnet string)
