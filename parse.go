package main

import (
	"github.com/Mimoja/MFT-Common"
	"github.com/linuxboot/fiano/pkg/uefi"
	"io/ioutil"
)

func analyse(entry MFTCommon.FlashImage) error {
	Bundle.Log.WithField("entry", entry).Infof("searching for EFI Partition inside %s\n", entry.ID.GetID())
	reader, err := Bundle.Storage.GetFile(entry.ID.GetID())
	if err != nil {
		Bundle.Log.WithField("entry", entry).WithError(err).Errorf("could not fetch file: %s : %v", entry.ID.GetID(), err)
		return err
	}
	defer reader.Close()

	defer func() {
		if r := recover(); r != nil {
			Bundle.Log.WithField("entry", entry).WithField("panic", r).Errorf("Could parse EFI. Panic while parsing!")
		}
	}()
	image, err := ioutil.ReadAll(reader)
	if err != nil {
		Bundle.Log.WithField("entry", entry).WithError(err).Error("Could not read File into buffer")
		return err
	}

	parsedRoot, err := uefi.Parse(image)
	if err != nil {
		if err.Error() == "no firmware volumes in BIOS Region" {
			Bundle.Log.WithField("entry", entry).WithError(err).Info("Not EFI found")
			return nil
		}
		Bundle.Log.WithField("entry", entry).WithError(err).Error("Could not parse UEFI")
		return err
	}

	var fileIndex uint64
	visitor := S3Extract{DirPath: entry.ID.GetID(), Index: &fileIndex}
	visitor.Run(parsedRoot)

	//jsonEncode.Encode(parsedRoot)

	id := entry.ID.GetID()
	Bundle.DB.StoreElement("efi", nil, parsedRoot, &id)

	return nil
}
