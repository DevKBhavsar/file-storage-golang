// src/pages/Home.jsx
import React from "react";
import { Link } from "react-router-dom";

function Home() {
  return (
    <div style={{ padding: "20px" }}>
      <h2>Welcome to File Storage App 🚀</h2>
      <p>
        <Link to="/register">Register</Link> |{" "}
        <Link to="/login">Login</Link> |{" "}
        <Link to="/dashboard">Go to Dashboard</Link>
      </p>
    </div>
  );
}

export default Home;
