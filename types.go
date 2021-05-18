package sdk_go

import "github.com/obada-protocol/sdk-go/properties"

type ObitIdDto struct {
	SerialNumberHash string `validate:"required"`
	Manufacturer     string `validate:"required"`
	PartNumber       string `validate:"required"`
}

type ObitDto struct {
	ObitIdDto
	ownerDid string `validate:"required"`
	obdDid   string `validate:"required"`
}

type Obit struct {
	obitId properties.ObitId
	serialNumberHash properties.StringProperty
	manufacturer properties.StringProperty
	partNumber properties.StringProperty
}
