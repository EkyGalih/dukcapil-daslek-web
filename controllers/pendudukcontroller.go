package controllers

import (
	"net/http"

	"github.com/EkyGalih/dukcapil-web/config"
	"github.com/EkyGalih/dukcapil-web/helpers"
	"github.com/EkyGalih/dukcapil-web/utils"
	"github.com/gin-gonic/gin"
)

type Penduduk struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	KeluargaID uint64 `gorm:"not null;index" json:"keluarga_id"`

	UrutanNIK        int    `json:"urutan_nik"`
	NIK              string `gorm:"index" json:"nik"`
	NamaLengkap      string `json:"nama_lengkap"`
	JenisKelamin     string `json:"jenis_kelamin"`
	TempatLahir      string `json:"tempat_lahir"`
	TanggalLahir     string `json:"tanggal_lahir"`
	Agama            string `json:"agama"`
	StatusPernikahan string `json:"status_pernikahan"`
	DudaJanda        string `json:"duda_janda"`
	GolonganDarah    string `json:"golongan_darah"`
	Pekerjaan        string `json:"pekerjaan"`

	// Relations
	KeluargaIDRef *Keluarga `gorm:"foreignKey:KeluargaID;references:ID" json:"keluarga,omitempty"`
	// Pendidikan          *Pendidikan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pendidikan,omitempty"`
	// Kesehatan           *Kesehatan  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"kesehatan,omitempty"`
	// Pendataans          []Pendataan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pendataans,omitempty"`
}

func PendudukIndex(c *gin.Context) {
	var datas []Penduduk
	err := utils.GetJSON(config.API_BASE_URL+"/penduduks", &datas)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %v", err)
		return
	}
	data := map[string]any{
		"title":    "Penduduk",
		"penduduk": datas,
	}
	helpers.RenderTemplate(c, "pages/penduduk/penduduk.html", data)
}
