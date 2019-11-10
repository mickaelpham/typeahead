package main

import (
	"reflect"
	"testing"
)

func TestWords(t *testing.T) {
	words := []string{
		"she",
		"sells",
		"sea",
		"shells",
		"by",
		"the",
		"sea",
		"shore",
	}

	trie := Trie{}
	for _, w := range words {
		trie.Put(w)
	}

	got := trie.Words("sh")
	want := []string{"she", "shells", "shore"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
