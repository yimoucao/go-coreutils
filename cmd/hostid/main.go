package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

/*
In the glibc implementation, the hostid is stored in the file
/etc/hostid.  (In glibc versions before 2.2, the file /var/adm/hostid
was used.)

In the glibc implementation, if gethostid() cannot open the file
containing the host ID, then it obtains the hostname using
gethostname(2), passes that hostname to gethostbyname_r(3) in order
to obtain the host's IPv4 address, and returns a value obtained by
bit-twiddling the IPv4 address.  (This value may not be unique.)
*/

// TODO: add flags

func main() {
	if hostidFromFile() == nil {
		return
	}
	if err := hostidFromIP(); err != nil {
		panic(err)
	}
}

func hostidFromFile() error {
	f, err := os.Open("/etc/hostid")
	if err != nil {
		return err
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	fmt.Printf("%x", bs)
	return nil
}

func hostidFromIP() error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	ip, err := net.ResolveIPAddr("ip4", hostname)
	if err != nil {
		return err
	}
	ipv4 := ip.IP.To4()
	bs := []byte{ipv4[1], ipv4[0], ipv4[3], ipv4[2]}
	fmt.Printf("%x\n", bs)
	return nil
}
