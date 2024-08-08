package main

import (
	pb "garden-managment-service/generated/gardenManagement"
	"garden-managment-service/service"
	"garden-managment-service/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	gardenManagement := service.GardenManagementServer{
		Garden: postgres.NewGardenManagementRepo(db),
		Plant: postgres.NewGardenPlantManagementRepo(db),
	}

	pb.RegisterGardenManagementServer(s, &gardenManagement)

	log.Println("server is running on :50051 ...")
	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
