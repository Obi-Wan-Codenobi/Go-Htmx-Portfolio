package main

import (
	"fmt"
	"net/http"
    //"io/ioutil"
    //"net/http"
    //"os"
    //"path/filepath"
    "strings"

    "github.com/russross/blackfriday/v2"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./../../client/website-htmx")))
	http.HandleFunc("/api/data", apiDataHandler)
	http.HandleFunc("/api/send-email", sendEmailHandler)
    http.HandleFunc("/api/blogposts", blogPostsHandler)

	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate API response (replace with actual API call)
	apiResponse := "API Data Loaded!"
	fmt.Fprint(w, apiResponse)
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate sending an email (replace with actual email sending logic)
	fmt.Fprint(w, "Email Sent!")
}

func blogPostsHandler(w http.ResponseWriter, r *http.Request) {
    // Simulate fetching Markdown files from GitHub (replace with actual logic)
    markdownFiles, err := fetchMarkdownFiles()
    if err != nil {
        http.Error(w, "Error fetching blog posts", http.StatusInternalServerError)
        return
    }

    // Convert Markdown to HTML
    var htmlContent strings.Builder
    for _, markdown := range markdownFiles {
        html := blackfriday.Run([]byte(markdown))
        htmlContent.Write(html)
    }

    // Respond with the generated HTML
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(htmlContent.String()))
}

func fetchMarkdownFiles() ([]string, error) {
    // Simulate fetching Markdown files from a GitHub repository
    // Replace this logic with actual code to fetch Markdown files from your GitHub repository
    markdownFiles := []string{
        "# Blog Post 1\n\nThis is the content of blog post 1.",
        "# Blog Post 2\n\nThis is the content of blog post 2.",
    }
    return markdownFiles, nil
}
