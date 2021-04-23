// Copyright (c) 2018 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package vaccel

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
	"strconv"
)



// vAccelInfo contains information related to the vAccelrt agent
type vAccelInfo struct {
	PID     int
	Version string
}


type vAccelType int


// firecracker is an Hypervisor interface implementation for the firecracker VMM.
type Vaccel struct {
	VaccelPath    string //Path to vaccel installation on host
	HostBackend   string //Vaccel backend framework
	GuestBackend  string //Vaccel transport layer (vsock or virtio)
	SocketPath    string //vsock specific.. move this to guestBackend
	SocketPort    uint32 //vsock specific.. move this to guestBackend

	info vAccelInfo //vaccelrt-agent info, also vsock specific

	vaccelrtd *exec.Cmd           //Tracks the vaccelrt-agent, vsock specific
}

// This is the vsock implementation (will be renamed to vaccelrtAgent
func (vaccel *Vaccel) VaccelInit() error {

	var cmd *exec.Cmd
	var args []string

	// Create the right environment for the vaccelrt-agent
	vaccelrtBin := filepath.Join(vaccel.VaccelPath, "bin", "vaccelrt-agent")
	vaccelrtLibs := filepath.Join(vaccel.VaccelPath, "lib")
	vaccelrtBack := filepath.Join(vaccelrtLibs, vaccel.HostBackend)
	vaccel_backends := "VACCEL_BACKENDS=" + vaccelrtBack
	vaccel_debug := "VACCEL_DEBUG_LEVEL=" + "4"
	ld_path := "LD_LIBRARY_PATH=" + vaccelrtLibs

	server_address := "unix://" + vaccel.SocketPath + "_" + strconv.FormatUint(uint64(vaccel.SocketPort), 10)
	args = append(args, "--server-address", server_address)

	cmd = exec.Command(vaccelrtBin, args...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, vaccel_backends, ld_path)

	if err := cmd.Start(); err != nil {
		return err
	}

	vaccel.info.PID = cmd.Process.Pid
	fmt.Println("Started wih PID:", vaccel.info.PID)
	vaccel.vaccelrtd = cmd
	return nil
}

func (vaccel *Vaccel) VaccelEnd() (err error) {

	pid := vaccel.info.PID

	// Send a SIGTERM to the vAccel agent to try to stop it properly
	if err = syscall.Kill(pid, syscall.SIGTERM); err != nil {
		if err == syscall.ESRCH {
			return nil
		}
		return err
	}

	// Wait for the vAccel process to terminate
	tInit := time.Now()
	for {
		if err = syscall.Kill(pid, syscall.Signal(0)); err != nil {
			return nil
		}

		if time.Since(tInit).Seconds() >= 5 {
			break
		}

		// Let's avoid to run a too busy loop
		time.Sleep(time.Duration(50) * time.Millisecond)
	}

	// Let's try with a hammer now, a SIGKILL should get rid of the
	// VM process.
	return syscall.Kill(pid, syscall.SIGKILL)
}
