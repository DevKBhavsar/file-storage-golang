// src/pages/Dashboard.jsx
import React from "react";
import UploadFile from "../components/UploadFile";
import FileList from "../components/FileList";
import StorageInfo from "../components/StorageInfo";

function Dashboard() {
  return (
    <div style={{ padding: "20px" }}>
      <h2>Dashboard 📂</h2>
      <StorageInfo />
      <hr />
      <UploadFile />
      <hr />
      <FileList />
    </div>
  );
}

export default Dashboard;
