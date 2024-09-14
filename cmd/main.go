package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/masum-osman/dns-server/internal/dns"
	"github.com/masum-osman/dns-server/internal/trie"
)

func main() {
	trie := trie.NewTrie()
	trie.Insert("www.facebook.com.", net.ParseIP("31.13.71.36"))
	trie.Insert("www.google.com.", net.ParseIP("142.250.74.238"))

	printCmd := flag.Bool("print", false, "Print all DNS records")
	flag.Parse()

	if *printCmd {
		fmt.Println("Printing all DNS records:")
		trie.PrintRecords()
		return
	}

	address := "127.0.0.1:8053"
	server := dns.NewServer(address, trie)

	log.Printf("Starting DNS server on %s", address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start DNS server: %s", err.Error())
	}
}
