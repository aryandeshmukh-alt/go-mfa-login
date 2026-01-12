package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"go-mfa-login/models"
	"go-mfa-login/utils"
)

var users = make(map[string]models.User)

func Register(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if data.Username == "" || data.Password == "" {
		http.Error(w, "Username and password required", http.StatusBadRequest)
		return
	}

	// 🔐 NEW: Password strength validation
	if !utils.IsStrongPassword(data.Password) {
		http.Error(
			w,
			"Password too weak. Use 8+ chars with upper, lower & number",
			http.StatusBadRequest,
		)
		return
	}

	log.Println("Register attempt:", data.Username)

	hash, _ := utils.HashPassword(data.Password)

	users[data.Username] = models.User{
		Username: data.Username,
		Password: hash,
	}

	log.Println("User registered:", data.Username)
	w.Write([]byte("Registration successful. Please login."))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&data)

	log.Println("Login attempt:", data.Username)

	user, ok := users[data.Username]
	if !ok || utils.CheckPassword(user.Password, data.Password) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	otp := utils.GenerateOTP()
	user.OTP = otp
	users[data.Username] = user

	log.Println("OTP generated for", data.Username, "OTP:", otp)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password verified. OTP generated.",
		"otp":     otp,
	})
}

func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	json.NewDecoder(r.Body).Decode(&data)

	log.Println("OTP verification:", data["username"])

	user := users[data["username"]]
	if user.OTP != data["otp"] {
		log.Println("Invalid OTP:", data["username"])
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	}

	log.Println("Login successful:", data["username"])
	w.Write([]byte("Login successful"))
}
