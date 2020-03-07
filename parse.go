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

	biosPart := image[entry.FirmwareOffset:]

	parsedRoot, err := uefi.Parse(biosPart)
	if err != nil {
		if err.Error() == "no firmware volumes in BIOS Region" {
			Bundle.Log.WithField("entry", entry).WithError(err).Info("Not EFI found")
			return nil
		}
		Bundle.Log.WithField("entry", entry).WithError(err).Error("Could not parse UEFI")
		return err
	}

	visitor := MFTExtract{}

	visitor.Run(parsedRoot)

	//efiJSON := visitor.JSON

	//id := entry.ID.GetID()

/*
	Bundle.DB.StoreElement("efi", nil, parsedRoot, &id)

	_, err = Bundle.DB.ES.Update().
		Index("flashimages").
		Type("flashimage").
		Id(entry.ID.GetID()).
		Doc(map[string]interface{}{"EFIBlob": efiJSON}).
		Do(context.Background())


	if err != nil {
		Bundle.Log.WithField("entry", entry).
			WithError(err).
			Errorf("Cannot update efi: %v", err)
		return err
	}

	Bundle.Log.WithField("entry", entry).
		Infof("Updated efi")
*/
	return nil
}
