package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransaksiCreateDTO struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NamaBank            string    `json:"nama_bank" form:"nama_bank" binding:"required"`
	Jumlah_Donasi_Event float64   `gorm:"type:float" json:"jumlah_donasi" binding:"required"`
	SisaDonasi          float64   `json:"-"`
	Tanggal_Transaksi   time.Time `gorm:"timestamp with time zone" json:"tangal_transaksi" binding:"required"`

	EventID      uuid.UUID `gorm:"type:uuid" json:"event_id" form:"user_id" binding:"required"`
	PembayaranID uuid.UUID `gorm:"type:uuid" json:"pembayaran_id" binding:"required"`
	UserID       uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
}
