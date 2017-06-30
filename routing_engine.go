package main

import "github.com/czerwonk/asn_lookup/asn"

type RoutingEngine interface {
	GetAs(asn string) (*asn.AutonomousSystem, error)
}
