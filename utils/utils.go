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
    log.Print("Failed to convert string into byte.")
  }
  return byte(octet)
}

/*
Create window of subnet calculator with information of subnet
*/
func CreateCalculator(mySubnet Subnet) {
  var IPAddrEntry, NWAddrEntry, BCAddrEntry *widget.Entry
  var SubnetEntry  *widget.SelectEntry = widget.NewSelectEntry(subnets)

  myApp := app.New()
  myWindow := myApp.NewWindow("Subnet Calculator")

  // Define labels for forms
  IPAddrLabel := widget.NewLabel("IP Address")
  SubnetLabel := widget.NewLabel("Subnet")
  NWAddrLabel := widget.NewLabel("Network Address")
  BCAddrLabel := widget.NewLabel("Broadcast Address")

  // Define binding for each forms
  IPAddrBound := binding.NewString()
  SubnetBound := binding.NewString()
  NWAddrBound := binding.NewString()
  BCAddrBound := binding.NewString()

  if mySubnet.Addr != nil {
    IPAddrBound.Set(mySubnet.Addr.String())
  }
  IPAddrEntry = widget.NewEntryWithData(IPAddrBound)
  
  if mySubnet.Net != nil {
    ones, _ := mySubnet.Net.Mask.Size()
    SubnetBound.Set(subnets[ones - 1])
  }
  SubnetEntry.Bind(SubnetBound)

  if mySubnet.Net != nil {
    NWAddrBound.Set(mySubnet.Net.IP.String())
  }
  NWAddrEntry = widget.NewEntryWithData(NWAddrBound)

  if mySubnet.BCAddr != nil {
    BCAddrBound.Set(mySubnet.BCAddr.String())
  }
  BCAddrEntry = widget.NewEntryWithData(BCAddrBound)

  grid := container.New(layout.NewFormLayout(), 
                        IPAddrLabel, IPAddrEntry,
			SubnetLabel, SubnetEntry,
			NWAddrLabel, NWAddrEntry,
			BCAddrLabel, BCAddrEntry)

  CalcButton := widget.NewButton("Execute calculation",
                                 func(){
				   var IPAddr, Subnet string
				   var err error
				   IPAddr, err = IPAddrBound.Get()
				   if err != nil {
				     log.Print(err)
				   }

				   Subnet, err = SubnetBound.Get()
				   if err != nil {
				     log.Print(err)
				   }
				   mask := ExtractMask(IPAddr + "/" + Subnet)
				   mySubnet := CalcSubnet(mask)
				   NWAddrBound.Set(mySubnet.Net.IP.String())
				   BCAddrBound.Set(mySubnet.BCAddr.String())
                                 })

  button := container.New(layout.NewCenterLayout(), CalcButton)

  wrapperContainer := container.New(layout.NewVBoxLayout(),
                                    grid,
				    button)

  myWindow.SetContent(wrapperContainer)

  myWindow.Resize(fyne.NewSize(500, 200))
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
      log.Print(err)
    }
    mySubnet.BCAddr = CalcBCAddr(mySubnet)
  } else if strings.ContainsAny(str, " '\t'") {
    splitText := strings.Fields(str)
    mySubnet.Addr = net.ParseIP(splitText[0])
    if mySubnet.Addr == nil {
      log.Print("Wrong IP address notation")
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
    log.Print("Invalid format in clipboard")
  }

  return mySubnet
}

/*
Extract bit length of subnet mask from string like 24(255.255.255.0) into IPMask
*/
func ExtractMask(subnet string) string {
  SubnetString := strings.Split(subnet, "(")
  return SubnetString[0]
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
