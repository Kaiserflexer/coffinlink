package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func checkDomain(domain string) string {
	resp, err := http.Get("http://" + domain)
	if err != nil {
		return color.RedString(fmt.Sprintf("%s: Failed to connect", domain))
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return color.GreenString(fmt.Sprintf("%s: 200 OK", domain))
	}
	return color.RedString(fmt.Sprintf("%s: %d", domain, resp.StatusCode))
}

func processFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		result := checkDomain(domain)
		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func main() {
	dirPath := "."               // Путь к каталогу, который нужно проверить
	fileToCheck := "domains.txt" // Имя файла, который нужно проверить
	filePath := filepath.Join(dirPath, fileToCheck)

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("Processing:", filePath)
		processFile(filePath)
	} else {
		fmt.Println("File not found:", filePath)
	}
}
