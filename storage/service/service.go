package service

import (
	"context"
	pb "garden-managment-service/generated/gardenManagement"
	postgres "garden-managment-service/storage/postgres"
)

type GardenManagementServer struct {
	pb.UnimplementedGardenManagementServer
	Garden *postgres.GardenManagementRepo
	Plant  *postgres.GardenPlantManagementRepo
}

func (s *GardenManagementServer) CreateGarden(ctx context.Context, in *pb.CreateGardenRequest) (*pb.CreateGardenResponse, error) {
	resp, err := s.Garden.CreateGarden(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *GardenManagementServer) ViewGarden(ctx context.Context, in *pb.ViewGardenRequest) (*pb.ViewGardenResponse, error) {
	resp, err := s.Garden.ViewGarden(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *GardenManagementServer) UpdateGarden(ctx context.Context, in *pb.UpdateGardenRequest) (*pb.UpdateGardenResponse, error) {
	resp, err := s.Garden.UpdateGarden(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *GardenManagementServer) DeleteGarden(ctx context.Context, in *pb.DeleteGardenRequest) (*pb.DeleteGardenResponse, error) {
	resp, err := s.Garden.DeleteGarden(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *GardenManagementServer) ViewUserGardens(ctx context.Context, in *pb.ViewUserGardensRequest) (*pb.ViewUserGardensResponse, error) {
	resp, err := s.Garden.ViewUserGardens(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *GardenManagementServer) AddPlanttoGarden(ctx context.Context, in *pb.AddPlanttoGardenRequest) (*pb.AddPlanttoGardenResponse, error) {
	resp, err := s.Plant.AddPlanttoGarden(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *GardenManagementServer) ViewGardenPlants(ctx context.Context, in *pb.ViewGardenPlantsRequest) (*pb.ViewGardenPlantsResponse, error) {
	resp, err := s.Plant.ViewGardenPlants(in.GardenId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *GardenManagementServer) UpdatePlant(ctx context.Context, in *pb.UpdatePlantRequest) (*pb.UpdatePlantResponse, error) {
	resp, err := s.Plant.UpdatePlant(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *GardenManagementServer) DeletePlant(ctx context.Context, in *pb.DeletePlantRequest) (*pb.DeletePlantResponse, error) {
	resp, err := s.Plant.DeletePlant(in.Id)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *GardenManagementServer) AddPlantCareLog(ctx context.Context, in *pb.AddPlantCareLogResquest) (*pb.AddPlantCareLogResponse, error) {
	resp, err := s.Plant.AddPlantCareLog(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *GardenManagementServer) ViewPlantCareLogs(ctx context.Context, in *pb.ViewPlantCareLogsRequest) (*pb.ViewPlantCareLogsResponse, error) {
	resp, err := s.Plant.ViewPlantCareLogs(in.PlantId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
