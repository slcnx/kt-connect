package exec

import (
	"github.com/alibaba/kt-connect/pkg/kt/exec/kubectl"
	"github.com/alibaba/kt-connect/pkg/kt/exec/portforward"
	"github.com/alibaba/kt-connect/pkg/kt/exec/ssh"
	"github.com/alibaba/kt-connect/pkg/kt/exec/sshchannel"
	"github.com/alibaba/kt-connect/pkg/kt/exec/sshuttle"
	"github.com/alibaba/kt-connect/pkg/kt/exec/tunnel"
)

// CliInterface ...
type CliInterface interface {
	Kubectl() kubectl.CliInterface
	Sshuttle() sshuttle.CliInterface
	SSH() ssh.CliInterface
	Tunnel() tunnel.CliInterface
	SshChannel() sshchannel.Channel
	PortForward() portforward.CliInterface
}

// Cli ...
type Cli struct {
	KubeOptions []string
	TunName     string
	SourceIP    string
	DestIP      string
	// MaskLen the net mask length of tun cidr
	MaskLen string
}

// PortForward ...
func (c *Cli) PortForward() portforward.CliInterface {
	return &portforward.Cli{}
}

// SshChannel ...
func (c *Cli) SshChannel() sshchannel.Channel {
	return &sshchannel.SSHChannel{}
}

// Kubectl ...
func (c *Cli) Kubectl() kubectl.CliInterface {
	return &kubectl.Cli{KubeOptions: c.KubeOptions}
}

// Sshuttle ...
func (c *Cli) Sshuttle() sshuttle.CliInterface {
	return &sshuttle.Cli{}
}

// SSH ...
func (c *Cli) SSH() ssh.CliInterface {
	return &ssh.Cli{}
}

func (c *Cli) Tunnel() tunnel.CliInterface {
	return &tunnel.Cli{
		TunName:  c.TunName,
		SourceIP: c.SourceIP,
		DestIP:   c.DestIP,
		MaskLen:  c.MaskLen,
	}
}
