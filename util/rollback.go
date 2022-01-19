package util

import "gorm.io/gorm"

func CommitOrRollback(db *gorm.DB) {
	err := recover()
	if err != nil {
		db.Rollback()
	} else {
		db.Commit()
	}
}
