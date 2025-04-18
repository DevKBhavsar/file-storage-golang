
# Project Title

A brief description of what this project does and who it's for

# 📚  File Storage API with User Authentication and Storage Limits  

---

## 🚀 Overview
This project implements a **secure file storage API** that allows users to:  
- 📂 Upload and manage files.  
- 🔐 Authenticate using JWT tokens.  
- 📊 Enforce storage quotas per user.  

---

## 🎯 Features
- ✅ User Registration & Login with JWT Authentication.
- 📂 File Upload with Storage Limit Validation.
- 📊 Storage Quota Enforcement and Remaining Storage Check.
- 📃 File Retrieval and Metadata Display.
- 🔄 Secure API Endpoints with Middleware Protection.

---

## ⚙️ Tech Stack
- 🐹 **Backend:** Golang + Gorilla Mux + MongoDB  
- ⚛️ **Frontend:** React (Vite) + Tailwind CSS  
- 🌐 **API Hosting:** Railway  
- 🎯 **Frontend Hosting:** Vercel  

---

## 📚 API Documentation
### 🔥 Base URL
```
https://wobot-production.up.railway.app
```

### ✅ API Endpoints

| #  | Endpoint             | Method | Description                  | Auth Required | Expected Response             |
|----|---------------------|--------|------------------------------|---------------|-------------------------------|
| 1  | `/`                  | GET    | Check if API is running      | ❌ No         | `API is running...`           |
| 2  | `/register`          | POST   | Register a new user          | ❌ No         | `User registered successfully` |
| 3  | `/login`             | POST   | Authenticate and get JWT     | ❌ No         | `JWT Token`                   |
| 4  | `/upload`            | POST   | Upload a file                | ✅ Yes        | `File uploaded successfully`  |
| 5  | `/storage/remaining` | GET    | Check remaining storage      | ✅ Yes        | `Remaining and total storage` |
| 6  | `/files`             | GET    | Get uploaded files list      | ✅ Yes        | `List of files`               |

---

## 📝 Setup Instructions
### ⚡️ 1. Clone Repository
```bash
git clone 
cd wobot
```

---

### 🔥 2. Backend Setup
```bash
cd file-storage-api
go mod tidy
go run main.go
```
✅ Backend should run at: `http://localhost:8080`  

---

### ⚛️ 3. Frontend Setup
```bash
cd frontend
npm install
npm run dev
```
✅ Frontend should run at: `http://localhost:5173`  

---

### 🔐 4. Environment Variables
Create a `.env` file in the `/api` and `/frontend` directories with the following:  

#### 📂 **Backend `/api/.env`**
```
MONGO_URI=<your-mongodb-uri>
JWT_SECRET=<your-jwt-secret>
PORT=8080
```

#### ⚛️ **Frontend `/frontend/.env`**
```
VITE_API_URL=https://wobot-production.up.railway.app
```

---

## 🧪 Testing
- Use **Postman** to test API endpoints.


---

## 📩 Deployment
### ✅ Backend on Railway
```
https://wobot-production.up.railway.app
```

### ✅ Frontend on Vercel
```
https://file-storage-frontend-khaki.vercel.app
```

---

## 🤝 Acknowledgements
- Vercel for seamless frontend deployment.  
- Railway for efficient backend hosting.  

---
