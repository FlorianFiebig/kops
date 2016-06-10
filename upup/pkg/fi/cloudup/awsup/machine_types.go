package awsup

import (
	"fmt"
	"github.com/golang/glog"
)

// I believe one vCPU ~ 3 ECUS, and 60 CPU credits would be needed to use one vCPU for an hour
const BurstableCreditsToECUS float32 = 3.0 / 60.0

type AWSMachineTypeInfo struct {
	Name           string
	MemoryGB       float32
	ECU            float32
	Cores          int
	EphemeralDisks []int
	Burstable      bool
}

type EphemeralDevice struct {
	DeviceName  string
	VirtualName string
	SizeGB      int
}

func (m *AWSMachineTypeInfo) EphemeralDevices() ([]*EphemeralDevice) {
	var disks []*EphemeralDevice
	for i, sizeGB := range m.EphemeralDisks {
		d := &EphemeralDevice{
			SizeGB: sizeGB,
		}

		if i >= 20 {
			// TODO: What drive letters do we use?
			glog.Fatalf("ephemeral devices for > 20 not yet implemented")
		}
		d.DeviceName = "/dev/sd" + string('c'+i)
		d.VirtualName = fmt.Sprintf("ephemeral%d", i)

		disks = append(disks, d)
	}
	return disks
}

func GetMachineTypeInfo(machineType string) (*AWSMachineTypeInfo, error) {
	for i := range MachineTypes {
		m := &MachineTypes[i]
		if m.Name == machineType {
			return m, nil
		}
	}

	return nil, fmt.Errorf("instance type not handled: %q", machineType)
}

var MachineTypes []AWSMachineTypeInfo = []AWSMachineTypeInfo{
	// This is tedious, but seems simpler than trying to have some logic and then a lot of exceptions

	// t2 family
	{
		Name:           "t2.nano",
		MemoryGB:       0.5,
		ECU:            3 * BurstableCreditsToECUS,
		Cores:          1,
		EphemeralDisks: nil,
		Burstable:      true,
	},
	{
		Name:           "t2.micro",
		MemoryGB:       1,
		ECU:            6 * BurstableCreditsToECUS,
		Cores:          1,
		EphemeralDisks: nil,
		Burstable:      true,
	},
	{
		Name:           "t2.small",
		MemoryGB:       2,
		ECU:            12 * BurstableCreditsToECUS,
		Cores:          1,
		EphemeralDisks: nil,
		Burstable:      true,
	},
	{
		Name:           "t2.medium",
		MemoryGB:       4,
		ECU:            24 * BurstableCreditsToECUS,
		Cores:          2,
		EphemeralDisks: nil,
		Burstable:      true,
	},
	{
		Name:           "t2.large",
		MemoryGB:       8,
		ECU:            36 * BurstableCreditsToECUS,
		Cores:          2,
		EphemeralDisks: nil,
		Burstable:      true,
	},

	// m3 family
	{
		Name:           "m3.medium",
		MemoryGB:       3.75,
		ECU:            3,
		Cores:          1,
		EphemeralDisks: []int{4},
	},
	{
		Name:           "m3.large",
		MemoryGB:       7.5,
		ECU:            6.5,
		Cores:          2,
		EphemeralDisks: []int{32},
	},
	{
		Name:           "m3.xlarge",
		MemoryGB:       15,
		ECU:            13,
		Cores:          4,
		EphemeralDisks: []int{40, 40},
	},
	{
		Name:           "m3.2xlarge",
		MemoryGB:       30,
		ECU:            26,
		Cores:          8,
		EphemeralDisks: []int{80, 80},
	},

	// m4 family
	{
		Name:           "m4.large",
		MemoryGB:       8,
		ECU:            6.5,
		Cores:          2,
		EphemeralDisks: nil,
	},
	{
		Name:           "m4.xlarge",
		MemoryGB:       16,
		ECU:            13,
		Cores:          4,
		EphemeralDisks: nil,
	},
	{
		Name:           "m4.2xlarge",
		MemoryGB:       32,
		ECU:            26,
		Cores:          8,
		EphemeralDisks: nil,
	},
	{
		Name:           "m4.4xlarge",
		MemoryGB:       64,
		ECU:            53.5,
		Cores:          16,
		EphemeralDisks: nil,
	},
	{
		Name:           "m4.10xlarge",
		MemoryGB:       160,
		ECU:            124.5,
		Cores:          40,
		EphemeralDisks: nil,
	},

	// c3 family
	{
		Name:           "c3.large",
		MemoryGB:       3.75,
		ECU:            7,
		Cores:          2,
		EphemeralDisks: []int{16, 16},
	},
	{
		Name:           "c3.xlarge",
		MemoryGB:       7.5,
		ECU:            14,
		Cores:          4,
		EphemeralDisks: []int{40, 40},
	},
	{
		Name:           "c3.2xlarge",
		MemoryGB:       15,
		ECU:            28,
		Cores:          8,
		EphemeralDisks: []int{80, 80},
	},
	{
		Name:           "c3.4xlarge",
		MemoryGB:       30,
		ECU:            55,
		Cores:          16,
		EphemeralDisks: []int{160, 160},
	},
	{
		Name:           "c3.8xlarge",
		MemoryGB:       60,
		ECU:            108,
		Cores:          32,
		EphemeralDisks: []int{320, 320},
	},

	// c4 family
	{
		Name:           "c4.large",
		MemoryGB:       3.75,
		ECU:            8,
		Cores:          2,
		EphemeralDisks: nil,
	},
	{
		Name:           "c4.xlarge",
		MemoryGB:       7.5,
		ECU:            16,
		Cores:          4,
		EphemeralDisks: nil,
	},
	{
		Name:           "c4.2xlarge",
		MemoryGB:       15,
		ECU:            31,
		Cores:          8,
		EphemeralDisks: nil,
	},
	{
		Name:           "c4.4xlarge",
		MemoryGB:       30,
		ECU:            62,
		Cores:          16,
		EphemeralDisks: nil,
	},
	{
		Name:           "c4.8xlarge",
		MemoryGB:       60,
		ECU:            132,
		Cores:          32,
		EphemeralDisks: nil,
	},
}
