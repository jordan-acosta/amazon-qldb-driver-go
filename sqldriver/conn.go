package sqldriver

import (
	"context"
	"database/sql/driver"
	"fmt"

	"github.com/amzn/ion-go/ion"
	"github.com/awslabs/amazon-qldb-driver-go/qldbdriver"
)

type Conn struct {
	qldbDriver *qldbdriver.QLDBDriver
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

func (c *Conn) Begin() (driver.Tx, error) {
	return nil, nil
}

func (c *Conn) Close() error {
	c.qldbDriver.Shutdown(context.Background())
	return nil
}

func (c *Conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	qldbRes, err := c.qldbDriver.Execute(ctx, func(txn qldbdriver.Transaction) (interface{}, error) {
		_, err := txn.Execute(query)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	result := &Result{
		qldbRes: qldbRes,
	}
	return result, nil
}

func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	p, err := driver.Execute(ctx, func(txn qldbdriver.Transaction) (interface{}, error) {
		result, err := txn.Execute(query)
		if err != nil {
			return nil, err
		}

		// Assume the result is not empty
		hasNext := result.Next(txn)
		if !hasNext && result.Err() != nil {
			return nil, result.Err()
		}

		ionBinary := result.GetCurrentData()

		temp := new(Person)
		err = ion.Unmarshal(ionBinary, temp)
		if err != nil {
			return nil, err
		}

		return *temp, nil
	})
	if err != nil {
		panic(err)
	}

	var returnedPerson Person
	returnedPerson = p.(Person)

	if returnedPerson != person {
		fmt.Print("Queried result does not match inserted struct")
	}
}
