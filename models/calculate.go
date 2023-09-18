package models

type Calculate struct {
	Id          int64   `gorm:"primaryKey" json:"id"`
	Number      float64 `json:"number"`
	Akarkuadrat float64 `json:"akar_kuadrat"`
}