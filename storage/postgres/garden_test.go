package postgres

import (
	"fmt"
	pb "garden-managment-service/generated/gardenManagement"
	"reflect"
	"testing"
)

func TestCreateGarden(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	garden := NewGardenManagementRepo(db)

	respcreate, err := garden.CreateGarden(&pb.CreateGardenRequest{
		Id:      "92a97aa1-a9d3-4b30-b24e-3681995ee86d",
		UserId:  "00e6f3c0-d302-4877-8e35-984ab1ae3c81",
		Name:    "ANY",
		AreaSqm: 12.2})

	if err != nil {
		fmt.Println(err)
	}
	waitcreate := pb.CreateGardenResponse{
		Success: true,
	}

	if !reflect.DeepEqual(respcreate, &waitcreate) {
		t.Errorf("have %v , wont %v", respcreate, &waitcreate)
	}
}

func TestViewGarden(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	garden := NewGardenManagementRepo(db)
	respview, err := garden.ViewGarden(&pb.ViewGardenRequest{
		Id: "92a97aa1-a9d3-4b30-b24e-3681995ee86d",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitview := pb.ViewGardenResponse{
		Id:      "92a97aa1-a9d3-4b30-b24e-3681995ee86d",
		UserId:  "00e6f3c0-d302-4877-8e35-984ab1ae3c81",
		Name:    "ANY",
		AreaSqm: 12.2,
	}
	if !reflect.DeepEqual(respview, &waitview) {
		t.Errorf("have %v , wont %v", respview, &waitview)
	}
}

func TestUpdateGarden(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	garden := NewGardenManagementRepo(db)
	respupdate, err := garden.UpdateGarden(&pb.UpdateGardenRequest{
		Id:      "92a97aa1-a9d3-4b30-b24e-3681995ee86d",
		UserId:  "00e6f3c0-d302-4877-8e35-984ab1ae3c81",
		Name:    "ANY",
		Type:    "",
		AreaSqm: 12.0})
	if err != nil {
		fmt.Println(err)
	}
	waitupdate := pb.UpdateGardenResponse{
		Success: true,
	}
	if !reflect.DeepEqual(respupdate, &waitupdate) {
		t.Errorf("have %v , wont %v", respupdate, &waitupdate)
	}
}

func TestDeleteGarden(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	garden := NewGardenManagementRepo(db)
	respdelete, err := garden.DeleteGarden(&pb.DeleteGardenRequest{
		Id: "92a97aa1-a9d3-4b30-b24e-3681995ee86d",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitdelete := pb.DeleteGardenResponse{
		Success: true,
	}
	if !reflect.DeepEqual(respdelete, &waitdelete) {
		t.Errorf("have %v , wont %v", respdelete, &waitdelete)
	}
}

func TestViewUserGarden(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	garden := NewGardenManagementRepo(db)
	resusergarden, err := garden.ViewUserGardens(&pb.ViewUserGardensRequest{UserId: "00e6f3c0-d302-4877-8e35-984ab1ae3c81"})
	if err != nil {
		fmt.Println(err)
	}

	waitusergarden := pb.ViewUserGardensResponse{
		Gardens: []*pb.Garden{
			{Id: "3c02ad1a-d7f6-438c-82fc-544b7fb1cf1e",
				UserId:  "00e6f3c0-d302-4877-8e35-984ab1ae3c81",
				Name:    "First",
				Type:    "rooftop",
				AreaSqm: 12.0},
		},
	}
	if !reflect.DeepEqual(resusergarden, &waitusergarden) {
		t.Errorf("have %v , wont %v", resusergarden, &waitusergarden)
	}

}
