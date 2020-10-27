package sqldriver

import (
	"database/sql/driver"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"github.com/awslabs/amazon-qldb-driver-go/qldbdriver"
)

type Driver struct {
}

func (d *Driver) Open(name string) (driver.Conn, error) {
	// "name" is typically a DB connection string.
	// QLDB has an unusual connection string format, compared to "real" SQL DBs,
	region := "us-east-1"
	ledger := "quick-start"

	awsSession := session.Must(session.NewSession(aws.NewConfig().WithRegion(region)))
	qldbSession := qldbsession.New(awsSession)
	drvr, err := qldbdriver.New(
		ledger,
		qldbSession,
		// TODO: is it possible to support functional opts while adhering the driver interface?
		// func(options *qldbdriver.DriverOptions) {
		// 	options.LoggerVerbosity = qldbdriver.LogInfo
		// },
	)
	if err != nil {
		return err
	}

	return &Conn{
		qldbdriver: drvr,
	}
}
