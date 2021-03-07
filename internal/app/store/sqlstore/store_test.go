package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("TEST_DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=advertisement sslmode=disable user=postgres password=yunis port=5432"
	}
	os.Exit(m.Run())
}
