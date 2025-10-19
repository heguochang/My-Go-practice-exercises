package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Migrator().DropTable(&Account{})
	db.Migrator().DropTable(&Transaction{})
	db.AutoMigrate(&Account{}, &Transaction{})
	fromAccount := &Account{
		Balance: 100,
	}
	toAccount := &Account{
		Balance: 100,
	}

	db.Create(fromAccount)
	db.Create(toAccount)

	tx := db.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}

	// 查询并检查转出账户
	var fa Account
	firstQueryResult := tx.First(&fa, fromAccount.ID)
	if firstQueryResult.Error != nil {
		tx.Rollback()
		fmt.Println("转出账户不存在 ", firstQueryResult.Error)
		return
	} else {
		if fa.Balance < 100 {
			tx.Rollback()
			fmt.Printf("转出账户余额不足,当前账户余额为 %d", fa.Balance)
			return
		}
	}

	// 更新转出账户
	updateAccountResult := tx.Model(&Account{}).Where("id = ?", fromAccount.ID).Update("balance", fa.Balance-100)
	if updateAccountResult.Error != nil {
		tx.Rollback()
		fmt.Println("扣减账户余额失败", updateAccountResult.Error)
		return
	}

	// 更新收款账户
	updateToAccountResult := tx.Model(&Account{}).Where("id = ?", toAccount.ID).Update("balance", gorm.Expr("balance + ?", 100))
	if updateToAccountResult.Error != nil {
		tx.Rollback()
		fmt.Println("更新收款帐户余额失败", updateToAccountResult.Error)
		return
	}

	// 创建交易记录
	t := &Transaction{
		FromAccountId: fromAccount.ID,
		ToAccountId:   toAccount.ID,
		Amount:        100,
	}

	createTsResult := tx.Create(t)
	if createTsResult.Error != nil {
		tx.Rollback()
		fmt.Println("记录交易记录失败", createTsResult.Error)
		return
	}

	if err = tx.Commit().Error; err != nil {
		fmt.Println("事务提交失败", err)
	}
}

type Account struct {
	gorm.Model
	Balance int
}

type Transaction struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        int
}
