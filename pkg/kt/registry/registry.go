// +build !windows

package registry

func SetGlobalProxy(port int, config *ProxyConfig) error {
	return nil
}

func CleanGlobalProxy(config *ProxyConfig) {
}

func SetHttpProxyEnvironmentVariable(protocol string, port int, config *ProxyConfig) error {
	return nil
}

func CleanHttpProxyEnvironmentVariable(config *ProxyConfig) {

}

func ResetGlobalProxyAndEnvironmentVariable() {

}
