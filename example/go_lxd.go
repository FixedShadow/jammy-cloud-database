package main

import (
	"context"
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
	instanceServer, err := lxd.ConnectLXD("https://192.168.2.99:8443", &param)
	if err != nil {
		panic(err)
	}
	instance := api.InstancesPost{}
	instance.Name = "test-mysql-instance01"
	instance.Type = "virtual-machine"
	instance.Devices = map[string]map[string]string{
		"root": {
			"path": "/",
			"pool": "new-storage-dir",
			"size": "12GiB",
			"type": "disk",
		},
	}
	instance.Config = map[string]string{
		"limits.cpu":    "2",
		"limits.memory": "2GiB",
	}
	instance.Source = api.InstanceSource{
		Type:        "image",
		BaseImage:   "913ae47e658993fb7d5e89995ffe91b2194c14843aaa9032e69ec8737055db75", //image fingerprint
		Fingerprint: "913ae47e658993fb7d5e89995ffe91b2194c14843aaa9032e69ec8737055db75", //image fingerprint
	}
	op, err := instanceServer.CreateInstance(instance)
	if err != nil {
		panic(err)
	}
	fmt.Println("operation id: ", op.Get().ID)
	fmt.Println("operation status code: ", op.Get().StatusCode.String())
	fmt.Println("operation err: ", op.Get().Err)

	err = op.WaitContext(context.Background()) //wait for the operation to complete
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
