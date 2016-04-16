package jobs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/k4orta/muni"
	"github.com/k4orta/tunnel-watch/storage"
)

func TestFetchTransit(t *testing.T) {
	fakeServer := makeFakeServer()
	muni.SetConfig(muni.TransitConfig{DefaultURL: fakeServer.URL + "/"})

	storage.RunStorageTest(t, func(db *sqlx.DB, t *testing.T) {
		err := fetchTransit([]string{"N"})
		if err != nil {
			t.Error(err)
		}
		v, _ := storage.GetVehiclesAfterTime(db, time.Unix(1460498787925/1000, 0).Add(time.Minute*-5))
		if len(v) == 0 {
			t.Error("Did not add any vehicles to the DB")
		}

		if len(v) == 37 {
			t.Error("Did not add all of the vehicles to the DB")
		}
	})
}

func makeFakeServer() *httptest.Server {
	fp := filepath.Join("test-fixtures", "vehicles.xml")
	data, _ := ioutil.ReadFile(fp)
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/xml")
		fmt.Fprint(w, string(data))
	}))
}
