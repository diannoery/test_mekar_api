package models

type Users struct {
	UserId             string     `json:"user_id"`
	Nik                string     `json:"nik"`
	Nama               string     `json:"nama"`
	TanggalLahir       string     `json:"tanggal_lahir"`
	Pekerjaan          Pekerjaan  `json:"pekerjaan"`
	PendidikanTerakhir Pendidikan `json:"pendidikan_terakhir"`
	UserStatus         string     `json:"user_status"`
	CreatedDate        string     `json:"createdDate,omitempty"`
	UpdatedDate        string     `json:"updatedDate,omitempty"`
}

type Pekerjaan struct {
	IDPekerjaan string `json:"pekerjaan_id""`
	Pekerjaan   string `json:"pekerjaan"`
}

type Pendidikan struct {
	IdPendidikan string `json:"id_pendidikan"`
	Pendidikan   string `json:"pendidikan""`
}

type LoginUser struct {
	Id_user  string `json:"id_user"`
	Username string `json:"username"`
	Password string `json:"password"`
}