package main

type EFI struct {
	Elements []struct {
		Type  string `json:"Type"`
		Value struct {
			FileSystemGUID struct {
				GUID string `json:"GUID"`
			} `json:"FileSystemGUID"`
			Length          int `json:"Length"`
			Signature       int `json:"Signature"`
			Attributes      int `json:"Attributes"`
			HeaderLen       int `json:"HeaderLen"`
			Checksum        int `json:"Checksum"`
			ExtHeaderOffset int `json:"ExtHeaderOffset"`
			Revision        int `json:"Revision"`
			Blocks          []struct {
				Count int `json:"Count"`
				Size  int `json:"Size"`
			} `json:"Blocks"`
			FVName struct {
				GUID string `json:"GUID"`
			} `json:"FVName"`
			ExtHeaderSize int `json:"ExtHeaderSize"`
			Files         []struct {
				Header struct {
					GUID struct {
						GUID string `json:"GUID"`
					} `json:"GUID"`
					Type       int `json:"Type"`
					Attributes int `json:"Attributes"`
					State      int `json:"State"`
				} `json:"Header"`
				Type      string `json:"Type"`
				NVarStore struct {
					Entries []struct {
						Header struct {
							Size       int `json:"Size"`
							Attributes int `json:"Attributes"`
						} `json:"Header"`
						GUID struct {
							GUID string `json:"GUID"`
						} `json:"GUID"`
						GUIDIndex int    `json:"GUIDIndex"`
						Name      string `json:"Name"`
						NVarStore struct {
							Entries []struct {
								Header struct {
									Size       int `json:"Size"`
									Attributes int `json:"Attributes"`
								} `json:"Header"`
								GUID struct {
									GUID string `json:"GUID"`
								} `json:"GUID"`
								GUIDIndex   int    `json:"GUIDIndex"`
								Name        string `json:"Name"`
								Type        int    `json:"Type"`
								Offset      int    `json:"Offset"`
								NextOffset  int    `json:"NextOffset"`
								ExtractPath string `json:"ExtractPath"`
								DataOffset  int    `json:"DataOffset"`
							} `json:"Entries"`
							GUIDStore []struct {
								GUID string `json:"GUID"`
							} `json:"GUIDStore"`
							FreeSpaceOffset int `json:"FreeSpaceOffset"`
							GUIDStoreOffset int `json:"GUIDStoreOffset"`
							Length          int `json:"Length"`
						} `json:"NVarStore"`
						Type        int    `json:"Type"`
						Offset      int    `json:"Offset"`
						NextOffset  int    `json:"NextOffset"`
						ExtractPath string `json:"ExtractPath"`
						DataOffset  int    `json:"DataOffset"`
					} `json:"Entries"`
					GUIDStore []struct {
						GUID string `json:"GUID"`
					} `json:"GUIDStore"`
					FreeSpaceOffset int `json:"FreeSpaceOffset"`
					GUIDStoreOffset int `json:"GUIDStoreOffset"`
					Length          int `json:"Length"`
				} `json:"NVarStore"`
				ExtractPath string `json:"ExtractPath"`
				DataOffset  int    `json:"DataOffset"`
			} `json:"Files"`
			DataOffset  int    `json:"DataOffset"`
			FVOffset    int    `json:"FVOffset"`
			ExtractPath string `json:"ExtractPath"`
			Resizable   bool   `json:"Resizable"`
		} `json:"Value,omitempty"`
		Value struct {
			Offset      int    `json:"Offset"`
			ExtractPath string `json:"ExtractPath"`
		} `json:"Value,omitempty"`
		Value struct {
			Offset      int    `json:"Offset"`
			ExtractPath string `json:"ExtractPath"`
		} `json:"Value,omitempty"`
	} `json:"Elements"`
	ExtractPath string      `json:"ExtractPath"`
	Length      int         `json:"Length"`
	FRegion     interface{} `json:"FRegion"`
	RegionType  int         `json:"RegionType"`
}