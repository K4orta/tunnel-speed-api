package storage

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
)

// Schema is a type with create and drop commands for the db
type Schema struct {
	create string
	drop   string
}

var dbSchema = Schema{
	create: `
    create TABLE vehicles (
      route_tag text NOT NULL,
      vehicle_id text NOT NULL,
      time_received timestamp default now(),
      heading int,
      dir_tag text,
      lat double precision,
      lng double precision,
      leading_vehicle_id text,
      predictable boolean,
      secs_since_report int,
      speed_km_hr float
    );
  `,
	drop: `
    drop TABLE vehicles;
  `,
}

var pgdb *sqlx.DB

// SetupDBForTesting initializes the storage module for testing
func SetupDBForTesting() {
	connectionDSN = os.Getenv("MUNI_TEST_POSTGRES_DSN")
	db, err := CreateConnection()
	if err != nil {
		fmt.Printf("Error connecting to test DB:\n %v\n", err)
	}
	pgdb = db
}

// MultiExec is a helper function for running multiple queries.
func MultiExec(e sqlx.Execer, query string) {
	stmts := strings.Split(query, ";\n")
	if len(strings.Trim(stmts[len(stmts)-1], " \n\t\r")) == 0 {
		stmts = stmts[:len(stmts)-1]
	}
	for _, s := range stmts {
		_, err := e.Exec(s)
		if err != nil {
			fmt.Println(err, s)
		}
	}
}

func (s Schema) postgres() (string, string) {
	return s.create, s.drop
}

// InjectMockData inserts some sample entries for testing
func InjectMockData() {

}

// RunStorageTest is a utility for running a test with a temporary schema
func RunStorageTest(t *testing.T, test func(db *sqlx.DB, t *testing.T)) {
	runner := func(db *sqlx.DB, t *testing.T, create, drop string) {
		defer func() {
			MultiExec(db, drop)
		}()

		MultiExec(db, create)
		test(db, t)
	}

	create, drop := dbSchema.postgres()
	runner(pgdb, t, create, drop)
}
