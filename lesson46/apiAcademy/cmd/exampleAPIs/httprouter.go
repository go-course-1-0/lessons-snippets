package main

import (
	"apiAcademy/internal/database/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
)

const (
	DBHost     = "localhost" // 127.0.0.1
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "academy_db"
)

type handler struct {
	DB *gorm.DB
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe", DBHost, DBPort, DBUser, DBPassword, DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected")

	h := handler{DB: db}

	db.AutoMigrate(
		&models.Admin{},
	)

	router := httprouter.New()

	router.GET("/admins", h.GetAllAdmins)
	router.POST("/admins", h.CreateAdmin)
	router.GET("/admins/:id", h.GetOneAdmin)
	router.PUT("/admins/:id", h.UpdateAdmin)
	router.DELETE("/admins/:id", h.DeleteAdmin)

	fmt.Println("starting server at :4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func (h *handler) GetAllAdmins(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var users []models.Admin
	query := r.URL.Query()
	search := query.Get("name")

	if search != "" {
		h.DB.Where("full_name ilike ?", "%"+search+"%").Find(&users)
	} else {
		h.DB.Find(&users)
	}
	usersJSON, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(usersJSON)
}

func (h *handler) GetOneAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "you try to see the admin with ID = %s\n", ps.ByName("id"))
}

func (h *handler) CreateAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	body, _ := io.ReadAll(r.Body)

	var user models.Admin
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	res := h.DB.Create(&user)
	if res.Error != nil {
		log.Fatal(res.Error)
	}

	w.Write([]byte(fmt.Sprintf("Successfully created an admin. ID: %d", user.ID)))
}

func (h *handler) UpdateAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "you try to update %s\n", ps.ByName("id"))
}

func (h *handler) DeleteAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "you try to delete %s\n", ps.ByName("id"))
}
