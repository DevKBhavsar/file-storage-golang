// src/components/FileList.jsx
import React, { useEffect, useState } from "react";
import { getUploadedFiles } from "../api/api";
import { getToken } from "../utils/auth";

function FileList() {
  const [files, setFiles] = useState([]);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchFiles = async () => {
      try {
        const token = getToken();
        const response = await getUploadedFiles(token);
        console.log(response.files)
        setFiles(response.files || []);
      } catch (error) {
        setError(error.response?.data?.error || "Error fetching files.");
      }
    };
    fetchFiles();
  }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h3>Your Uploaded Files 📄</h3>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {files.length === 0 ? (
        <p>No files uploaded yet.</p>
      ) : (
        <ul>
          {files.map((file, index) => (
            <li key={index}>
              {file} 
            
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

// Format file size in MB or KB
const formatFileSize = (bytes) => {
  if (bytes < 1024) return `${bytes} Bytes`;
  if (bytes < 1024 * 1024)
    return `${(bytes / 1024).toFixed(2)} KB`;
  return `${(bytes / (1024 * 1024)).toFixed(2)} MB`;
};

// Format upload date properly
const formatDate = (timestamp) => {
  const date = new Date(timestamp);
  if (isNaN(date.getTime())) return "Invalid Date";
  return date.toLocaleString();
};

export default FileList;
