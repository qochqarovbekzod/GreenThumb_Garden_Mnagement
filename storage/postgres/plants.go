package postgres

import (
	"database/sql"
	"fmt"
	pb "garden-managment-service/generated/gardenManagement"
	"garden-managment-service/pkg"
)

type GardenPlantManagementRepo struct {
	DB *sql.DB
}

func NewGardenPlantManagementRepo(db *sql.DB) *GardenPlantManagementRepo {
	return &GardenPlantManagementRepo{DB: db}
}

func (p *GardenPlantManagementRepo) AddPlanttoGarden(plant *pb.AddPlanttoGardenRequest) (*pb.AddPlanttoGardenResponse, error) {
	_, err := p.DB.Exec(`
		INSERT INTO plants (
			id,
			garden_id,
			species,
			quantity,
			planting_date,
			status
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
	`, plant.Id, plant.GardenId, plant.Species, plant.Quantity, plant.PlantingDate, plant.Status)

	if err != nil {
		return &pb.AddPlanttoGardenResponse{Success: false}, err
	}

	return &pb.AddPlanttoGardenResponse{Success: true}, nil
}

func (p *GardenPlantManagementRepo) ViewGardenPlants(gardenID string) (*pb.ViewGardenPlantsResponse, error) {
	var plants []*pb.Plant

	rows, err := p.DB.Query(`
		SELECT
			p.id,
			p.garden_id,
			p.species,
			p.quantity,
			p.planting_date,
			p.status
		FROM
			plants as p
		INNER JOIN
			gardens as g ON p.garden_id = g.id
		WHERE
			g.deleted_at = 0 and p.deleted_at = 0 and g.id = $1
	`, gardenID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var plant pb.Plant

		err = rows.Scan(&plant.Id, &plant.GardenId, &plant.Species, &plant.Quantity, &plant.PlantingDate, &plant.Status)

		if err != nil {
			return nil, err
		}

		plants = append(plants, &plant)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.ViewGardenPlantsResponse{Plants: plants}, nil
}

func (p *GardenPlantManagementRepo) UpdatePlant(plant *pb.UpdatePlantRequest) (*pb.UpdatePlantResponse, error) {
	params := make(map[string]interface{})
	var query = "UPDATE plants as p SET "

	if plant.GardenId != "" {
		query += "garden_id = :gardenId, "
		params["gardenId"] = plant.GardenId
	}
	if plant.Species != "" {
		query += "species = :species, "
		params["species"] = plant.Species
	}
	if plant.Quantity != 0 {
		query += "quantity = :quantity, "
		params["quantity"] = plant.Quantity
	}
	if plant.PlantingDate != "" {
		query += "planting_date = :planting_date, "
		params["planting_date"] = plant.PlantingDate
	}
	if plant.Status != "" {
		query += "status = :status, "
		params["status"] = plant.Status
	}

	query += `
			updated_at = CURRENT_TIMESTAMP 
		FROM 
			gardens as g 
		WHERE 
			p.garden_id = g.id and g.deleted_at = 0 AND p.deleted_at = 0 AND p.id = :id
		`

	params["id"] = plant.Id
	query, args := pkg.ReplaceQueryParams(query, params)

	res, err := p.DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return &pb.UpdatePlantResponse{Success: false}, fmt.Errorf("no rows affected, plant with id %s not found", plant.Id)
	}
	return &pb.UpdatePlantResponse{Success: true}, nil

}
func (p *GardenPlantManagementRepo) DeletePlant(id string) (*pb.DeletePlantResponse, error) {
	res, err := p.DB.Exec(`
	UPDATE plants AS p
SET deleted_at = EXTRACT(epoch FROM CURRENT_TIMESTAMP)::INT
FROM gardens AS g
WHERE g.id = p.garden_id 
  AND g.deleted_at = 0 
  AND p.deleted_at = 0 
  AND p.id = $1
	`, id)

	if err != nil {
		return &pb.DeletePlantResponse{Success: false}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return &pb.DeletePlantResponse{Success: false}, fmt.Errorf("no rows affected, plant with id %s not found", id)
	}
	return &pb.DeletePlantResponse{Success: true}, nil
}
func (p *GardenPlantManagementRepo) AddPlantCareLog(careLog *pb.AddPlantCareLogResquest) (*pb.AddPlantCareLogResponse, error) {
	_, err := p.DB.Exec(`
		INSERT INTO care_logs (
			id,
			plant_id,
			action,
			notes
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		)
	`, careLog.Id, careLog.PlantId, careLog.Action, careLog.Notes)

	if err != nil {
		return &pb.AddPlantCareLogResponse{Success: true}, err
	}
	return &pb.AddPlantCareLogResponse{Success: true}, nil
}
func (p *GardenPlantManagementRepo) ViewPlantCareLogs(plantID string) (*pb.ViewPlantCareLogsResponse, error) {
	var carLogs []*pb.CareLog
	rows, err := p.DB.Query(`
		SELECT
			c.id,
			plant_id,
			c.action,
			c.notes
		FROM
			care_logs AS c
		INNER JOIN
			plants AS p ON c.plant_id = p.id
		INNER JOIN
			gardens AS g ON p.garden_id = g.id
		WHERE
			g.deleted_at = 0 AND p.deleted_at = 0 and p.id=$1
	`, plantID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var carLog pb.CareLog

		err = rows.Scan(&carLog.Id, &carLog.PlantId, &carLog.Action, &carLog.Notes)

		if err != nil {
			return nil, err
		}

		carLogs = append(carLogs, &carLog)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.ViewPlantCareLogsResponse{CoreLogs: carLogs}, nil
}
