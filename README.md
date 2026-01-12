# 🔐 Go MFA Login System

A full-stack authentication system built using Go (backend) and HTML/CSS/JavaScript (frontend).  
It demonstrates secure password handling, login flow, and Multi-Factor Authentication (OTP).

This project is created as part of the Authentication & Authorization assignment.

* * *

## 📌 Problem Statement

Build a minimal login system that:

*   Allows user registration  
      
    
*   Hashes passwords securely  
      
    
*   Allows user login  
      
    
*   Adds a basic MFA layer using OTP  
      
    
*   Gives clear success and error feedback  
      
    

* * *

## 🎯 Why This Project

Authentication is a core backend concept.  
This project demonstrates:

*   Secure password storage  
      
    
*   Defense-in-depth security  
      
    
*   Proper login flow  
      
    
*   Frontend + backend integration  
      
    
*   Real-world authentication patterns  
      
    

* * *

## 🛠️ Tech Stack Used

### Backend

*   Go (Golang)  
      
    
*   net/http  
      
    
*   bcrypt (golang.org/x/crypto)  
      
    

### Frontend

*   HTML  
      
    
*   CSS  
      
    
*   JavaScript (Fetch API)  
      
    

### Tools

*   Git & GitHub  
      
    
*   Terminal  
      
    
*   Browser DevTools  
      
    

* * *

## 📁 Project Structure

go-mfa-login/

├── backend/

│   ├── main.go

│   ├── go.mod

│   ├── handlers/

│   │   └── auth.go

│   ├── models/

│   │   └── user.go

│   └── utils/

│       ├── hash.go

│       ├── otp.go

│       └── password.go

├── frontend/

│   ├── index.html

│   ├── style.css

│   └── script.js

└── README.md

  

* * *

## 🔄 Application Flow

### 1\. Register User

*   User enters username and password  
      
    
*   Password strength is validated  
      
    
*   Password is hashed using bcrypt  
      
    
*   User is stored in memory  
      
    

### 2\. Login User

*   User enters username and password  
      
    
*   Password is verified against stored hash  
      
    
*   OTP is generated after successful password check  
      
    

### 3\. OTP Verification

*   OTP is shown on screen (demo purpose)  
      
    
*   User enters OTP  
      
    
*   OTP is verified  
      
    
*   Login is completed  
      
    

### 4\. Logout

*   Session is cleared  
      
    
*   UI is reset  
      
    
*   User can register or login again  
      
    

* * *

## 🔐 Security Concepts Implemented

### Password Hashing

*   bcrypt is used for hashing passwords  
      
    
*   Plain text passwords are never stored  
      
    

### Password Strength Validation

Password must:

*   Be at least 8 characters  
      
    
*   Contain uppercase letters  
      
    
*   Contain lowercase letters  
      
    
*   Contain numbers  
      
    

Validation is done on:

*   Frontend (UX)  
      
    
*   Backend (security)  
      
    

### Multi-Factor Authentication

*   OTP is generated after password verification  
      
    
*   OTP must be entered correctly to complete login  
      
    

* * *

## 🧠 Data Flow (Simplified)

Frontend

   ↓

Register/Login Request

   ↓

Backend Handler

   ↓

Validation + Hashing

   ↓

OTP Generation

   ↓

OTP Verification

   ↓

Login Success

  

* * *

## 🖥️ Backend Logic Explanation

### Register

*   Decode request body  
      
    
*   Validate inputs  
      
    
*   Validate password strength  
      
    
*   Hash password  
      
    
*   Store user  
      
    

### Login

*   Verify username exists  
      
    
*   Compare password hash  
      
    
*   Generate OTP  
      
    
*   Send OTP to frontend  
      
    

### Verify OTP

*   Match OTP with stored value  
      
    
*   Allow login on success  
      
    

* * *

## 🎨 Frontend Features

*   Clean card-based UI  
      
    
*   Smooth animations and transitions  
      
    
*   Real-time password strength feedback  
      
    
*   Clear success and error messages  
      
    
*   OTP display for demo  
      
    
*   Logout functionality  
      
    
*   Fully reset state on logout  
      
    

* * *

## 🧪 How to Run the Project

### 1️⃣ Start Backend

cd backend

go mod tidy

go run main.go

  

Backend runs on:

http://localhost:8080

  

* * *

### 2️⃣ Open Frontend

cd frontend

xdg-open index.html

  

Or open index.html manually in a browser.


* * *

## 🧾 Conclusion

This project demonstrates a complete authentication flow with:

*   Secure password handling  
      
    
*   MFA using OTP  
      
    
*   Backend + frontend integration  
      
    
*   Clean and explainable code  

* * *

## Demo

[Screencast from 2026-01-12 23-36-11.webm](https://github.com/user-attachments/assets/952076c1-1374-4940-ad74-ef334a7ad4c1)
