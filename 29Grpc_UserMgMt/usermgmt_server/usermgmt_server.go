package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	pb "usermgmt/usermgmt"

	"google.golang.org/grpc"
)

const port = ":50051"

//this is our implimentation of grpc service
type UserMgmtServer struct {
	pb.UnimplementedUserManagementServer
}

//Here we can see server recieving Context and a protobuff file carring data of a new user from client, i.e. name & age.
//Server gnerates an random id for new user and return user packed in protobuff with name & age recieved from client and user_id gnerattted and assigned by server to that user, Else an error
func (s *UserMgmtServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {

	log.Printf("Recieved: %v\n", in.GetName())

	//Generate Random Id
	var user_id int32 = int32(rand.Intn(1000))

	//returning &pb.User
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen0: ", err)
	}

	//Initialize new Server
	s := grpc.NewServer()

	//Regester the server as a new grpc service
	pb.RegisterUserManagementServer(s, &UserMgmtServer{})
	log.Println("server listening at ", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}

}
