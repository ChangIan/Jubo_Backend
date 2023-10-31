package controllerV1

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	model "changian.com/jubo/model"

	modelBase "changian.com/jubo/model/base"

	modelCommon "changian.com/jubo/model/common"

	serviceV1 "changian.com/jubo/service/v1"
)

// Controller
// Method - Get
// 查詢患者醫囑
func GetOrders(context *gin.Context) {

	result := new(modelBase.Response)

	// 參數檢查
	patientId, patientIdError := strconv.Atoi(context.Param("id"))

	if patientIdError != nil {
		log.Fatal(patientIdError)

		result.ErrorCode = 1
		result.ErrorMessage = "查詢失敗，請稍候執行"

		context.JSON(200, result)

		return
	}

	// 沒有 JWT 先給 OperatorID = 1
	operatorClaims := modelCommon.OperatorClaims{
		OperatorID: 1,
	}

	orders, err := serviceV1.QueryOrders(&operatorClaims, patientId)

	if err != nil {
		log.Fatal(err)

		result.ErrorCode = 1
		result.ErrorMessage = "查詢失敗，請稍候執行"
	} else {
		result.ErrorCode = 0
		result.ErrorMessage = "查詢成功"
		result.Data = orders
	}

	context.JSON(200, result)
}

// Controller
// Method - Post
// 新增患者醫囑
func PostOrders(context *gin.Context) {

	result := new(modelBase.Response)

	var orders model.Orders

	if bindError := context.ShouldBindJSON(&orders); bindError != nil {
		log.Fatal(bindError)

		result.ErrorCode = 1
		result.ErrorMessage = "新增失敗，請稍候執行"

		context.JSON(200, result)

		return
	}

	// 沒有 JWT 先給 OperatorID = 1
	operatorClaims := modelCommon.OperatorClaims{
		OperatorID: 1,
	}

	err := serviceV1.AddOrders(&operatorClaims, &orders)

	if err != nil {
		log.Fatal(err)

		result.ErrorCode = 1
		result.ErrorMessage = "新增失敗，請稍候執行"
	} else {
		result.ErrorCode = 0
		result.ErrorMessage = "新增成功"
	}

	context.JSON(200, result)
}

// Controller
// Method - Put
// 更新患者醫囑
func PutOrders(context *gin.Context) {

	result := new(modelBase.Response)

	var orders model.Orders

	if bindError := context.ShouldBindJSON(&orders); bindError != nil {
		log.Fatal(bindError)

		result.ErrorCode = 1
		result.ErrorMessage = "更新失敗，請稍候執行"

		context.JSON(200, result)

		return
	}

	// 沒有 JWT 先給 OperatorID = 1
	operatorClaims := modelCommon.OperatorClaims{
		OperatorID: 1,
	}

	err := serviceV1.ModifyOrders(&operatorClaims, &orders)

	if err != nil {
		log.Fatal(err)

		result.ErrorCode = 1
		result.ErrorMessage = "更新失敗，請稍候執行"
	} else {
		result.ErrorCode = 0
		result.ErrorMessage = "更新成功"
	}

	context.JSON(200, result)
}

// Controller
// Method - Delete
// 刪除患者醫囑
func DeleteOrders(context *gin.Context) {

	result := new(modelBase.Response)

	// 參數檢查
	orderId, orderIdError := strconv.Atoi(context.Param("id"))

	if orderIdError != nil {
		log.Fatal(orderIdError)

		result.ErrorCode = 1
		result.ErrorMessage = "查詢失敗，請稍候執行"

		context.JSON(200, result)

		return
	}

	// 沒有 JWT 先給 OperatorID = 1
	operatorClaims := modelCommon.OperatorClaims{
		OperatorID: 1,
	}

	err := serviceV1.DeleteOrders(&operatorClaims, orderId)

	if err != nil {
		log.Fatal(err)

		result.ErrorCode = 1
		result.ErrorMessage = "刪除失敗，請稍候執行"
	} else {
		result.ErrorCode = 0
		result.ErrorMessage = "刪除成功"
	}

	context.JSON(200, result)
}
