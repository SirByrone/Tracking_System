package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db            *gorm.DB
	adminUsername = "SCIADMIN"
	adminPassword = os.Getenv("SCI123")
	templates     = template.Must(template.ParseGlob("templates/*.html"))
)

type User struct {
	ID                   uint `gorm:"primaryKey"`
	Name                 string
	Username             string `gorm:"unique"`
	Password             string
	IsAdmin              bool
	RegNumber            string
	Course               string
	YearOfStudy          string
	Department           string
	College              string
	Address              string
	MobileNumber         string
	AttachmentFrom       string
	AttachmentTo         string
	LecturerInChargeName string
	LecturerContact      string
	SName                string
	SAddress             string
	Sphone               string
	SCounty              string
	SConstituency        string
}

func main() {
	dsn := "byrone:incorrect@tcp(127.0.0.1:3306)/attachee_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&User{})

	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/attachee", attacheeHandler)
	http.HandleFunc("/attachee/update", updateAttacheeHandler)
	http.HandleFunc("/admin/add", addAttacheeHandler)
	http.HandleFunc("/admin/delete", deleteAttacheeHandler)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		log.Println("Received login attempt:", username, password)

		var user User
		result := db.Where("username = ? AND password = ?", username, password).First(&user)
		if result.Error != nil {
			log.Println("Error fetching user:", result.Error)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		log.Println("User found:", user)

		// Check if the user is an admin
		if user.IsAdmin {
			http.SetCookie(w, &http.Cookie{
				Name:  "username",
				Value: username,
				Path:  "/",
			})
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		// Set a cookie with the username
		http.SetCookie(w, &http.Cookie{
			Name:  "username",
			Value: username,
			Path:  "/",
		})
		http.Redirect(w, r, "/attachee", http.StatusSeeOther)
		return
	}
	templates.ExecuteTemplate(w, "login.html", nil)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		log.Println("Error fetching cookie:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	username := cookie.Value

	var user User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil || !user.IsAdmin {
		log.Println("Unauthorized access attempt by:", username)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	log.Println("Admin logged in")
	users := []User{}
	db.Find(&users)
	templates.ExecuteTemplate(w, "admin.html", users)
}

func attacheeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	username := cookie.Value

	log.Println("Fetching data for user:", username)

	var user User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		log.Println("Error fetching user:", result.Error)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	log.Println("User data fetched:", user)
	templates.ExecuteTemplate(w, "attachee.html", user)
}

func updateAttacheeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		cookie, err := r.Cookie("username")
		if err != nil {
			log.Println("Error fetching cookie:", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		username := cookie.Value

		log.Println("Updating data for user:", username)

		var user User
		result := db.Where("username = ?", username).First(&user)
		if result.Error != nil {
			log.Println("Error fetching user for update:", result.Error)
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		user.Name = r.FormValue("name")
		user.RegNumber = r.FormValue("reg_number")
		user.Course = r.FormValue("course")
		user.YearOfStudy = r.FormValue("year_of_study")
		user.Department = r.FormValue("department")
		user.College = r.FormValue("college")
		user.Address = r.FormValue("address")
		user.MobileNumber = r.FormValue("mobile_number")
		user.AttachmentFrom = r.FormValue("attachment_from")
		user.AttachmentTo = r.FormValue("attachment_to")
		user.LecturerInChargeName = r.FormValue("lecturer_in_charge")
		user.LecturerContact = r.FormValue("lecturer's_contact")
		user.SName = r.FormValue("sponsor_name")
		user.Sphone = r.FormValue("sponsor_contact")
		user.SAddress = r.FormValue("sponsor_address")
		user.SCounty = r.FormValue("sponsor_county")
		user.SConstituency = r.FormValue("sponsor_constituency")

		db.Save(&user)
		log.Println("User data updated:", user)
		http.Redirect(w, r, "/attachee", http.StatusSeeOther)
	}
}

func addAttacheeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		username := r.FormValue("username")
		password := r.FormValue("password")
		user := User{Name: name, Username: username, Password: password, IsAdmin: false}
		db.Create(&user)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func deleteAttacheeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		db.Delete(&User{}, id)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
