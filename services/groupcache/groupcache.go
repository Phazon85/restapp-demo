package groupcache

import (
	"fmt"
	"net"
	"os"

	"github.com/mailgun/groupcache"
	"go.uber.org/zap"
)

//Newservice to return a new http pool opts

//Set peers to the options

//]

const (
	addrTemplate = "http://%s:%s"
	defaultPort  = "8080"
)

type Service struct {
	logger  *zap.Logger
	Address string
	Port    string
	Pool    *groupcache.HTTPPool
}

func New(logger *zap.Logger) (*Service, error) {
	addr, err := findlocaladdress()
	if err != nil {
		return nil, err
	}

	tempAddress := fmt.Sprintf(addrTemplate, addr, defaultPort)

	return &Service{
		logger:  logger,
		Address: tempAddress,
		Port:    defaultPort,
	}, nil
}

func findlocaladdress() (string, error) {
	hostName, err := os.Hostname()
	if err != nil {
		return "", err
	}
	addrs, err := net.LookupHost(hostName)
	if err != nil {
		return "", err
	}
	if len(addrs) < 1 {
		return "", fmt.Errorf("Couldn't find an IP")
	}
	return addrs[2], nil
}

func (s *Service) SetPeers() {
	peerList := []string{s.Address}

	peers := []string{}
	if len(peers) < 1 {
		s.Pool.Set(peerList[0])

		return
	}
	for _, peer := range peers {
		tempPeer := fmt.Sprintf(addrTemplate, peer, s.Port)
		peerList = append(peerList, tempPeer)
	}
	s.Pool.Set(peerList...)
}
