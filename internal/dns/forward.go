package dns

import (
	"fmt"

	"github.com/miekg/dns"
)

func forwardQuery(req *dns.Msg, upstream string) (*dns.Msg, error) {
	fmt.Println("Inside Forward Query: ")
	client := new(dns.Client)
	resp, _, err := client.Exchange(req, upstream)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
