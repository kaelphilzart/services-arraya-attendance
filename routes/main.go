package routes

import (
	"log"
	"net/http"
	"os"
	"runtime"
	middleware "services-arraya-attendance/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Run will start the server
func Run() {
	utilsMid := new(middleware.UtilsMiddleware)

	// Nonaktifkan redirect otomatis
	router.RedirectFixedPath = false // Tambahkan ini untuk mencegah redirect otomatis

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust based on frontend URL
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.Use(utilsMid.CORSMiddleware())
	router.Use(utilsMid.RequestIDMiddleware())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	getRoutes()
	port := os.Getenv("PORT")
	router.LoadHTMLGlob("./public/html/*")

	router.Static("/public", "./public")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	if os.Getenv("SSL") == "TRUE" {

		//Generated using sh generate-certificate.sh
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{
			CERT: "./cert/myCA.cer",
			KEY:  "./cert/myCA.key",
		}

		router.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		router.Run(":" + port)
	}
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.Group("/v1")
	pingRoutes(v1)

	authRoutes(v1)
	userProfileRoutes((v1))
	userRoutes(v1)
	roleRoutes(v1)
	departmentRoutes(v1)
	CompanyRoutes(v1)
	branchRoutes(v1)
	positionRoutes(v1)

	attendanceRoutes(v1)
	shiftRoutes(v1)
	typeLeaveRoutes(v1)
	leaveRoutes(v1)
	leaveApprovalRoutes(v1)
}
