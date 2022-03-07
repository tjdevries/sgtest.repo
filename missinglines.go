package main

// ExecuteFirecracker executes the firecracker process using the Go SDK
func ExecuteFirecracker(vm *api.VM, fcIfaces firecracker.NetworkInterfaces) (err error) {
	drivePath := vm.SnapshotDev()

	vCPUCount := int64(vm.Spec.CPUs)
	memSizeMib := int64(vm.Spec.Memory.MBytes())

	cmdLine := vm.Spec.Kernel.CmdLine
	if len(cmdLine) == 0 {
		// if for some reason cmdline would be unpopulated, set it to the default
		cmdLine = constants.VM_DEFAULT_KERNEL_ARGS
	}

	return nil
}
