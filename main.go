package main

import (
	"github.com/awesomeProject/src"
	"github.com/gin-gonic/gin"
	"net/http"
)

type provinciasResponse struct {
	Status   string          `json:"status,omitempty"`
	Response []src.Provincia `json:"response,omitempty"`
}
type departamentosResponse struct {
	Status   string             `json:"status,omitempty"`
	Response []src.Departamento `json:"response,omitempty"`
}
type distritosResponse struct {
	Status   string         `json:"status,omitempty"`
	Response []src.Distrito `json:"response,omitempty"`
}

var provincias = []src.Provincia{
	{ID: "01", Name: "HUANUCO"},
	{ID: "02", Name: "AMBO"},
	{ID: "03", Name: "DOS DE MAYO"},
	{ID: "04", Name: "HUACAYBAMBA"},
	{ID: "05", Name: "HUAMALIES"},
	{ID: "06", Name: "LEONCIO PRADO"},
	{ID: "07", Name: "MARAÑON"},
	{ID: "08", Name: "PACHITEA"},
	{ID: "09", Name: "PUERTO INCA"},
	{ID: "10", Name: "LAURICOCHA"},
	{ID: "11", Name: "YAROWILCA"},
}

var departamentos = []src.Departamento{
	{ID: "01", Name: "SANTIAGO"},
	{ID: "02", Name: "VIÑA"},
	{ID: "03", Name: "HUINCA"},
}
var distritos = []src.Distrito{
	{ID: "01", Name: "AMBO"},
	{ID: "02", Name: "CAYNA"},
	{ID: "03", Name: "COLPAS"},
	{ID: "04", Name: "CONCHAMARCA"},
	{ID: "05", Name: "HUACAR"},
	{ID: "06", Name: "SAN FRANCISCO"},
	{ID: "07", Name: "SAN RAFAEL"},
	{ID: "08", Name: "TOMAY KICHWA"},
}
var respuestaProvincia = provinciasResponse{
	Status:   "success",
	Response: provincias,
}
var respuestaDepartamento = departamentosResponse{
	Status:   "success",
	Response: departamentos,
}
var respuestaDistrito = distritosResponse{
	Status:   "success",
	Response: distritos,
}

func main() {
	router := gin.Default()
	router.GET("/provincias", getProvincias)
	router.GET("/departamentos", getDepartamentos)
	router.GET("/distritos", getDistritos)

	router.Run("localhost:8080")
}

func getDepartamentos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, respuestaDepartamento)
}

func getDistritos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, respuestaDistrito)
}

// getAlbums responds with the list of all albums as JSON.
func getProvincias(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, respuestaProvincia)
}
