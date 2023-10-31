package repositoryV1

import (
	"fmt"
	"time"

	global "changian.com/jubo/startup"

	model "changian.com/jubo/model"
)

// 查詢患者
func QueryPatient(patientId int) (result []model.Patient, err error) {

	defer func() {
		fmt.Println("Query Patient ", time.Now().Format("2006-01-02 15:03:04"))
	}()

	rows, err := global.Postgresql.Query(
		`SELECT 
		Id, 
		Name, 
		Phone, 
		Email,
		Gender, 
		CreateTime,
		ModifyTime
		FROM Patient;`)

	if err != nil {
		return nil, err
	}

	result = make([]model.Patient, 0)

	for rows.Next() {

		var obj model.Patient

		err := rows.Scan(
			&obj.Id,
			&obj.Name,
			&obj.Phone,
			&obj.Email,
			&obj.Gender,
			&obj.CreateTime,
			&obj.ModifyTime,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, obj)
	}

	return result, nil
}
