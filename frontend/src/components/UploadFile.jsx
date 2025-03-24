// src/components/UploadFile.jsx
import React, { useState } from "react";
import { uploadFile } from "../api/api";
import { getToken } from "../utils/auth";

function UploadFile() {
  const [file, setFile] = useState(null);
  const [message, setMessage] = useState("");

  // Handle file change
  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  // Handle file upload
  const handleUpload = async (e) => {
    e.preventDefault();

    if (!file) {
      setMessage("Please select a file to upload.");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);

    try {
      const token = getToken();
      const response = await uploadFile(formData, token);
      setMessage(response.message || "File uploaded successfully! Refresh to view in info");
    } catch (error) {
      setMessage(
        error.response?.data?.error || "Error uploading file. Try again."
      );
    }
  };

  return (
    <div style={{ padding: "20px" }}>
      <h3>Upload a File 📂</h3>
      <form onSubmit={handleUpload}>
        <input type="file" onChange={handleFileChange} required />
        <button type="submit">Upload</button>
      </form>
      <p>{message}</p>
    </div>
  );
}

export default UploadFile;
