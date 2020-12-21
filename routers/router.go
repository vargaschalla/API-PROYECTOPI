package routers

import (
	"PROYECTintegrador/ProyectoGOI/app"
	"PROYECTintegrador/ProyectoGOI/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {

	conn, err := connectDBmysql()
	if err != nil {
		panic("failed to connect database: " + err.Error())
		//return
	}
	// Migrate the schema
	conn.AutoMigrate(
		&models.Person{},
		&models.User{},
		&models.Rol{},
		&models.Sesiones{},
		&models.SesionActividad{},
		&models.Nivel{},
		&models.Grado{},
		&models.Curso{},
		&models.Alumno{},
		&models.Docente{},
		&models.Academic_Period{},
		&models.Academic_Plane{},
		&models.SesionMaterial{},
		&models.Seccion{},
		&models.MaterialActividad{},
		&models.CargaAcademica{},
		&models.CategoriaCurso{},
		&models.Modulo{},
	)

	r := gin.Default()

	//config := cors.DefaultConfig() https://github.com/rs/cors
	//config.AllowOrigins = []string{"http://localhost", "http://localhost:8086"}

	r.Use(CORSMiddleware())

	r.Use(dbMiddleware(*conn))

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", app.ItemsIndex)

		v1.GET("/personas", app.PersonsIndex)
		v1.POST("/personas", authMiddleWare(), app.PersonsCreate)
		v1.GET("/personas/:id", app.PersonsGet)
		v1.PUT("/personas/:id", app.PersonsUpdate)
		v1.DELETE("/personas/:id", app.PersonsDelete)

		v1.GET("/users", app.UsersIndex)
		v1.POST("/users", app.UsersCreate)
		v1.GET("/users/:id", app.UsersGet)
		v1.PUT("/users/:id", app.UsersUpdate)
		v1.DELETE("/users/:id", app.UsersDelete)
		v1.POST("/login", app.UsersLogin)
		v1.POST("/logout", app.UsersLogout)

		v1.GET("/rol", app.RolLista)
		v1.POST("/rol", authMiddleWare(), app.RolCreate)
		v1.GET("/rol/:id", app.RolGetID)
		v1.PUT("/rol/:id", app.RolUpdate)
		v1.DELETE("/rol/:id", app.RolDelete)

		v1.GET("/sesiones", app.SesionIndex)
		v1.POST("/sesiones", authMiddleWare(), app.SesionCreate)
		v1.GET("/sesiones/:id", app.SesionGet)
		v1.PUT("/sesiones/:id", app.SesionUpdate)
		v1.DELETE("/sesiones/:id", app.SesionDelete)

		v1.GET("/sesionmaterial", app.SesionMaterialIndex)
		v1.POST("/sesionmaterial", authMiddleWare(), app.SesionMaterialCreate)
		v1.GET("/sesionmaterial/:id", app.SesionMaterialGet)
		v1.PUT("/sesionmaterial/:id", app.SesionMaterialUpdate)
		v1.DELETE("/sesionmaterial/:id", app.SesionMaterialDelete)

		v1.GET("/nivel", app.NivelIndex)
		v1.POST("/nivel", authMiddleWare(), app.NivelCreate)
		v1.GET("/nivel/:id", app.NivelGet)
		v1.PUT("/nivel/:id", app.NivelUpdate)
		v1.DELETE("/nivel/:id", app.NivelDelete)

		v1.GET("/grado", app.GradoIndex)
		v1.POST("/grado", authMiddleWare(), app.GradoCreate)
		v1.GET("/grado/:id", app.GradoGet)
		v1.PUT("/grado/:id", app.GradoUpdate)
		v1.DELETE("/grado/:id", app.GradoDelete)

		v1.GET("/alumno", app.AlumnoIndex)
		v1.GET("/alumno/:id", app.AlumnoGETID)
		v1.POST("/alumno", authMiddleWare(), app.AlumnoCreate)
		v1.PUT("/alumno/:id", app.AlumnoUpdate)
		v1.DELETE("/alumno/:id", app.AlumnoDelete)

		v1.GET("/docente", app.DocenteIndex)
		v1.GET("/docente/:id", app.DocenteGETID)
		v1.POST("/docente", authMiddleWare(), app.DocenteCreate)
		v1.PUT("/docente/:id", app.DocenteUpdate)
		v1.DELETE("/docente/:id", app.DocenteDelete)

		v1.GET("/curso", app.CursoIndex)
		v1.GET("/curso/:id", app.CursoGet)
		v1.POST("/curso", authMiddleWare(), app.CursoCreate)
		v1.PUT("/curso/:id", app.CursoUpdate)
		v1.DELETE("/curso/:id", app.CursoDelete)

		v1.GET("/academicperiod", app.AcademicPeriodIndex)
		v1.POST("/academicperiod", authMiddleWare(), app.AcademicPeriodCreate)
		v1.GET("/academicperiod/:id", app.AcademicPeriodGet)
		v1.PUT("/academicperiod/:id", app.AcademicPeriodUpdate)
		v1.DELETE("/academicperiod/:id", app.AcademicPeriodDelete)

		v1.GET("/academicplan", app.AcademicPlanIndex)
		v1.POST("/academicplan", authMiddleWare(), app.AcademicPlanCreate)
		v1.GET("/academicplan/:id", app.AcademicPlanGet)
		v1.PUT("/academicplan/:id", app.AcademicPlanUpdate)
		v1.DELETE("/academicplan/:id", app.AcademicPlanDelete)

		v1.GET("/seccion", app.SeccionIndex)
		v1.POST("/seccion", authMiddleWare(), app.SeccionCreate)
		v1.GET("/seccion/:id", app.SeccionGet)
		v1.PUT("/seccion/:id", app.SeccionUpdate)
		v1.DELETE("/seccion/:id", app.SeccionDelete)

		v1.GET("/modulo", app.ModuloIndex)
		v1.POST("/modulo", authMiddleWare(), app.ModuloCreate)
		v1.GET("/modulo/:id", app.ModuloGet)
		v1.PUT("/modulo/:id", app.ModuloUpdate)
		v1.DELETE("/modulo/:id", app.ModuloDelete)

		v1.GET("/materialactividad", app.MaterialActividadIndex)
		v1.POST("/materialactividad", authMiddleWare(), app.MaterialActividadCreate)
		v1.GET("/materialactividad/:id", app.MaterialActividadGet)
		v1.PUT("/materialactividad/:id", app.MaterialActividadUpdate)
		v1.DELETE("/materialactividad/:id", app.MaterialActividadDelete)

		v1.GET("/sesionactividad", app.SesionActividadIndex)
		v1.POST("/sesionactividad", authMiddleWare(), app.SesionActividadCreate)
		v1.GET("/sesionactividad/:id", app.SesionActividadGet)
		v1.PUT("/sesionactividad/:id", app.SesionActividadUpdate)
		v1.DELETE("/sesionactividad/:id", app.SesionActividadDelete)

		v1.GET("/categoriacurso", app.CategoriaCursoCreate)
		v1.POST("/categoriacurso", authMiddleWare(), app.CategoriaCursoCreate)
		v1.GET("/categoriacurso/:id", app.CategoriaCursoGet)
		v1.PUT("/categoriacurso/:id", app.CategoriaCursoUpdate)
		v1.DELETE("/categoriacurso/:id", app.CategoriaCursoDelete)

		v1.GET("/cargaAcademica", app.CargaAcademicaIndex)
		v1.POST("/cargaAcademica", authMiddleWare(), app.CargaAcademicaCreate)
		v1.GET("/cargaAcademica/:id", app.CargaAcademicaGETID)
		v1.PUT("/cargaAcademica/:id", app.CargaAcademicaUpdate)
		v1.DELETE("/cargaAcademica/:id", app.CargaAcademicaDelete)
	}

	return r
}

func connectDBmysql() (c *gorm.DB, err error) {

	dsn := "root:aracelybriguit@tcp(localhost:3306)/wagner?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	return conn, err
}

func dbMiddleware(conn gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//c.Header("Access-Control-Allow-Origin", "http://localhost, http://localhost:8086,")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE ")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//https://dev.to/stevensunflash/a-working-solution-to-jwt-creation-and-invalidation-in-golang-4oe4

//https://www.nexmo.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
func authMiddleWare() gin.HandlerFunc { //ExtractToken
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := models.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated (IsTokenValid)."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
