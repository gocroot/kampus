package riziq

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tabel Pelanggan
type Pelanggan struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Alamat    string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	NoTelepon string             `bson:"no_telepon,omitempty" json:"no_telepon,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
}

// Tabel Tagihan
type Tagihan struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	IDPelanggan      primitive.ObjectID `bson:"id_pelanggan,omitempty" json:"id_pelanggan,omitempty"`
	TanggalTagihan   string             `bson:"tanggal_tagihan,omitempty" json:"tanggal_tagihan,omitempty"`
	TotalTagihan     string             `bson:"total_tagihan,omitempty" json:"total_tagihan,omitempty"`
	StatusPembayaran string             `bson:"status_pembayaran,omitempty" json:"status_pembayaran,omitempty"`
}

// Tabel Pembayaran
type Pembayaran struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	IDTagihan         primitive.ObjectID `bson:"id_tagihan,omitempty" json:"id_tagihan,omitempty"`
	TanggalPembayaran string             `bson:"tanggal_pembayaran,omitempty" json:"tanggal_pembayaran,omitempty"`
	JumlahPembayaran  string             `bson:"jumlah_pembayaran,omitempty" json:"jumlah_pembayaran,omitempty"`
	MetodePembayaran  string             `bson:"metode_pembayaran,omitempty" json:"metode_pembayaran,omitempty"`
}

// Tabel Produk
type Produk struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NamaProduk  string             `bson:"nama_produk,omitempty" json:"nama_produk,omitempty"`
	HargaProduk string             `bson:"harga_produk,omitempty" json:"harga_produk,omitempty"`
}

// Tabel About
type About struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	IsiSatu string             `bson:"isisatu,omitempty" json:"isisatu,omitempty"`
	IsiDua  string             `bson:"isidua,omitempty" json:"isidua,omitempty"`
	Image   string             `bson:"image,omitempty" json:"image,omitempty"`
}
