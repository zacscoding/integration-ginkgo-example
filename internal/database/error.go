package database

import "github.com/go-sql-driver/mysql"

// IsDuplicateEntryError returns a true if duplicate entry error, otherwise false
func IsDuplicateEntryError(err error) bool {
	if err == nil {
		return false
	}
	switch err.(type) {
	case *mysql.MySQLError:
		e := err.(*mysql.MySQLError)
		if e.Number == 1062 {
			return true
		}
	}
	return false
}
