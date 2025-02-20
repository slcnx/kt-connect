package cmd

import (
	"github.com/alibaba/kt-connect/pkg/kt/options"
	"github.com/alibaba/kt-connect/pkg/kt/util"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd/api"
)

// GlobalOptions ...
type GlobalOptions struct {
	// global
	Labels    string
	Image     string
	Debug     bool
	currentNs string
	Timeout   int

	// common
	args                   []string
	userSpecifiedNamespace string
	restConfig             *rest.Config
	rawConfig              api.Config
	clientset              kubernetes.Interface
}

// ExchangeOptions ...
type ExchangeOptions struct {
	GlobalOptions

	configFlags *genericclioptions.ConfigFlags
	genericclioptions.IOStreams

	// exchange
	Target string
	Expose string
}

// ConnectOptions ...
type ConnectOptions struct {
	GlobalOptions

	genericclioptions.IOStreams
	configFlags *genericclioptions.ConfigFlags

	Method     string
	Labels     string
	Proxy      int
	DisableDNS bool
	Cidr       string
	Dump2hosts string
	Port       int
	TunName    string
	TunCidr    string
}

// MeshOptions ...
type MeshOptions struct {
	GlobalOptions

	genericclioptions.IOStreams
	configFlags *genericclioptions.ConfigFlags

	// mesh
	Target  string
	Expose  string
	Version string
}

// ProvideOptions ...
type ProvideOptions struct {
	GlobalOptions

	genericclioptions.IOStreams
	configFlags *genericclioptions.ConfigFlags

	// run
	Expose   int
	External bool
	Target   string
}

func (o *GlobalOptions) transportGlobalOptions() *options.DaemonOptions {
	return &options.DaemonOptions{
		Image:      o.Image,
		Debug:      o.Debug,
		WithLabels: o.Labels,
		Namespace:  o.currentNs,
		WaitTime:   o.Timeout,
		RuntimeOptions: &options.RuntimeOptions{
			UserHome:   util.UserHome,
			AppHome:    util.KtHome,
			Clientset:  o.clientset,
			RestConfig: o.restConfig,
		},
		ConnectOptions: &options.ConnectOptions{},
	}
}
