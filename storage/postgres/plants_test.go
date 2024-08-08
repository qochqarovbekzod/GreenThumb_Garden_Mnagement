package postgres

import (
	"fmt"
	pb "garden-managment-service/generated/gardenManagement"
	"reflect"
	"testing"
)

func TestAddPlantGarden(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	plant := NewGardenPlantManagementRepo(db)

	resadd, err := plant.AddPlanttoGarden(&pb.AddPlanttoGardenRequest{
		Id:           "41278da8-9505-41f8-a338-be29e7db817b",
		GardenId:     "3c02ad1a-d7f6-438c-82fc-544b7fb1cf1e",
		Species:      "any",
		Quantity:     3,
		PlantingDate: "2024-01-01",
		Status:       "planned",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitadd := pb.AddPlanttoGardenResponse{
		Success: true,
	}
	if !reflect.DeepEqual(resadd, &waitadd) {
		t.Errorf("have %v , wont %v", resadd, &waitadd)
	}

}

func TestViewGardenPlants(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	plant := NewGardenPlantManagementRepo(db)
	resview, err := plant.ViewGardenPlants("3c02ad1a-d7f6-438c-82fc-544b7fb1cf1e")
	if err != nil {
		fmt.Println(err)
	}
	waitview := pb.ViewGardenPlantsResponse{
		Plants: []*pb.Plant{
			{
				Id:           "41278da8-9505-41f8-a338-be29e7db817b",
				GardenId:     "3c02ad1a-d7f6-438c-82fc-544b7fb1cf1e",
				Species:      "any",
				Quantity:     3,
				PlantingDate: "2024-01-01T00:00:00Z",
				Status:       "planned",
			},
		},
	}
	if !reflect.DeepEqual(resview, &waitview) {
		t.Errorf("have %v , wont %v", resview, &waitview)
	}
}

func TestUpdateGardenPlant(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	plant := NewGardenPlantManagementRepo(db)
	resupdate, err := plant.UpdatePlant(&pb.UpdatePlantRequest{
		Id:           "41278da8-9505-41f8-a338-be29e7db817b",
		GardenId:     "3c02ad1a-d7f6-438c-82fc-544b7fb1cf1e",
		Species:      "any",
		Quantity:     13,
		PlantingDate: "2020-01-01T00:00:00Z",
		Status:       "planned",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitupdate := pb.UpdatePlantResponse{
		Success: true,
	}
	if !reflect.DeepEqual(resupdate, &waitupdate) {
		t.Errorf("have %v , wont %v", resupdate, &waitupdate)
	}
}

func TestDeleteGardenPlant(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	plant := NewGardenPlantManagementRepo(db)
	resdelete, err := plant.DeletePlant("41278da8-9505-41f8-a338-be29e7db817b")
	if err != nil {
		fmt.Println(err)
	}
	waitdelete := pb.DeletePlantResponse{
		Success: true,
	}
	if !reflect.DeepEqual(resdelete, &waitdelete) {
		t.Errorf("have %v , wont %v", resdelete, &waitdelete)
	}
}

func TestAddPlantCareLog(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	plant := NewGardenPlantManagementRepo(db)
	resadd, err := plant.AddPlantCareLog(&pb.AddPlantCareLogResquest{
		Id:      "5caa1d3a-de14-44ed-92a6-744a2beaca70",
		PlantId: "41278da8-9505-41f8-a338-be29e7db817b",
		Action:  "Nma gap",
		Notes:   "tinchmi",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitadd := pb.AddPlantCareLogResponse{
		Success: true,
	}
	if !reflect.DeepEqual(resadd, &waitadd) {
		t.Errorf("have %v , wont %v", resadd, &waitadd)
	}
}

func TestViewPlantCareLog(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	plant := NewGardenPlantManagementRepo(db)
	resview, err := plant.ViewPlantCareLogs("41278da8-9505-41f8-a338-be29e7db817b")
	if err != nil {
		fmt.Println(err)
	}
	waitview := pb.ViewPlantCareLogsResponse{
		CoreLogs: []*pb.CareLog{
			{
				Id:      "5caa1d3a-de14-44ed-92a6-744a2beaca70",
				PlantId: "41278da8-9505-41f8-a338-be29e7db817b",
				Action:  "Nma gap",
				Notes:   "tinchmi",
			},
		},
	}
	if !reflect.DeepEqual(resview, &waitview) {
		t.Errorf("have %v , wont %v", resview, &waitview)
	}
}
