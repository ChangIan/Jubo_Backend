package repositoryV1

import (
	"fmt"
	"time"

	global "changian.com/jubo/startup"

	model "changian.com/jubo/model"
)

// 查詢醫囑
func QueryOrders(patientId int) (result []model.Orders, err error) {

	defer func() {
		fmt.Println("Query Orders ", time.Now().Format("2006-01-02 15:03:04"))
	}()

	rows, err := global.Postgresql.Query(
		`SELECT 
		Id, 
		Patient_Id, 
		Message, 
		CreateTime,
		ModifyTime
		FROM Orders
		WHERE Patient_Id = $1
		AND IsDelete = 0;`,
		patientId)

	if err != nil {
		return nil, err
	}

	result = make([]model.Orders, 0)

	for rows.Next() {

		var obj model.Orders

		err := rows.Scan(
			&obj.Id,
			&obj.Patient_Id,
			&obj.Message,
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

// 新增醫囑
func InsertOrders(input *model.Orders) (err error) {

	defer func() {
		fmt.Println("Insert Orders ", time.Now().Format("2006-01-02 15:03:04"))
	}()

	_, err = global.Postgresql.Exec(
		`INSERT INTO Orders 
		(Patient_Id, Message) 
		VALUES 
		($1, $2);`,
		input.Patient_Id,
		input.Message,
	)

	return err
}

// 更新醫囑
func UpdateOrders(input *model.Orders) (err error) {

	defer func() {
		fmt.Println("Update Orders ", time.Now().Format("2006-01-02 15:03:04"))
	}()

	_, err = global.Postgresql.Exec(
		`UPDATE Orders 
		SET Message = $1, ModifyTime = NOW()
		WHERE Id = $2;`,
		input.Message,
		input.Id,
	)

	return err
}

// 刪除醫囑
// 只更新是否刪除的狀態，資料不會直接抹去痕跡
func DeleteOrders(orderId int) (err error) {

	defer func() {
		fmt.Println("Delete Orders ", time.Now().Format("2006-01-02 15:03:04"))
	}()

	_, err = global.Postgresql.Exec(
		`UPDATE Orders 
		SET IsDelete = 1, ModifyTime = NOW()
		WHERE Id = $1;`,
		orderId,
	)

	return err
}
