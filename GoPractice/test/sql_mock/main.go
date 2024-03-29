package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// @NOTE: the real connection is not required for tests
	db, err := sql.Open("mysql", "root@/blog")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = recordStats(db, 1 /*some user id*/, 5 /*some product id*/); err != nil {
		panic(err)
	}
}

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
		return
	}
	return
}

func get(db *sql.DB, key string) (users []*UserInfo, err error) {
	sqlStr := fmt.Sprintf(`SELECT id,account_type FROM user WHERE id IN(%s)`, key)
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := &UserInfo{}
		err = rows.Scan(&user.ID, &user.AccountType)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return
}

type UserInfo struct {
	ID          int64 `json:"id"`
	AccountType int64 `json:"account_type"`
}
