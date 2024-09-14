package trie

import (
	"fmt"
	"net"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	ip       net.IP
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(domain string, ip net.IP) {
	node := t.root
	for _, ch := range domain {
		if _, exist := node.children[ch]; !exist {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.ip = ip
}

func (t *Trie) Search(domain string) (net.IP, bool) {
	node := t.root
	for _, ch := range domain {
		if _, exist := node.children[ch]; !exist {
			return nil, false
		}
		node = node.children[ch]
	}
	if node.isEnd {
		return node.ip, true
	}
	return nil, false
}

func (t *Trie) PrintRecords() {
	t.printRecordsHelper(t.root, "")
}

func (t *Trie) printRecordsHelper(node *TrieNode, prefix string) {
	if node.isEnd {
		fmt.Printf("%s -> %s\n", prefix, node.ip)
	}
	for char, child := range node.children {
		t.printRecordsHelper(child, prefix+string(char))
	}
}
