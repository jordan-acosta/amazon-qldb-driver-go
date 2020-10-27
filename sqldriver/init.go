package sqldriver

import "database/sql"

func init() {
	sql.Register("qldb", &Driver{})
}
