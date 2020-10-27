package sqldriver

type Result struct {
	qldbRes interface{}
}

// TODO: Can we get this info from the result QLDB returns?

func (r *Result) LastInsertId() (int64, error) {
	return 0, nil
}

func (r *Result) RowsAffected() (int64, error) {
	return 0, nil
}
