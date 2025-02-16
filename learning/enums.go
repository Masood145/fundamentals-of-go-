package learning

import (
	"fmt"
	"math/rand"
)

type Protocol string

const (
	TCP   Protocol = "TCP"
	UDP   Protocol = "UDP"
	HTTP  Protocol = "HTTP"
	HTTPS Protocol = "HTTPS"
)

type PacketStatus int

const (
	Accepted PacketStatus = iota
	Rejected
	Suspicoius
)

type PacketAnalyzer struct {
	TotalPackets      int
	AcceptedPackets   int
	RejectedPackets   int
	Suspiciouspackets int
}

type Packet struct {
	ID            int
	Protocol      Protocol
	SourceIP      string
	Destinationip string
	Size          int
}

func (ps PacketStatus) String() string {
	return [...]string{"Accepted", "Rejected", "Suspicious"}[ps]

}
func (pa *PacketAnalyzer) Analyze(p Packet) PacketStatus {
	pa.TotalPackets++

	if p.Protocol == HTTP {
		pa.Suspiciouspackets++
		return Suspicoius
	}
	if p.Size > 1500 {
		pa.RejectedPackets++
		return Rejected

	}
	pa.AcceptedPackets++
	return Accepted
}

func PacketGenerator() Packet {
	Proctocol := []Protocol{"TCP", "UDP", "HTTP", "HTTPS"}
	return Packet{

		ID:            rand.Intn(10000),
		Protocol:      Proctocol[rand.Intn(len(Proctocol))],
		SourceIP:      fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)),
		Destinationip: fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)),
		Size:          rand.Intn(2000),
	}

}
