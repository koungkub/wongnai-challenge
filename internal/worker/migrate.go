package worker

import (
	"database/sql"
	"os"

	"github.com/pkg/errors"
)

// OpenFile open raw file
func OpenFile(name string) (*os.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, errors.Wrap(err, "can not open file")
	}

	return f, nil
}

// MigrateReview migrate schema table `review`
func MigrateReview(db *sql.DB, f *os.File) error {

}

// MigrateFoodDic migrate schema table `food_dictionary`
func MigrateFoodDic(db *sql.DB, f *os.File) error {

}
