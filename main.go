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
		VaccelPath: "/home/olagkasn/vaccel_featkata_Release",
		HostBackends: "noop,jetson,plugin1",
		//guestBackend: guestback,
		SocketPath: "unix:///home/olagkasn/testvaccel.vsock",
		SocketPort: 2048,
	}
	os.Setenv("VACCEL_DEBUG_LEVEL", "4")
	fmt.Println("ENV:", vaccel.VaccelEnv())
	fmt.Println("main: calling VaccelInit")
	vaccel.VaccelInit()
	fmt.Println("ENV:", vaccel.VaccelEnv())
	fmt.Println("main: calling VaccelEnd")
	//vaccel.VaccelEnd()

}
