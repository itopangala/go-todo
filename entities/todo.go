package entities

type Todo struct {
	Id           int64
	Kegiatan     string `validate:"required"`
	Catatan      string `validate:"required"`
	Prioritas    string `validate:"required"`
	TenggatWaktu string `validate:"required" label:"Tenggat Waktu"`
}
