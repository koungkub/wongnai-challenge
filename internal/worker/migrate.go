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

	sqlCreateTableReview = `
	CREATE TABLE IF NOT EXISTS review (
		review_id INT(11),
		comment TEXT NOT NULL,
		created_at DATETIME DEFAULT current_timestamp(),
		updated_at DATETIME DEFAULT current_timestamp() ON UPDATE current_timestamp(),
		PRIMARY KEY (review_id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`
	sqlCreateTableFoodDic = `
	CREATE TABLE IF NOT EXISTS food_dictionary (
		food_id INT(11) AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		created_at DATETIME DEFAULT current_timestamp(),
		updated_at DATETIME DEFAULT current_timestamp() ON UPDATE current_timestamp(),
		PRIMARY KEY (food_id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`
)

// OpenFile open raw file
func OpenFile(name string) (*os.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, errors.Wrap(err, "can not open file")
	}

	return f, nil
}

// MigrateSchema create table schema in database
func MigrateSchema(db *sql.DB) error {
	_, err := db.ExecContext(context.TODO(), sqlCreateTableReview)
	if err != nil {
		return errors.Wrap(err, "migrate schema review")
	}

	_, err = db.ExecContext(context.TODO(), sqlCreateTableFoodDic)
	if err != nil {
		return errors.Wrap(err, "migrate schema food_dic")
	}

	return nil
}

// MigrateReview migrate schema table `review`
func MigrateReview(db *sql.DB, f io.Reader) error {
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
func MigrateFoodDic(db *sql.DB, f io.Reader) error {
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
