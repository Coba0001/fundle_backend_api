package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transaksi struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Jumlah_Donasi_Event float64   `gorm:"type:float" json:"jumlah_donasi"`
	Tanggal_Transaksi   time.Time `gorm:"datetime" json:"tangal_transaksi"`

	HistoryTransaksiUser HistoryTransaksiUser `gorm:"foreignKey:TransaksiID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"history_transaksi_users,omitempty"`
	EventID              uuid.UUID            `gorm:"type:uuid" json:"event_id"`
	Event                Event                `gorm:"foreignKey:EventID" json:"event"`
	Pembayaran           []Pembayaran         `gorm:"foreignKey:TransaksiID" json:"pembayaran"`
	UserID               uuid.UUID            `gorm:"type:uuid" json:"user_id"`
	User                 User                 `gorm:"foreignKey:UserID" json:"user"`
}