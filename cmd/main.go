package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	sql "google.golang.org/api/sqladmin/v1beta4"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

func main() {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, sqladmin.SqlserviceAdminScope)
	if err != nil {
		fmt.Print("error")
	}
	sqlAdmin, _ := sql.New(client)
	cr := &sql.InstancesCloneRequest{
		CloneContext: &sql.CloneContext{
			DestinationInstanceName: "cloned-big",
		},
	}
	fmt.Print("2 \n")
	op, err := sqlAdmin.Instances.Clone("gke-verification", "opssight-4-7-2018-06-06-19-23-34", cr).Do()
	for {
		inst, err := sqlAdmin.Instances.Get("gke-verification", "cloned-big").Do()
		fmt.Print(fmt.Sprintf("%v %v", op, err))
		fmt.Print(fmt.Sprintf("instance: %v , disk: %v", inst.State, inst.CurrentDiskSize))
		time.Sleep(10 * time.Second)
	}
}
