package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"go-core-04/GoSearch/pkg/crawler/spider"
)

// Функция для проверки наличия слова в тексте
func containsWord(text, word string) bool {
	return strings.Contains(text, word)
}

func main() {
	s := spider.New() // Создаем экземпляр поискового робота

	// Обработка флага "-s"
	searchWord := flag.String("s", "", "Search word")
	flag.Parse()

	// Сканирование сайта "go.dev"
	goDevData, err := s.Scan("https://go.dev", 2)
	if err != nil {
		log.Fatal(err)
	}

	// Сканирование сайта "golang.org"
	golangOrgData, err := s.Scan("https://golang.org", 2)
	if err != nil {
		log.Fatal(err)
	}

	// Объединение результатов сканирования
	allData := append(goDevData, golangOrgData...)

	// Вывод результатов сканирования
	for _, doc := range allData {
		fmt.Printf("URL: %s\nTitle: %s\n\n", doc.URL, doc.Title)
	}

	// Поиск по заданному слову и вывод соответствующих ссылок
	if *searchWord != "" {
		fmt.Printf("Search results for word '%s':\n", *searchWord)
		for _, doc := range allData {
			if containsWord(doc.Body, *searchWord) {
				fmt.Printf("- %s\n", doc.URL)
			}
		}
	}
}
