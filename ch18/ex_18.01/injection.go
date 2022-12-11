package sql_injection

import (
	"database/sql"
	"fmt"
)

type UserDetails struct {
	Id         string
	CardNumber string
}

func GetCardNumber(db *sql.DB, userID string) (resp string, err error) {
	// WARN: don't use user provided data, directly and without sanitizaion. Use
	// prepared statements instead
	query := `SELECT CARD_NUMBER FROM USER_DETAILS WHERE USER_ID = ` + userID
	row := db.QueryRow(query)
	switch err = row.Scan(&resp); err {
	case sql.ErrNoRows:
		return resp, fmt.Errorf("no rows returned")
	case nil:
		return resp, err
	default:
		return resp, err
	}
}

func GetCardNumberSecure(db *sql.DB, userID string) (resp string, err error) {
	query := `SELECT CARD_NUMBER FROM USER_DETAILS WHERE USER_ID = ?`
	row := db.QueryRow(query, userID)
	switch err = row.Scan(&resp); err {
	case sql.ErrNoRows:
		return resp, fmt.Errorf("no rows returned")
	case nil:
		return resp, err
	default:
		return resp, err
	}
}
