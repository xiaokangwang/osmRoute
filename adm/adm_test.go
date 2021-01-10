package adm

import (
	"os"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	os.MkdirAll("/run/shm/testdb", 0700)
	CreateIndex("/run/shm/ireland.osm.pbf", "/run/shm/testdb")
}
