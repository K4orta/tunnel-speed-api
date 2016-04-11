package storage

import (
	"testing"
)

func TestCreateConnection(t *testing.T) {
	db, err := CreateConnection()
	if err != nil {
		t.Error(err)
	}
	db.Close()
}
