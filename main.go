package main

import (
	"context"
	"fmt"

	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/gophercloud/utils/v2/openstack/clientconfig"
)

func main() {
	opts := new(clientconfig.ClientOpts)
	// We could configure the cloud manually but we don't. Instead we'll leave it unset causing gophercloud to load the
	// cloud from the 'OS_CLOUD' environment variable. We could also simplify this further and pass 'nil' below instead
	// of generating and passing a 'ClientOpts' object but this shows how you _could_ configure things if you so chose.
	// opts.Cloud = "devstack-admin"

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
