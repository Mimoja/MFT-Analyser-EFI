// Copyright 2018 the LinuxBoot Authors. All rights reserved
// Copyright 2018 Mimoja
// Use of this source code is governed by a BSD-style

package main

import (
	"fmt"
	"path/filepath"

	"github.com/linuxboot/fiano/pkg/uefi"
)

// S3Extract extracts any Firmware node to DirPath
type S3Extract struct {
	DirPath string
	Index   *uint64
}

// Run wraps Visit and performs some setup and teardown tasks.
func (v *S3Extract) Run(f uefi.Firmware) error {

	if err := f.Apply(v); err != nil {
		return err
	}
	return nil
}

// Visit applies the S3Extract visitor to any Firmware type.
func (v *S3Extract) Visit(f uefi.Firmware) error {
	// The visitor must be cloned before modification; otherwise, the
	// sibling's values are modified.
	v2 := *v

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

func store(buf []byte, dirPath string, filename string) (string, error) {

	fp := filepath.Join(dirPath, filename)
	//TODO do not ignore
	return fp, nil
}
