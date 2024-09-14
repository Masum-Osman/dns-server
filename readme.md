# Go DNS Server

This is a simple DNS server implemented in Go that uses a Trie data structure to store domain-IP mappings. It can respond to DNS queries and dynamically add missing domains by querying an upstream DNS server.

## Features

- Store and resolve domain-IP mappings using a Trie.
- Automatically query and add domains that are not present in the Trie from an upstream DNS server.
- Command to print all current domain-IP mappings in the Trie.

## Project Structure:
```
dns-server/
│
├── cmd/
│   └── main.go            # Entry point of the application
│
├── internal/
│   ├── trie/
│   │   ├── trie.go        # Trie data structure and methods
│   │
│   ├── dns/
│   │   ├── server.go      # DNS server logic and handlers
│   │   ├── forward.go     # DNS forwarding logic
│   │
│   └── config/
│       └── config.go      # Configuration for the DNS server
│
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
    ```sh
    git clone github.com/masum-osman/dns-server
    cd dns-server
    ```

2. Initialize the project:
    ```sh
    go mod init github.com/masum-osman/dns-server
    go mod tidy
    ```

### Usage

1. Run the DNS server:
    ```sh
    go run main.go
    ```

2. To query the DNS server, you can use `dig` or `nslookup`:
    ```sh
    dig @127.0.0.1 -p 8053 www.example.com
    ```

3. To print all current domain-IP mappings in the Trie:
    ```sh
    go run main.go print
    ```

## Example

When you query for a domain that is not in the Trie, the server will query the upstream DNS server (e.g., Google DNS) and add the result to the Trie. Subsequent queries for the same domain will be served from the Trie.

### Sample Query
```sh
dig @127.0.0.1 -p 8053 www.google.com
