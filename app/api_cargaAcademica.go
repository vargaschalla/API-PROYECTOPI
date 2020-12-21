package app

import (
	"PROYECTintegrador/ProyectoGOI/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CargaAcademicaIndex(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	lis := []models.CargaAcademica{}
	conn.Find(&lis)
	conn.Preload("Academic_Plane").Preload("Academic_Period").Preload("Alumno").Preload("Seccion").Find(&lis) // Preload("Alumno") carga los objetos Alumno relacionado
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})

}

func CargaAcademicaGETID(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.CargaAcademica
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &d)

}

func CargaAcademicaCreate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)
	var d models.CargaAcademica

	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)

}

func CargaAcademicaUpdate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.CargaAcademica
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&d)
	conn.Save(&d)
	c.JSON(http.StatusOK, &d)

}

func CargaAcademicaDelete(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.CargaAcademica

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
