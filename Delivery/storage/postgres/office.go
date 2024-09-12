package postgres

import (
	"database/sql"
	"delivery/genproto/delivery"
	"fmt"
)

func (stg *DeliveryStorage) CreateOffice(req *delivery.CreateOfficeRequest) (*delivery.InfoResponse, error) {
	query := `INSERT INTO offices (name, address, phone_number, email) VALUES ($1, $2, $3, $4)`
	_, err := stg.db.Exec(query, req.Name, req.Address, req.PhoneNumber, req.Email)
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}
	return &delivery.InfoResponse{Success: true, Message: "Office created successfully"}, nil
}

func (stg *DeliveryStorage) GetOffice(req *delivery.GetOfficeRequest) (*delivery.OfficeResponse, error) {
	query := `SELECT office_id, name, address, phone_number, email FROM offices WHERE office_id = $1`
	row := stg.db.QueryRow(query, req.OfficeId)

	var office delivery.OfficeResponse
	err := row.Scan(&office.OfficeId, &office.Name, &office.Address, &office.PhoneNumber, &office.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("office with ID %s not found", req.OfficeId)
		}
		return nil, err
	}

	return &office, nil
}

func (stg *DeliveryStorage) GetAllOffices(req *delivery.GetAllOfficesRequest) (*delivery.GetAllOfficesResponse, error) {
	offset := (req.Page - 1) * req.Limit
	query := `SELECT office_id, name, address, phone_number, email FROM offices LIMIT $1 OFFSET $2`
	rows, err := stg.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offices []*delivery.OfficeResponse
	for rows.Next() {
		var office delivery.OfficeResponse
		if err := rows.Scan(&office.OfficeId, &office.Name, &office.Address, &office.PhoneNumber, &office.Email); err != nil {
			return nil, err
		}
		offices = append(offices, &office)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	countQuery := `SELECT COUNT(*) FROM offices`
	var totalCount int
	err = stg.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	return &delivery.GetAllOfficesResponse{
		Offices:    offices,
		TotalCount: int32(totalCount),
	}, nil
}

func (stg *DeliveryStorage) UpdateOffice(req *delivery.UpdateOfficeRequest) (*delivery.InfoResponse, error) {
	query := `UPDATE offices SET name = $1, address = $2, phone_number = $3, email = $4 WHERE office_id = $5`
	result, err := stg.db.Exec(query, req.Name, req.Address, req.PhoneNumber, req.Email, req.OfficeId)
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	if rowsAffected == 0 {
		return &delivery.InfoResponse{Success: false, Message: "Office not found"}, nil
	}

	return &delivery.InfoResponse{Success: true, Message: "Office updated successfully"}, nil
}

func (stg *DeliveryStorage) DeleteOffice(req *delivery.DeleteOfficeRequest) (*delivery.InfoResponse, error) {
	query := `DELETE FROM offices WHERE office_id = $1`
	result, err := stg.db.Exec(query, req.OfficeId)
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	if rowsAffected == 0 {
		return &delivery.InfoResponse{Success: false, Message: "Office not found"}, nil
	}

	return &delivery.InfoResponse{Success: true, Message: "Office deleted successfully"}, nil
}
