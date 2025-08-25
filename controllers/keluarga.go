package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.gom/EkyGalih/dukcapil-web/config"
	"github.gom/EkyGalih/dukcapil-web/utils"
)

type Keluarga struct {
    ID           int    `json:"id"`
    Kepala       string `json:"kepala"`
    Alamat       string `json:"alamat"`
    JumlahAnggota int   `json:"jumlah_anggota"`
}

func KeluargaIndex(c *gin.Context) {
    var data []Keluarga
    err := utils.GetJSON(config.API_BASE_URL+"/keluargas", &data)
    if err != nil {
        c.String(http.StatusInternalServerError, "Error: %v", err)
        return
    }

    c.HTML(http.StatusOK, "pages/keluarga/keluarga.html", gin.H{
        "title": "Data Keluarga",
        "keluargas": data,
    })
}
