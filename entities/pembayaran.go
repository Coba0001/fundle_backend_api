package entities

import (
	"github.com/google/uuid"
)

type Pembayaran struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Jumlah          float64   `gorm:"type:float" json:"jumlah"`

	Transaksi          []Transaksi `gorm:"foreignKey:PembayaranID" json:"transaksi"`
	ListBankID         uint        `gorm:"type:uint" json:"list_bank_id"`
}
