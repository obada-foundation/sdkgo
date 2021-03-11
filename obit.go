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

func NewObit(serialNumberHash string, manufacturer string, partNumber string) (Obit, error) {
	var o Obit

	snh, err := properties.NewSerialNumberHash(serialNumberHash)
	m, err := properties.NewManufacturer(manufacturer)
	pn, err := properties.NewPartNumber(partNumber)
	id, err := properties.NewObitId(snh, m, pn)

	if err != nil {
		return o, err
	}

	o.obitId = id
	o.serialNumberHash = snh
	o.manufacturer = m
	o.partNumber = pn

	return o, nil
}

func (o *Obit) GetSerialNumberHash() properties.SerialNumberHash  {
	return o.serialNumberHash
}
