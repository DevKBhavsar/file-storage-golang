// src/components/StorageInfo.jsx
import React, { useEffect, useState } from "react";
import { getRemainingStorage } from "../api/api";
import { getToken } from "../utils/auth";

function StorageInfo() {
  const [storage, setStorage] = useState({remaining_storage_mb: '0', total_storage_mb: '0', used_storage_mb: '0'});
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchStorage = async () => {
      try {
        const token = getToken();
        const response = await getRemainingStorage(token);
        // console.log("we got response",response)
        setStorage(response);
      } catch (error) {
        setError(
          error.response?.data?.error || "Error fetching storage info."
        );
      }
    };
    fetchStorage();
  }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h3>Storage Info 📊</h3>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <p>

        <strong>Total Storage:</strong> {(storage.total_storage_mb)}
      </p>
      <p>
        <strong>Used Storage:</strong> {(storage.used_storage_mb)}
      </p>
      <p>
        <strong>Remaining Storage:</strong> {(storage.remaining_storage_mb)}
      </p>
    </div>
  );
}


export default StorageInfo;
