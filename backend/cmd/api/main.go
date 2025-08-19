package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Server struct{ db *sqlx.DB }

type User struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

func main() {
	_ = godotenv.Load(".env")

	dsn := os.Getenv("DB_DSN") // e.g. myapp:myapp@tcp(localhost:3306)/myapp?parseTime=true&charset=utf8mb4
	if dsn == "" {
		log.Fatal("DB_DSN is required")
	}

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := &Server{db: db}
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	r.GET("/api/users", s.listUsers)
	r.POST("/api/users", s.createUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) listUsers(c *gin.Context) {
	var users []User
	if err := s.db.Select(&users, "SELECT id, name, email FROM users ORDER BY id DESC"); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (s *Server) createUser(c *gin.Context) {
	var in User
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(400, gin.H{"error": "invalid payload"})
		return
	}
	res, err := s.db.Exec("INSERT INTO users(name, email) VALUES(?,?)", in.Name, in.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	id, _ := res.LastInsertId()
	in.ID = id
	c.JSON(201, in)
}
