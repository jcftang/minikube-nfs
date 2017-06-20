package nfsd

import (
	"fmt"
	"net"
	"strings"
)

const ExportsFile = "/etc/exports"

type Hosts struct {
	Names   []string
	Network *net.IPNet
}

func (h Hosts) String() string {
	var s string

	if len(h.Names) > 0 {
		s = strings.Join(h.Names, " ")
	}

	if h.Network != nil {
		if h.Network.Mask != nil {
			s = fmt.Sprintf("-network %s -mask %s", h.Network.IP, net.IP(h.Network.Mask))
		} else {
			s = fmt.Sprintf("-network %s", h.Network.IP)
		}
	}

	return s
}

type Export struct {
	Directories []string
	Flags       []string
	Hosts       Hosts
}

func (e Export) String() string {
	var (
		directories = strings.Join(e.Directories, " ")
		flags       = strings.Join(e.Flags, " ")
		hosts       = e.Hosts
	)

	return fmt.Sprintf("%s %s %v", directories, flags, hosts)
}