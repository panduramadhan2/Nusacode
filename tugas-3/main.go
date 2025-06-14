package main

import (
	"fmt"
	"tugas-3/client"
)

func main() {
	fmt.Println("== GET Posts ==")
	client.GetPosts()

	fmt.Println("\n== POST with Authorization ==")
	client.PostWithAuth()

	fmt.Println("\n== PUT Update Post ==")
	client.UpdatePost()
}
