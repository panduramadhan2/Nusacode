package client

import (
	"encoding/json"
	"fmt"
	"io"
	"tugas-3/models"
)

func GetPosts() {
	resp, err := HTTPClient.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var posts []models.Post
	json.Unmarshal(body, &posts)

	for _, post := range posts[:5] { // tampilkan hanya 5 pertama
		fmt.Printf("ID: %d, Title: %s\n", post.ID, post.Title)
	}
}
