package dto

import (
	"fmt"
	"net"
)

type PeerIn struct {
	Address AddressIn `json:"publicAddress" validate:"required"`
}

type AddressIn struct {
	Ip   net.IP `json:"ip" validate:"required"`
	Port int    `json:"port" validate:"required"`
}

func (a *AddressIn) String() string {
	return fmt.Sprintf("%s:%d", a.Ip, a.Port)
}

type FilesIn struct {
	DistributionFiles []string `json:"distributingFiles" validate:"required"`
}
