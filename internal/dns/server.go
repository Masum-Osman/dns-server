package dns

import (
	"fmt"
	"log"

	"github.com/masum-osman/dns-server/internal/trie"
	"github.com/miekg/dns"
)

func NewServer(address string, trie *trie.Trie) *dns.Server {
	mux := dns.NewServeMux()
	mux.HandleFunc(".", handleRequest(trie))

	return &dns.Server{
		Addr:    address,
		Net:     "udp",
		Handler: mux,
	}
}

func handleRequest(trie *trie.Trie) func(w dns.ResponseWriter, r *dns.Msg) {
	return func(w dns.ResponseWriter, r *dns.Msg) {
		m := new(dns.Msg)
		m.SetReply(r)

		for _, q := range r.Question {
			fmt.Println("Inside Check: ", q)
			if ip, found := trie.Search(q.Name); found {
				rr := &dns.A{
					Hdr: dns.RR_Header{Name: q.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 3600},
					A:   ip,
				}
				m.Answer = append(m.Answer, rr)
			} else {
				upstream := "8.8.8.8:53"
				resp, err := forwardQuery(r, upstream)
				if err != nil {
					log.Printf("Error forwarding query: %s", err)
					m.SetRcode(r, dns.RcodeServerFailure)
					w.WriteMsg(m)
					return
				}
				m = resp
			}
		}
		w.WriteMsg(m)
	}
}
