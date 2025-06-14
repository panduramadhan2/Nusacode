package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"tugas-3/models"
)

func UpdatePost() {
	updatedPost := models.Post{
		Title:  "Updated Post",
		Body:   "This post has been updated.",
		UserID: 1,
	}

	jsonData, _ := json.Marshal(updatedPost)
	req, _ := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

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
