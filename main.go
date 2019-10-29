package main

import (
	"MimojaFirmwareToolkit/pkg/Common"
	"encoding/json"
)

func worker(id int, file chan MFTCommon.FlashImage) {

	for true {
		load := <-file
		Bundle.Log.WithField("entry", load).Infof("Handeling %s in Worker %d\n", load.ID.GetID(), id)
		analyse(load)
	}
}

const NumberOfWorker = 3

var Bundle MFTCommon.AppBundle

func main() {
	Bundle = MFTCommon.Init("EFIUnpacker")

	biosImages := make(chan MFTCommon.FlashImage, NumberOfWorker)
	for w := 1; w <= NumberOfWorker; w++ {
		go worker(w, biosImages)
	}

	Bundle.MessageQueue.BiosImagesQueue.RegisterCallback("EFIUnpacker", func(payload string) error {

		Bundle.Log.WithField("payload", payload).Debug("Got new Message!")
		var file MFTCommon.FlashImage
		err := json.Unmarshal([]byte(payload), &file)
		if err != nil {
			Bundle.Log.WithError(err).Error("Could not unmarshall json")
			return err
		}
		biosImages <- file

		return nil
	})
	Bundle.Log.Info("Starting up!")
	select {}
}
