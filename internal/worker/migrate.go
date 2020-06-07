package worker

import (
	"database/sql"
	"os"
)

// OpenFileFoodDic open raw file food_dictionary.txt
func OpenFileFoodDic(name string) (*os.File, error) {

}

// OpenFileReview open raw file test_file.csv
func OpenFileReview(name string) (*os.File, error) {

}

// MigrateReview migrate schema table `review`
func MigrateReview(db *sql.DB, f *os.File) error {

}

// MigrateFoodDic migrate schema table `food_dictionary`
func MigrateFoodDic(db *sql.DB, f *os.File) error {

}
