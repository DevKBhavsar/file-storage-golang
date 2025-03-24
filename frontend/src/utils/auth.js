// src/utils/auth.js

// Save token to local storage
export const saveToken = (token) => {
    localStorage.setItem("token", token);
  };
  
  // Get token from local storage
  export const getToken = () => {
    return localStorage.getItem("token");
  };
  
  // Remove token from local storage (logout)
  export const removeToken = () => {
    localStorage.removeItem("token");
  };
  