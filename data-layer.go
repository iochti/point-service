package main

import (
	"fmt"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/namsral/flag"
)

// DataLayerInterface abstracts the db connection
type DataLayerInterface interface {
	CreatePoint(pt *client.Point) error
	QueryDB(cmd string) (res []client.Result, err error)
}

var (
	DBName *string
)

// InfluxDL implements the DataLayerInterface
type InfluxDL struct {
	Client      client.Client
	BatchConfig client.BatchPointsConfig
}

//Init the database
func (i *InfluxDL) Init() error {
	url := flag.String("influx-url", "localhost", "Influx host url")
	DBName = flag.String("influx-db-name", "iochti", "Influx default database name")
	username := flag.String("influx-user", "iochti", "Influx username")
	inflxPwd := flag.String("influx-pwd", "", "Influx user's password")
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s:8086", *url),
		Username: *username,
		Password: *inflxPwd,
	})

	if err != nil {
		return err
	}
	q := client.Query{
		Command: fmt.Sprintf(`
            CREATE DATABASE IF NOT EXISTS %s;
            CREATE RETENTION POLICY \"ten_min_only\" ON \"%s\" DURATION 10min REPLICATION 1;
        `, *DBName, *DBName),
		Database: *DBName,
	}
	if response, err := c.Query(q); err != nil {
		if response.Error() != nil {
			return response.Error()
		}
	}
	i.Client = c
	i.BatchConfig = client.BatchPointsConfig{
		Database: *DBName,
	}
	return nil
}

// CreatePoint creates a point in the DB
func (i *InfluxDL) CreatePoint(pt *client.Point) error {
	bp, err := client.NewBatchPoints(i.BatchConfig)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)
	if err = i.Client.Write(bp); err != nil {
		return err
	}
	return nil
}

// QueryDB convenience function to query the database
func (i *InfluxDL) QueryDB(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: *DBName,
	}

	if response, err := i.Client.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
