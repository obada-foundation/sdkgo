package sdk_go

import (
	"github.com/obada-protocol/sdk-go/properties"
)

type Obit struct {
	obitId           properties.ObitId
	serialNumberHash properties.SerialNumberHash
	manufacturer     properties.Manufacturer
	partNumber       properties.PartNumber
	ownerDid         string
	obdDid           string
	metadata         string
	structuredData   string
	documents        string
	modifiedAt       string
	status           string
}

func (o *Obit) GetSerialNumberHash() properties.SerialNumberHash  {
	return o.serialNumberHash
}
