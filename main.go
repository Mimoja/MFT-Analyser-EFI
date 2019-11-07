package main

import (
	"github.com/Mimoja/MFT-Common"
	"encoding/json"
)

var Bundle MFTCommon.AppBundle

func main() {
	Bundle = MFTCommon.Init("EFIUnpacker")

	Bundle.MessageQueue.BiosImagesQueue.RegisterCallback("EFIUnpacker", func(payload string) error {

		Bundle.Log.WithField("payload", payload).Debug("Got new Message!")
		var file MFTCommon.FlashImage
		err := json.Unmarshal([]byte(payload), &file)
		if err != nil {
			Bundle.Log.WithError(err).Error("Could not unmarshall json")
			return err
		}
		Bundle.Log.WithField("entry", file).Infof("Handeling %s\n", file.ID.GetID())
		analyse(file)

		return nil
	})
	Bundle.Log.Info("Starting up!")
	select {}
}
