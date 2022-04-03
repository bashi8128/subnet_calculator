package calc

import (
  "log"
  "net"
  "strings"

  "subnet_calculator/utils"
)

type Subnet struct {
  Addr net.IP
  Net *net.IPNet
  BCAddr net.IP
}

/*
Calculate informations about given subnet
*/
func CalcSubnet(str string) Subnet {
  var mySubnet Subnet

  var err error

  if strings.Contains(str, "/") {
    mySubnet.Addr, mySubnet.Net, err = net.ParseCIDR(str)
    if err != nil {
      log.Fatal(err)
    }
    mySubnet.BCAddr = CalcBCAddr(mySubnet)
  } else if strings.Contains(str, " ") {
    splitText := strings.Split(str, " ")
    mySubnet.Addr = net.ParseIP(splitText[0])
    if mySubnet.Addr == nil {
      log.Fatal("Wrong IP address notation")
    }
    strMask := strings.Split(splitText[1], ".")
    mask := net.IPv4Mask(utils.AtoByte(strMask[0]), 
                         utils.AtoByte(strMask[0]),
                         utils.AtoByte(strMask[2]),
                         utils.AtoByte(strMask[3]))
    mySubnet.Net = &net.IPNet{IP: mySubnet.Addr.Mask(mask),
                              Mask: mask}
    mySubnet.BCAddr = CalcBCAddr(mySubnet)
  } else {
  }

  return mySubnet

}

/*
Calculate a broadcast address from network address and its subnet mask
*/
func CalcBCAddr(mySubnet Subnet) net.IP {
  BCAddr := make([]byte, 4)
  copy(BCAddr, mySubnet.Net.IP)
  for i, octet := range(mySubnet.Net.IP) {
    BCAddr[i] = octet | ^mySubnet.Net.Mask[i]
  }

  return BCAddr
}
