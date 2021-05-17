package sdk_go

import "github.com/obada-protocol/sdk-go/properties"

type ObitDto struct {
	serialNumberHash string `validate:"required"`
	manufacturer     string `validate:"required"`
	partNumber       string `validate:"required"`
	ownerDid         string `validate:"required"`
	obdDid           string `validate:"required"`
}

type Obit struct {
	obitId properties.ObitId
}
