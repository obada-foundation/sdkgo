package sdk_go

import "github.com/obada-protocol/sdk-go/properties"

type ObitIdDto struct {
	serialNumberHash string `validate:"required"`
	manufacturer     string `validate:"required"`
	partNumber       string `validate:"required"`
}

func (dto *ObitIdDto) GetSerialNumberHash() string {
	return dto.serialNumberHash
}

func (dto *ObitIdDto) GetManufacturer() string {
	return dto.manufacturer
}

func (dto *ObitIdDto) GetPartNumber() string {
	return dto.partNumber
}

type ObitDto struct {
	ObitIdDto
	ownerDid string `validate:"required"`
	obdDid   string `validate:"required"`
}

type Obit struct {
	obitId properties.ObitId
}
