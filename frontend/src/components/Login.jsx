// src/components/Login.jsx
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { loginUser } from "../api/api";
import { saveToken } from "../utils/auth";

function Login() {
  const [user, setUser] = useState({ username: "", password: "" });
  const [message, setMessage] = useState("");
  const navigate = useNavigate(); // Use navigate to redirect

  const handleChange = (e) => {
    setUser({ ...user, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await loginUser(user);
      saveToken(response.token); // Save JWT token
      setMessage("Login successful! Redirecting...");
      setTimeout(() => navigate("/dashboard"), 1000); // Redirect to Dashboard
    } catch (error) {
      setMessage(error.response?.data?.error || "Error logging in");
    }
  };

  return (
    <div style={{ padding: "20px" }}>
      <h2>Login </h2>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="username"
          placeholder="Username"
          value={user.username}
          onChange={handleChange}
          required
        />
        <input
          type="password"
          name="password"
          placeholder="Password"
          value={user.password}
          onChange={handleChange}
          required
        />
        <button type="submit">Login</button>
      </form>
      <p>{message}</p>
      <a href="/register">register</a>

    </div>
  );
}

export default Login;
