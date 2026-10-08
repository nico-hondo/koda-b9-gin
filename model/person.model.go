package model

type Person struct {
	Id            int    `db:"id"`
	Nama          string `db:"nama"`
	Alamat        string `db:"alamat"`
	Jenis_Kelamin string `db:"jenis_kelamin"`
	Email         string `db:"email"`
	No_Telp       string `db:"no_telp"`
}
