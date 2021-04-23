// Copyright (c) 2018 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"fmt"
	"os"
	"github.com/nubificus/vaccel-go-runtime/vaccel"
)

func main() {

	vaccel := vaccel.Vaccel{
		VaccelPath: "/root/olagkasn/vaccel-vsock-kata/vaccel-release/opt/",
		HostBackend: "libvaccel-noop.so",
		//guestBackend: guestback,
		SocketPath: "unix:///home/olagkasn/testvaccel.vsock",
		SocketPort: 2048,
	}
	os.Setenv("VACCEL_DEBUG_LEVEL", "4")
	fmt.Println("main: calling VaccelInit")
	fmt.Println("ENV:", os.Environ())
	vaccel.VaccelInit()
	fmt.Println("main: calling VaccelEnd")
	//vaccel.VaccelEnd()

}
