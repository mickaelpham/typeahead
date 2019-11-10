package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var trie Trie

func searchHandler(w http.ResponseWriter, r *http.Request) {
	prefix := r.URL.Query().Get("q")
	w.Header().Set("content-type", "application/json")
	words := trie.Words(prefix)
	err := json.NewEncoder(w).Encode(words)
	if err != nil {
		log.Fatal(err)
	}
}

func buildTrie() {
	trie = Trie{}
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trie.Put(scanner.Text())
	}
}

func main() {
	buildTrie()
	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
