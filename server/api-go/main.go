package main

import (
    "encoding/json"
    //"encoding/base64"
    "log"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

    "github.com/russross/blackfriday/v2"
)

type GitHubContent struct {
    Name string `json:"name"`
    Type string `json:"type"`
    Path     string `json:"path"` 
    //Content  string `json:"content"`
    //Encoding string `json:"encoding"`
}

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
	markdownFiles, err := fetchMarkdownFilesFromGitHub()
	if err != nil {
        log.Println("Error fetching blog posts:", err)
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

func fetchMarkdownFilesFromGitHub() ([]string, error) {

    repoURL := "https://raw.githubusercontent.com/Obi-Wan-Codenobi/Go-Htmx-Portfolio/main"
    treeURL := "https://api.github.com/repos/Obi-Wan-Codenobi/Go-Htmx-Portfolio/contents/content/blogposts"
    filePaths, err := fetchGitHubContent(treeURL)

	if err != nil {
        log.Println("Error fetching GitHub content:", err)
		return nil, err
	}

	// Fetch content of each Markdown file
	var result []string
	for _, item := range filePaths {
		if item.Type == "file" && strings.HasSuffix(item.Name, ".md") {
			fileURL := repoURL + "/" + item.Path
            log.Println("Fetching file content from URL:", fileURL)
			content, err := fetchFileContent(fileURL)
			if err != nil {
                log.Println("Error fetching file content:", err)
				return nil, err
			}
			result = append(result, content)
		}
	}

	return result, nil
}

func fetchGitHubContent(url string) ([]GitHubContent, error) {
	resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    
    log.Println("GitHub API response body:", string(body))

    var content []GitHubContent
    err = json.Unmarshal(body, &content)
    if err != nil {
        return nil, err
    }

    return content, nil
}


func fetchFileContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}


//func blogPostsHandler(w http.ResponseWriter, r *http.Request) {
//    // Simulate fetching Markdown files from GitHub (replace with actual logic)
//    markdownFiles, err := fetchMarkdownFilesFromGitHub()
//    if err != nil {
//        http.Error(w, "Error fetching blog posts", http.StatusInternalServerError)
//        return
//    }
//
//    // Convert Base64-encoded Markdown to HTML
//    var htmlContent strings.Builder
//    for _, base64Markdown := range markdownFiles {
//        // base64Markdown.Content is already Base64-encoded
//        markdown := []byte(base64Markdown.Content)
//
//        html := blackfriday.Run(markdown)
//        htmlContent.Write(html)
//    }
//
//    // Respond with the generated HTML
//    w.Header().Set("Content-Type", "text/html")
//    w.Write([]byte(htmlContent.String()))
//}
//
//func fetchMarkdownFilesFromGitHub() ([]string, error) {
//	// GitHub repository URL
//	repoURL := "https://api.github.com/repos/Obi-Wan-Codenobi/Go-Htmx-Portfolio/contents/content/blogposts"
//
//	// Fetch content of the blogposts directory
//    content, err := fetchGitHubContent(repoURL)
//	if err != nil {
//		return nil, err
//	}
//
//	// Extract file names from the directory listing
//	var fileNames []string
//	for _, item := range content {
//		if item.Type == "file" && strings.HasSuffix(item.Name, ".md") {
//			fileNames = append(fileNames, item.Name)
//		}
//	}
//
//    // Fetch content of each Markdown file
//	var result []string
//	for _, fileName := range fileNames {
//		fileURL := repoURL + "/" + fileName
//		content, err := fetchFileContent(fileURL)
//		if err != nil {
//			return nil, err
//		}
//		result = append(result, content)
//	}
//
//	return result, nil
//}
//
//func fetchGitHubContent(url string) ([]GitHubContent, error) {
//	resp, err := http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	var content []GitHubContent
//	err = json.Unmarshal(body, &content)
//	if err != nil {
//		return nil, err
//	}
//
//	return content, nil
//}
//
//func fetchFileContent(url string) (string, error) {
//	resp, err := http.Get(url)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//
//	return string(body), nil
//}
