package serviceV1

import (
	"encoding/json"
	"fmt"
	"log"

	common "changian.com/jubo/model/common"

	model "changian.com/jubo/model"

	repositoryV1 "changian.com/jubo/repository/v1"
)

// 查詢患者醫囑
func QueryOrders(user *common.OperatorClaims, patientId int) (result []model.Orders, err error) {

	defer func() {
		transactionLog := model.TransactionLog{
			OperatorID: user.OperatorID,
			Content:    fmt.Sprintf("查詢患者醫囑 Patient_Id : %d", patientId),
		}

		if err == nil {
			// 紀錄痕跡
			err = repositoryV1.InsertTransactionLog(&transactionLog)
		}
	}()

	result, err = repositoryV1.QueryOrders(patientId)

	return result, err
}

// 新增患者醫囑
func AddOrders(user *common.OperatorClaims, orders *model.Orders) (err error) {

	defer func() {

		outByte, outError := json.Marshal(orders)

		if outError != nil {
			log.Fatal(outError)
		}

		transactionLog := model.TransactionLog{
			OperatorID: user.OperatorID,
			Content:    fmt.Sprintf("新增患者醫囑 : %s", string(outByte)),
		}

		if err == nil {
			// 紀錄痕跡
			err = repositoryV1.InsertTransactionLog(&transactionLog)
		}
	}()

	err = repositoryV1.InsertOrders(orders)

	return err
}

// 更新患者醫囑
func ModifyOrders(user *common.OperatorClaims, orders *model.Orders) (err error) {

	defer func() {

		outByte, outError := json.Marshal(orders)

		if outError != nil {
			log.Fatal(outError)
		}

		transactionLog := model.TransactionLog{
			OperatorID: user.OperatorID,
			Content:    fmt.Sprintf("更新患者醫囑 : %s", string(outByte)),
		}

		if err == nil {
			// 紀錄痕跡
			err = repositoryV1.InsertTransactionLog(&transactionLog)
		}
	}()

	err = repositoryV1.UpdateOrders(orders)

	return err
}

// 刪除患者醫囑
func DeleteOrders(user *common.OperatorClaims, orderId int) (err error) {

	defer func() {

		transactionLog := model.TransactionLog{
			OperatorID: user.OperatorID,
			Content:    fmt.Sprintf("刪除患者醫囑 OrderId : %d", orderId),
		}

		if err == nil {
			// 紀錄痕跡
			err = repositoryV1.InsertTransactionLog(&transactionLog)
		}
	}()

	err = repositoryV1.DeleteOrders(orderId)

	return err
}
