package sqldriver

import "database/sql/driver"

type Rows struct {
}

func (r *Rows) Columns() []string {
}

func (r *Rows) Close() error {
}

func (r *Rows) Next(dest []driver.Value) error {
}
