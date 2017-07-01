package bird

import (
	"fmt"
	"net"
	"regexp"

	"github.com/czerwonk/asn_lookup/asn"
	"github.com/czerwonk/bird_socket"
)

type Bird struct {
	bird4SocketPath string
	bird6SocketPath string
	regex           *regexp.Regexp
}

func New(bird4Socket, bird6Socket string) *Bird {
	return &Bird{bird4SocketPath: bird4Socket, bird6SocketPath: bird6Socket, regex: regexp.MustCompile("([^\\s]+)[\\s]+(?:via|unreachable)")}
}

func (b *Bird) GetAs(a string) (*asn.AutonomousSystem, error) {
	as := asn.NewAs(a)
	var err error

	if len(b.bird4SocketPath) > 0 {
		as.Ipv4, err = b.GetPrefixes(b.bird4SocketPath, a)
		if err != nil {
			return nil, err
		}
	}

	if len(b.bird6SocketPath) > 0 {
		as.Ipv6, err = b.GetPrefixes(b.bird6SocketPath, a)
		if err != nil {
			return nil, err
		}
	}

	return as, nil
}

func (b *Bird) GetPrefixes(socketPath string, asn string) ([]string, error) {
	qry := fmt.Sprintf("show route filter { if bgp_path.last = %s then accept; reject; }", asn)
	resp, err := birdsocket.Query(socketPath, qry)
	if err != nil {
		return nil, err
	}

	return b.parse(resp[5:]), nil
}

func (b *Bird) parse(resp []byte) []string {
	matches := b.regex.FindAllStringSubmatch(string(resp), -1)
	if matches == nil {
		return nil
	}

	res := make([]string, 0)
	for _, m := range matches {
		_, p, _ := net.ParseCIDR(m[1])
		if p != nil {
			res = append(res, p.String())
		}
	}

	return res
}
