package controllers

import (
	"fmt"
	"net/http"

	"github.com/EkyGalih/dukcapil-web/config"
	"github.com/EkyGalih/dukcapil-web/helpers"
	"github.com/EkyGalih/dukcapil-web/utils"
	"github.com/gin-gonic/gin"
)

type Keluarga struct {
    ID           int    `json:"id"`
    Kepala       string `json:"nama_kepala_keluarga"`
    Alamat       string `json:"alamat"`
    JumlahAnggota int   `json:"jumlah_anggota"`
}

func KeluargaIndex(c *gin.Context) {
    var datas []Keluarga
    err := utils.GetJSON(config.API_BASE_URL+"/keluargas", &datas)
    if err != nil {
        c.String(http.StatusInternalServerError, "Error: %v", err)
        return
    }
    data := map[string]any{
        "title": "Keluarga",
        "keluarga": datas,
    }
    fmt.Printf("%+v\n", datas)
    helpers.RenderTemplate(c, "pages/keluarga/keluarga.html", data)
}
