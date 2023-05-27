package database

import "database/sql"

func ChangeLastRequestByLink(url string, db *sql.DB) error {
	_, err := db.Exec("UPDATE scrapers SET last_request = CURRENT_TIMESTAMP WHERE link = ?", url)
	if err != nil {
		return err
	}
	return err
}
