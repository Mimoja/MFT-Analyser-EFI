// Copyright 2018 the LinuxBoot Authors. All rights reserved
// Copyright 2018 Mimoja
// Use of this source code is governed by a BSD-style

package main

import (
	"encoding/json"
	"fmt"
	"github.com/linuxboot/fiano/pkg/uefi"
	"path/filepath"
)

// MFTExtract extracts any Firmware node to DirPath
type MFTExtract struct {
	JSON string
}

// Run wraps Visit and performs some setup and teardown tasks.
func (v *MFTExtract) Run(f uefi.Firmware) error {

	if err := f.Apply(v); err != nil {
		return err
	}
	return nil
}

// Visit applies the MFTExtract visitor to any Firmware type.
func (v *MFTExtract) Visit(f uefi.Firmware) error {

	v2 := *v

	b, err := json.MarshalIndent(f, "", "\t")
	if err != nil {
		return err
	}
	//v.JSON = string(b);
	var err error
	switch f := f.(type) {

	case *uefi.FirmwareVolume:
		v2.DirPath = filepath.Join(v.DirPath, fmt.Sprintf("%#x", f.FVOffset))
		if len(f.Files) == 0 {
			f.ExtractPath, err = store(f.Buf(), v2.DirPath, "fv.bin")
		} else {
			f.ExtractPath, err = store(f.Buf()[:f.DataOffset], v2.DirPath, "fvh.bin")
		}

	case *uefi.File:
		// For files we use the GUID as the folder name.
		v2.DirPath = filepath.Join(v.DirPath, f.Header.GUID.String())
		// Crappy hack to make unique ids unique
		v2.DirPath = filepath.Join(v2.DirPath, fmt.Sprint(*v.Index))
		*v.Index++
		if len(f.Sections) == 0 {
			f.ExtractPath, err = store(f.Buf(), v2.DirPath, fmt.Sprintf("%v.ffs", f.Header.GUID))
		}

	case *uefi.Section:
		// For sections we use the file order as the folder name.
		v2.DirPath = filepath.Join(v.DirPath, fmt.Sprint(f.FileOrder))
		if len(f.Encapsulated) == 0 {
			f.ExtractPath, err = store(f.Buf(), v2.DirPath, fmt.Sprintf("%v.sec", f.FileOrder))
		}

	}
	if err != nil {
		return err
	}

	return f.ApplyChildren(&v2)
}
