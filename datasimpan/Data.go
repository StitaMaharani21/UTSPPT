package datasimpan

type DataStruct struct {
	Nama   string `json: "Nama"`
	NIM    int    `json: "nim"`
	Alamat string `json: "alamat"`
}

func InputDataStruct() []DataStruct {
	Data := []DataStruct{
		DataStruct{
			Nama:   "Tita",
			NIM:    1234,
			Alamat: "Malang",
		},
		DataStruct{
			Nama:   "Rani",
			NIM:    5678,
			Alamat: "Kabupaten Malang",
		},
	}
	return Data
}
