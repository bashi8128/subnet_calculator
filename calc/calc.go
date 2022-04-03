package calc

import (
  "log"
  "net"
  "strings"

  "subnet_calculator/utils"
)

/*
Calculate informations about given subnet
*/
func CalcSubnet(str string) (net.IP, net.IP, net.IP) {
  var IPAddr, NWAddr, BCAddr net.IP
  var IPNet *net.IPNet
  var err error

  if strings.Contains(str, "/") {
    IPAddr, IPNet, err = net.ParseCIDR(str)
    if err != nil {
      log.Fatal(err)
    }
    NWAddr = IPNet.IP
    BCAddr = CalcBCAddr(NWAddr, IPNet.Mask)
  } else if strings.Contains(str, " ") {
    splitText := strings.Split(str, " ")
    IPAddr = net.ParseIP(splitText[0])
    if IPAddr == nil {
      log.Fatal("Wrong IP address notation")
    }
    strMask := strings.Split(splitText[1], ".")
    mask := net.IPv4Mask(utils.AtoByte(strMask[0]), 
                         utils.AtoByte(strMask[0]),
                         utils.AtoByte(strMask[2]),
                         utils.AtoByte(strMask[3]))
    NWAddr = IPAddr.Mask(mask)
    BCAddr = CalcBCAddr(NWAddr, mask)
  } else {
  }

  return IPAddr, NWAddr, BCAddr

}

/*
Calculate a broadcast address
*/
func CalcBCAddr(addr net.IP, mask net.IPMask) net.IP {
  BCAddr := make([]byte, 4)
  copy(BCAddr, addr)
  for i, octet := range(addr) {
    BCAddr[i] = octet | ^mask[i]
  }

  return BCAddr
}
