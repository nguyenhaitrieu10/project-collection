package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Custom Middleware
func LoggerCustom() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func SetupRouter() *gin.Engine {
	// Write logs to files
	gin.DisableConsoleColor()
	f, err := os.Create("logs/gin.log")
	if err != nil {
		fmt.Println(err)
	} else {
		gin.DefaultWriter = io.MultiWriter(f)
	}

	// router := gin.Default()
	router := gin.New()
	// router.Use(gin.Logger())
	// or Custom format logs (optional)
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//  /user/john
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	//  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// curl -X POST localhost:8080/form_post -H "Content-Type: application/json"   --data '{"message":"xyz"}'
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// curl -X POST localhost:8080/post -H "Content-Type: application/x-www-form-urlencoded"   --data 'name=manu&message=this_is_great'
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		str := fmt.Sprintf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.String(http.StatusOK, str)
	})

	// Serving static files
	router.Static("/public", "./public")
	router.StaticFS("/more_static", http.Dir("my_file_system"))

	// HTML rendering
	router.LoadHTMLGlob("templates/**/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	// Redirect
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	// View http://localhost:8080/public/index.html
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]

		for _, file := range files {
			filename := "./public/images/" + file.Filename
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files", len(files)))
	})

	// // Grouping routes
	// v1 := router.Group("/v1")
	// {
	// v1.POST("/login", loginEndpoint)
	// v1.POST("/submit", submitEndpoint)
	// }
	// authorized := r.Group("/")
	// authorized.Use(AuthRequired())
	// {
	// 	authorized.POST("/login", loginEndpoint)
	// 	authorized.POST("/submit", submitEndpoint)
	// 	// nested group
	// 	testing := authorized.Group("testing")
	// 	testing.GET("/analytics", analyticsEndpoint)
	// }

	return router
}

func main() {
	router := SetupRouter()
	s := &http.Server{
		Addr:           "0.0.0.0:9000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	// router.Run("0.0.0.0:9000")
}
