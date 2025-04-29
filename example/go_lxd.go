package main

import (
	"fmt"
	lxd "github.com/canonical/lxd/client"
	"github.com/canonical/lxd/shared/api"
	"os"
)

/**
this file show how to execute the lxd operation by go sdk.
https://pkg.go.dev/github.com/lxc/lxd/client
*/

var param = lxd.ConnectionArgs{InsecureSkipVerify: true}

func InitBasicParam() {
	certContent, err := os.ReadFile("my.crt")
	if err != nil {
		panic(err)
	}
	keyContent, err := os.ReadFile("my.key")
	if err != nil {
		panic(err)
	}
	param.TLSClientCert = string(certContent)
	param.TLSClientKey = string(keyContent)
}

func CreateVmInstance() {
	instanceServer, err := lxd.ConnectLXD("https://x.x.x.x:8443", &param)
	if err != nil {
		panic(err)
	}
	instance := api.InstancesPost{}
	instance.Name = "my-vm-001"
	instance.Type = "virtual-machine"
	instance.Devices = map[string]map[string]string{
		"root": {
			"path": "/",
			"pool": "default",
			"size": "12GiB",
			"type": "disk",
		},
	}
	instance.Config = map[string]string{
		"limits.cpu":    "1",
		"limits.memory": "2GiB",
	}
	instance.Source = api.InstanceSource{
		Type:        "image",
		BaseImage:   "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", //image fingerprint
		Fingerprint: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", //image fingerprint
	}
	op, err := instanceServer.CreateInstance(instance)
	if err != nil {
		panic(err)
	}
	fmt.Println("operation id: ", op.Get().ID)
	fmt.Println("operation status code: ", op.Get().StatusCode.String())
	fmt.Println("operation err: ", op.Get().Err)

	err = op.Wait() //wait for the operation to complete
	if err != nil {
		panic(err)
	}
	state := api.InstanceStatePut{
		Action:  "start",
		Timeout: -1,
	}
	op, err = instanceServer.UpdateInstanceState(instance.Name, state, "")
	if err != nil {
		panic(err)
	}
	err = op.Wait() //wait for the operation to complete.
	if err != nil {
		panic(err)
	}
}
