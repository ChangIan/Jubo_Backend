package repositoryV1

import (
	"fmt"
	"time"

	global "changian.com/jubo/startup"

	model "changian.com/jubo/model"
)

// 新增交易紀錄
func InsertTransactionLog(input *model.TransactionLog) (err error) {

	defer func() {
		fmt.Println("Insert TransactionLog ", time.Now().Format("2006-01-02 15:03:04"))
	}()

	_, err = global.Postgresql.Exec(
		`INSERT INTO TransactionLog 
		(OperatorID, Token, Content, IP, Location, Origin, UserAgent) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7);`,
		input.OperatorID,
		input.Content,
		input.Token,
		input.IP,
		input.Location,
		input.Origin,
		input.UserAgent,
	)

	return err
}
