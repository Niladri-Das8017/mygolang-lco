package main

import (
	"context"
	"log"
	"time"
	pb "usermgmt/usermgmt"

	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main()  {
	
	//Dial a connection to grpc Server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Dial failed: ", err)
	}
	defer conn.Close()

	//Create new Client
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//New Users
	var new_users = make(map[string]int32)
	new_users["Alice"] = 32
	new_users["Elon"] = 45

	for name, age := range new_users {
		response, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatal("Could not create user: ", err)
		}
		log.Printf(`User details:
		Name: %s
		Age: %d
		Id: %d`, response.GetName(), response.GetAge(), response.GetId())
	}

}