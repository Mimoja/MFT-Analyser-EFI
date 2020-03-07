package main

import (
	"encoding/json"
	"github.com/linuxboot/fiano/pkg/uefi"
)


type MFTExtract struct {
	JSON string
	Index    *uint64
}

func (v *MFTExtract) Run(f uefi.Firmware) error {
	b, err := json.MarshalIndent(f, "", "\t")
	if err != nil {
		return err
	}
	v.JSON = string(b);
	return f.Apply(v)
}

func (v *MFTExtract) Visit(f uefi.Firmware) error {
	// The visitor must be cloned before modification; otherwise, the
	// sibling's values are modified. Taken from fiano
	v2 := *v

	var err error
	switch f.(type) {

	case *uefi.FirmwareVolume:
		/*v2.DirPath = filepath.Join(v.DirPath, fmt.Sprintf("%#x", f.FVOffset))
		if len(f.Files) == 0 {
			f.ExtractPath, err = v2.extractBinary(f.Buf(), "fv.bin")
		} else {
			f.ExtractPath, err = v2.extractBinary(f.Buf()[:f.DataOffset], "fvh.bin")
		}*/

	case *uefi.File:
		/*
		// For files we use the GUID as the folder name.
		v2.DirPath = filepath.Join(v.DirPath, f.Header.GUID.String())
		// Crappy hack to make unique ids unique
		v2.DirPath = filepath.Join(v2.DirPath, fmt.Sprint(*v.Index))
		*v.Index++
		if len(f.Sections) == 0 && f.NVarStore == nil {
			f.ExtractPath, err = v2.extractBinary(f.Buf(), fmt.Sprintf("%v.ffs", f.Header.GUID))
		}
		 */

	case *uefi.Section:
		/*
		// For sections we use the file order as the folder name.
		v2.DirPath = filepath.Join(v.DirPath, fmt.Sprint(f.FileOrder))
		if len(f.Encapsulated) == 0 {
			f.ExtractPath, err = v2.extractBinary(f.Buf(), fmt.Sprintf("%v.sec", f.FileOrder))
		}
		*/


	case *uefi.NVar:
		/*// For NVar we use the GUID as the folder name the Name as file name and add the offset to links to make them unique
		v2.DirPath = filepath.Join(v.DirPath, f.GUID.String())
		if f.IsValid() {
			if f.NVarStore == nil {
				if f.Type == uefi.LinkNVarEntry {
					f.ExtractPath, err = v2.extractBinary(f.Buf()[f.DataOffset:], fmt.Sprintf("%v-%#x.bin", f.Name, f.Offset))
				} else {
					f.ExtractPath, err = v2.extractBinary(f.Buf()[f.DataOffset:], fmt.Sprintf("%v.bin", f.Name))
				}
			}
		} else {
			f.ExtractPath, err = v2.extractBinary(f.Buf(), fmt.Sprintf("%#x.nvar", f.Offset))
		}
		*/
	}
	if err != nil {
		return err
	}
	return f.ApplyChildren(&v2)
}
