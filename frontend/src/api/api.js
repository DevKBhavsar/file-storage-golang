// src/api/api.js
import axios from "axios";

// Define API base URL from Vite env variable or fallback to localhost
const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

// Create axios instance with base config
const api = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Register User
export const registerUser = async (userData) => {
  const response = await api.post("/register", userData);
  return response.data;
};

// Login User
export const loginUser = async (userData) => {
  const response = await api.post("/login", userData);
  return response.data;
};

// Upload File
export const uploadFile = async (formData, token) => {
  const response = await api.post("/upload", formData, {
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "multipart/form-data",
    },
  });
  return response.data;
};

// Get Remaining Storage
export const getRemainingStorage = async (token) => {
  const response = await api.get("/storage/remaining", {
    headers: { Authorization: `Bearer ${token}` },
  });
  return response.data;
};

// Get Uploaded Files
export const getUploadedFiles = async (token) => {
  const response = await api.get("/files", {
    headers: { Authorization: `Bearer ${token}` },
  });
 
  return response.data;
};

export default api;
