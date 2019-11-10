package main

import "sort"

type Trie struct {
	Root *Node
}

type Node struct {
	Val  string
	Next map[string]*Node
}

func (t *Trie) Put(word string) {
	if t.Root == nil {
		t.Root = &Node{Next: make(map[string]*Node)}
	}

	curr := t.Root
	for _, char := range word {
		if curr.Next[string(char)] == nil {
			curr.Next[string(char)] = &Node{Next: make(map[string]*Node)}
		}

		curr = curr.Next[string(char)]
	}

	curr.Val = word
}

func (t *Trie) Words(prefix string) (words []string) {
	if t.Root == nil {
		return
	}

	curr := t.Root
	for _, char := range prefix {
		curr = curr.Next[string(char)]
		if curr == nil {
			return
		}
	}

	s := []*Node{curr}
	for len(s) > 0 {
		curr, s = s[len(s)-1], s[:len(s)-1]

		if curr.Val != "" {
			words = append(words, curr.Val)
		}

		next := make([]string, 0)
		for k := range curr.Next {
			next = append(next, string(k))
		}
		sort.Strings(next)

		for i := len(next) - 1; i >= 0; i-- {
			s = append(s, curr.Next[next[i]])
		}
	}

	return
}
