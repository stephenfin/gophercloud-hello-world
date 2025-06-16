package main

import (
	"context"
	"fmt"

	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/v2/openstack/dns/v2/zones"
	"github.com/gophercloud/gophercloud/v2/openstack/keymanager/v1/secrets"
	"github.com/gophercloud/gophercloud/v2/openstack/messaging/v2/queues"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
	"github.com/gophercloud/utils/v2/openstack/clientconfig"
)

// listServers tests the Compute service, nova
func listServers() {
	opts := new(clientconfig.ClientOpts)
	client, err := clientconfig.NewServiceClient(context.TODO(), "compute", opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	pager, err := servers.List(client, nil).AllPages(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	servers, err := servers.ExtractServers(pager)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("servers:")
	for i, server := range servers {
		fmt.Printf("  server %d: id=%s\n", i, server.ID)
	}
}

// listZones tests the DNS service, designate
func listZones() {
	opts := new(clientconfig.ClientOpts)
	client, err := clientconfig.NewServiceClient(context.TODO(), "dns", opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	pager, err := zones.List(client, nil).AllPages(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	zones, err := zones.ExtractZones(pager)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("zones:")
	for i, zone := range zones {
		fmt.Printf("  zone %d: id=%s\n", i, zone.ID)
	}
}

// listSecrets tests the Key Manager service, barbican
func listSecrets() {
	opts := new(clientconfig.ClientOpts)
	client, err := clientconfig.NewServiceClient(context.TODO(), "key-manager", opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	pager, err := secrets.List(client, nil).AllPages(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	secrets, err := secrets.ExtractSecrets(pager)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("secrets:")
	for i, secret := range secrets {
		fmt.Printf("  secret %d: id=%s\n", i, secret.SecretRef)
	}
}

// listQueues tests the Messaging service, zaqar
func listQueues() {
	opts := new(clientconfig.ClientOpts)
	client, err := clientconfig.NewServiceClient(context.TODO(), "messaging", opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	pager, err := queues.List(client, nil).AllPages(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	queues, err := queues.ExtractQueues(pager)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("queues:")
	for i, queue := range queues {
		fmt.Printf("  queue %d: name=%s\n", i, queue.Name)
	}
}

// listNetworks tests the Networking service, neutron
func listNetworks() {
	opts := new(clientconfig.ClientOpts)
	client, err := clientconfig.NewServiceClient(context.TODO(), "network", opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	pager, err := networks.List(client, nil).AllPages(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	networks, err := networks.ExtractNetworks(pager)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("networks:")
	for i, network := range networks {
		fmt.Printf("  network %d: id=%s\n", i, network.ID)
	}
}

func main() {
	listServers()
	listZones()
	listSecrets()
	listQueues()
	listNetworks()
}
