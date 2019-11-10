// Copyright 2018 the LinuxBoot Authors. All rights reserved
// Copyright 2018 Mimoja
// Use of this source code is governed by a BSD-style

package main

import (
	"encoding/json"
	"github.com/linuxboot/fiano/pkg/uefi"
)

// MFTExtract extracts any Firmware node to DirPath
type MFTExtract struct {
	JSON string
}

// Run wraps Visit and performs some setup and teardown tasks.
func (v *MFTExtract) Run(f uefi.Firmware) error {

	return f.Apply(v)
}

// Visit applies the MFTExtract visitor to any Firmware type.
func (v *MFTExtract) Visit(f uefi.Firmware) error {
	b, err := json.MarshalIndent(f, "", "\t")
	if err != nil {
		return err
	}
	v.JSON = string(b);
	return nil;
}
