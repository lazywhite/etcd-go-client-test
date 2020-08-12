package main

import (
	"context"
	"fmt"
	//client "github.com/coreos/etcd/clientv3"
	client "github.com/coreos/etcd/clientv3"
)

var eps = []string{
	"http://localhost:12379",
	"http://localhost:22379",
	"http://localhost:32379",
}

func main() {
	cfg := client.Config{
		Endpoints: eps,
	}

	c, err := client.New(cfg)
	if err != nil{
		fmt.Println("failed to create client")
	}

	fmt.Println(c.Endpoints())
	defer c.Close()

	mResp, err := c.MemberList(context.TODO())
	fmt.Println(mResp.Members)

	_, err = c.Put(context.Background(), "/ns/bob", "value")
	if err != nil{
		fmt.Println(err.Error())
	}
	_, err = c.Put(context.Background(), "/ns/alice", "xx")
	_, err = c.Put(context.Background(), "/ns/alice/age", "xx")
	var resp *client.GetResponse
	resp, err = c.Get(context.Background(), "/ns/bob")
	fmt.Println(resp.Kvs)

	// no directory, use prefix instead
	resp, err = c.Get(context.Background(), "/ns", client.WithPrefix(), client.WithSort(client.SortByKey, client.SortDescend))
	fmt.Println(resp.Kvs)

}
