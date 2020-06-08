package worker

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/csv"
	"io"
	"os"

	"github.com/pkg/errors"
)

const (
	sqlMigrateReviewTable  = `INSERT INTO review (review_id, comment) VALUES (?, ?)`
	sqlMigrateFoodDicTable = `INSERT INTO food_dictionary (name) VALUES (?)`
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
	stmt, err := db.PrepareContext(context.TODO(), sqlMigrateReviewTable)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	csv := csv.NewReader(f)
	csv.Comma = ';'
	csv.Read()

	for {
		r, err := csv.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "read file")
		}

		if _, err := stmt.ExecContext(context.TODO(), r[0], r[1]); err != nil {
			return errors.Wrapf(err, "exec statement at id [%v]", r[0])
		}
	}
}

// MigrateFoodDic migrate schema table `food_dictionary`
func MigrateFoodDic(db *sql.DB, f *os.File) error {
	stmt, err := db.PrepareContext(context.TODO(), sqlMigrateFoodDicTable)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	scanner := bufio.NewScanner(f)
	for i := 0; i < 20000; i++ {
		if scanner.Scan() {
			if _, err := stmt.ExecContext(context.TODO(), scanner.Text()); err != nil {
				return errors.Wrapf(err, "exec statement at name [%v]", scanner.Text())
			}
		}
	}

	return nil
}
