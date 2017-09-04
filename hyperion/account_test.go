package hyperion

import (
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/xorm"
	"github.com/yl10/mdd"
)

var testServer *mdd.Server

func init() {
	db, _ := xorm.NewEngine("mssql", "odbc:server=localhost\\SQLExpress;user id=sa;password=123;database=mdd;connection timeout=30")
	testServer = mdd.NewServer(db, nil)
}
func TestSync(t *testing.T) {

	testServer.Sync()
}

func TestRegDim(t *testing.T) {
	if err := testServer.RegisterDimension(DimAccount); err != nil {
		t.Error(err)
	}
}
