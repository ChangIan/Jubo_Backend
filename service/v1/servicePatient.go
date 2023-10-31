package serviceV1

import (
	common "changian.com/jubo/model/common"

	model "changian.com/jubo/model"

	repositoryV1 "changian.com/jubo/repository/v1"
)

// 查詢患者
func QueryPatient(user *common.OperatorClaims) (result []model.Patient, err error) {

	defer func() {
		transactionLog := model.TransactionLog{
			OperatorID: user.OperatorID,
			Content:    "查詢患者",
		}

		if err == nil {
			// 紀錄痕跡
			err = repositoryV1.InsertTransactionLog(&transactionLog)
		}
	}()

	result, err = repositoryV1.QueryPatient(0)

	return result, err
}
