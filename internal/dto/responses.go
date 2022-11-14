package dto

import (
	"github.com/google/uuid"
	"net"
)

type PeerOut struct {
	Id uuid.UUID `json:"id"`
}

type AddressesOut struct {
	PeerAddresses []AddressOut `json:"peerIds"`
}

type AddressOut struct {
	Ip   net.IP `json:"ip"`
	Port int    `json:"port"`
}
