package dto

type Person struct {
	Id            int    `json:"id"`
	Nama          string `json:"nama"`
	Alamat        string `json:"alamat"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	Email         string `json:"email"`
	No_Telp       string `json:"no_telp"`
}
