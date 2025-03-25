// src/components/Register.jsx
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { registerUser } from "../api/api";
import { Link } from "react-router-dom";

function Register() {
  const [user, setUser] = useState({ username: "", password: "" });
  const [message, setMessage] = useState("");
  const navigate = useNavigate(); // Use navigate to redirect

  const handleChange = (e) => {
    setUser({ ...user, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await registerUser(user);
      setMessage("Registration successful! Redirecting to login...");
      setTimeout(() => navigate("/login"), 1000); // Redirect to Login
    } catch (error) {
      setMessage(error.response?.data?.error || "Error registering user");
    }
  };

  return (
    <div style={{ padding: "20px" }}>
      <h2>Register 📝</h2>
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
        <button type="submit">Register</button>
      </form>
      <p>{message}</p>
      <Link to="/login">Login</Link>    </div>
  );
}

export default Register;
