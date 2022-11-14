package repositories

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"litetorrent-tracker/internal/dto"
	"net"
)

type Repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (pr *Repository) CreatePeer(address string) (*uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("error while generate uuid: %w", err)
	}
	_, err = pr.db.Exec("INSERT INTO peer (id, address) values ($1, $2)", id.String(), address)
	if err != nil {
		return nil, fmt.Errorf("error while create peer: %w", err)
	}
	return &id, nil
}

func (pr *Repository) UpdateDistributedFiles(peerId uuid.UUID, hashes []string) error {
	if _, err := pr.db.Exec("DELETE FROM peer_file WHERE peer_id = $1", peerId); err != nil {
		return fmt.Errorf("error while deleting files with peer: %w", err)
	}
	if _, err := pr.db.Exec("DELETE FROM file WHERE NOT EXISTS(SELECT 1 FROM peer_file pf WHERE file.hash = pf.hash)"); err != nil {
		return fmt.Errorf("error while deleting files without any peers: %w", err)
	}
	for _, hash := range hashes {
		if _, err := pr.db.Exec("INSERT INTO file (hash) VALUES ($1) ON CONFLICT DO NOTHING", hash); err != nil {
			return fmt.Errorf("error while insert file: %w", err)
		}
		if _, err := pr.db.Exec("INSERT INTO peer_file (peer_id, hash) VALUES ($1, $2)", peerId.String(), hash); err != nil {
			return fmt.Errorf("error while insert peer_file: %w", err)
		}
	}
	return nil
}

func (pr *Repository) GetPeerAddressesByFile(hash string) (*dto.AddressesOut, error) {
	rows, err := pr.db.Query("SELECT peer.address FROM peer JOIN peer_file pf on peer.id = pf.peer_id WHERE pf.hash = $1", hash)
	if err != nil {
		return nil, fmt.Errorf("error while get peer addresses: %w", err)
	}
	addressesOut := make([]dto.AddressOut, 0)
	for rows.Next() {
		var address string
		if err := rows.Scan(&address); err != nil {
			return nil, fmt.Errorf("error while scan address: %w", err)
		}
		ip, err := net.ResolveUDPAddr("udp", address)
		if err != nil {
			return nil, fmt.Errorf("error while parse address: %w", err)
		}
		addressesOut = append(addressesOut, dto.AddressOut{Ip: ip.IP, Port: ip.Port})
	}
	return &dto.AddressesOut{PeerAddresses: addressesOut}, nil
}

func (pr *Repository) DeletePeer(peerId uuid.UUID) error {
	_, err := pr.db.Exec("DELETE FROM peer WHERE id = $1", peerId)
	if err != nil {
		return fmt.Errorf("error while deleting peer: %w", err)
	}
	return nil
}

func (pr *Repository) IsPeerExist(peerId uuid.UUID) (bool, error) {
	var count int
	err := pr.db.QueryRow("SELECT count(id) FROM peer WHERE id = $1", peerId.String()).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error while check is peer exist: %w", err)
	}
	return count > 0, nil
}
