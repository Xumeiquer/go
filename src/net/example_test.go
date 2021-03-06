// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net_test

import (
	"fmt"
	"io"
	"log"
	"net"
)

func ExampleListener() {
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}

func ExampleIPv4() {
	fmt.Println(net.IPv4(8, 8, 8, 8))

	// Output:
	// 8.8.8.8
}

func ExampleParseCIDR() {
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/16")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr)
	fmt.Println(ipv4Net)

	ipv6Addr, ipv6Net, err := net.ParseCIDR("2001:db8:a0b:12f0::1/32")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv6Addr)
	fmt.Println(ipv6Net)

	// Output:
	// 192.0.2.1
	// 192.0.0.0/16
	// 2001:db8:a0b:12f0::1
	// 2001:db8::/32
}

func ExampleCIDRMask() {
	// This mask corresponds to a /31 subnet for IPv4.
	fmt.Println(net.CIDRMask(31, 32))

	// This mask corresponds to a /64 subnet for IPv6.
	fmt.Println(net.CIDRMask(64, 128))

	// Output:
	// fffffffe
	// ffffffffffffffff0000000000000000
}

func ExampleIPv4Mask() {
	fmt.Println(net.IPv4Mask(255, 255, 255, 0))

	// Output:
	// ffffff00
}
