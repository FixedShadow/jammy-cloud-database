package container

import (
	lxd "github.com/canonical/lxd/client"
)

func NewClient() lxd.InstanceServer {
	return nil
}

func NewClientWithAuth(url string, cert []byte, key []byte) (lxd.InstanceServer, error) {
	param := lxd.ConnectionArgs{InsecureSkipVerify: true}
	param.TLSClientCert = string(cert)
	param.TLSClientKey = string(key)
	return lxd.ConnectLXD(url, &param)
}
