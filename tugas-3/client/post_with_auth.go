package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"tugas-3/models"
)

func PostWithAuth() {
	post := models.Post{
		Title:  "Post with Authorization",
		Body:   "This post includes an authorization header.",
		UserID: 1,
	}

	jsonData, _ := json.Marshal(post)
	req, _ := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 12345")

	resp, err := HTTPClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)
	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)
	fmt.Println(body.String())
}
