package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (a IPAddr) String() string {
  var str string
  for _, a := range a {
    str += fmt.Sprintf("%v.", a)
  }
  return str[: len(str) - 1]
}

func main() {
  hosts := map[string]IPAddr{
    "loopback":  {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }
  for name, ip := range hosts {
    fmt.Printf("%v: %v\n", name, ip)
  }
}
