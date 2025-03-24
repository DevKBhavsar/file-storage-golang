package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"file-storage-api/app/models"
	"file-storage-api/app/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UploadFile handles file uploads
func UploadFile(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	// Parse form data
	err := r.ParseMultipartForm(10 << 20) // Max upload size: 10 MB
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileSize := handler.Size
	collection := utils.GetCollection(utils.InitDB(), "users")

	// Get user details
	var user models.User
	err = collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check if user has enough storage left
	if user.UsedStorage+fileSize > user.StorageQuota {
		http.Error(w, "Insufficient storage space", http.StatusPaymentRequired)
		return
	}

	// Create user-specific folder
	userFolder := filepath.Join("uploads", username)
	if _, err := os.Stat(userFolder); os.IsNotExist(err) {
		os.MkdirAll(userFolder, os.ModePerm)
	}

	// Save file to user folder
	filePath := filepath.Join(userFolder, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy file data to destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error writing file", http.StatusInternalServerError)
		return
	}

	// Update user's used storage
	user.UsedStorage += fileSize
	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": user.ID},
		bson.M{"$set": bson.M{"used_storage": user.UsedStorage}},
	)
	if err != nil {
		http.Error(w, "Error updating storage", http.StatusInternalServerError)
		return
	}

	// Save file metadata (optional)
	fileMetadata := models.FileMetadata{
		ID:         primitive.NewObjectID(),
		Username:   username,
		Filename:   handler.Filename,
		Size:       fileSize,
		UploadTime: time.Now(),
	}
	fileCollection := utils.GetCollection(utils.InitDB(), "files")
	_, err = fileCollection.InsertOne(context.TODO(), fileMetadata)
	if err != nil {
		http.Error(w, "Error saving file metadata", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "File uploaded successfully"})
}

// GetFiles retrieves a list of uploaded files for the authenticated user
func GetFiles(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	userFolder := filepath.Join("uploads", username)

	// Check if the folder exists
	if _, err := os.Stat(userFolder); os.IsNotExist(err) {
		http.Error(w, "No files found", http.StatusNotFound)
		return
	}

	// Get list of files
	files, err := os.ReadDir(userFolder)
	if err != nil {
		http.Error(w, "Error reading files", http.StatusInternalServerError)
		return
	}

	fileListChan := make(chan string, len(files))
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(f os.DirEntry) {
			defer wg.Done()
			fileListChan <- f.Name()
		}(file)
	}

	// Close channel after all goroutines are done
	go func() {
		wg.Wait()
		close(fileListChan)
	}()

	var fileList []string
	for fileName := range fileListChan {
		fileList = append(fileList, fileName)
	}

	// Return list of filenames
	json.NewEncoder(w).Encode(map[string]interface{}{
		"files": fileList,
	})
}

// GetRemainingStorage returns the remaining storage for the authenticated user
// GetRemainingStorage returns the remaining storage for the authenticated user
func GetRemainingStorage(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	collection := utils.GetCollection(utils.InitDB(), "users")

	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Convert bytes to MB (1 MB = 1024 * 1024 bytes)
	toMB := func(bytes int64) float64 {
		return float64(bytes) / (1024 * 1024)
	}

	remainingStorageMB := toMB(user.StorageQuota - user.UsedStorage)
	totalStorageMB := toMB(user.StorageQuota)
	usedStorageMB := toMB(user.UsedStorage)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"total_storage_mb":     fmt.Sprintf("%.2f MB", totalStorageMB),
		"used_storage_mb":      fmt.Sprintf("%.2f MB", usedStorageMB),
		"remaining_storage_mb": fmt.Sprintf("%.2f MB", remainingStorageMB),
	})
}
