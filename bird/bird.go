package bird

import (
	"fmt"
	"net"
	"regexp"

	"errors"

	"github.com/czerwonk/asn_lookup/asn"
	"github.com/czerwonk/bird_socket"
)

type Bird struct {
	bird4SocketPath string
	bird6SocketPath string
	regex           *regexp.Regexp
}

func New(bird4Socket, bird6Socket string) *Bird {
	return &Bird{bird4SocketPath: bird4Socket, bird6SocketPath: bird6Socket, regex: regexp.MustCompile("([^\\s]+)[\\s]+(?:via|unreachable)[^\\n]+\\[AS(\\d+)")}
}

func (b *Bird) GetAs(a string) (*asn.AutonomousSystem, error) {
	as := asn.NewAs(a)
	var err error

	if len(b.bird4SocketPath) > 0 {
		as.Ipv4, err = b.getPrefixes(b.bird4SocketPath, a)
		if err != nil {
			return nil, err
		}
	}

	if len(b.bird6SocketPath) > 0 {
		as.Ipv6, err = b.getPrefixes(b.bird6SocketPath, a)
		if err != nil {
			return nil, err
		}
	}

	return as, nil
}

func (b *Bird) GetAsByIP(ip net.IP) (*asn.AutonomousSystem, error) {
	a := ""
	var err error

	if ip.To4() != nil && len(b.bird4SocketPath) > 0 {
		a, err = b.getSourceAsn(b.bird4SocketPath, ip)
		if err != nil {
			return nil, err
		}
	}

	if ip.To4() == nil && len(b.bird6SocketPath) > 0 {
		a, err = b.getSourceAsn(b.bird6SocketPath, ip)
		if err != nil {
			return nil, err
		}
	}

	if len(a) == 0 {
		return nil, errors.New(fmt.Sprintf("Source ASN for IP %s not found.", ip.String()))
	}

	return b.GetAs(a)
}

func (b *Bird) getSourceAsn(socketPath string, ip net.IP) (string, error) {
	qry := fmt.Sprintf("show route for %s", ip.String())
	resp, err := birdsocket.Query(socketPath, qry)
	if err != nil {
		return "", err
	}

	return b.parseSource(resp[5:]), nil
}

func (b *Bird) getPrefixes(socketPath string, asn string) ([]string, error) {
	qry := fmt.Sprintf("show route filter { if bgp_path.last = %s then accept; reject; }", asn)
	resp, err := birdsocket.Query(socketPath, qry)
	if err != nil {
		return nil, err
	}

	return b.parsePrefixes(resp[5:]), nil
}

func (b *Bird) parsePrefixes(resp []byte) []string {
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

func (b *Bird) parseSource(resp []byte) string {
	matches := b.regex.FindAllStringSubmatch(string(resp), -1)
	if matches == nil {
		return ""
	}

	top := 0
	a := ""
	for _, m := range matches {
		_, p, _ := net.ParseCIDR(m[1])

		if p != nil {
			_, l := p.Mask.Size()
			if l > top {
				a = m[2]
				top = l
			}
		}
	}

	return a
}
