package sdk_go

import (
	"fmt"
	"github.com/obada-protocol/sdk-go/properties"
	"log"
)

type Sdk struct {
	logger *log.Logger
	debug bool
}

func NewSdk(log *log.Logger, debug bool) (*Sdk, error)  {
	return &Sdk{
		logger: log,
		debug: debug,
	}, nil
}

func (sdk *Sdk) NewObit(serialNumberHash string, manufacturer string, partNumber string) (Obit, error) {
	var o Obit

	sdk.Debug(fmt.Sprintf("NewObit(%q, %q, %q)", serialNumberHash, manufacturer, partNumber))

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

func (sdk *Sdk) NewObitId(serialNumberHash string, manufacturer string, partNumber string) (properties.ObitId, error) {
	var obitId properties.ObitId

	sdk.Debug(fmt.Sprintf("NewObit(%q, %q, %q)", serialNumberHash, manufacturer, partNumber))

	snh, err := properties.NewSerialNumberHash(serialNumberHash)

	if err != nil {
		return obitId, err
	}

	snhh := snh.GetHash()

	sdk.Debug(fmt.Sprintf("serialNumberHash = %q :: hash = %q :: decHash = %q", snh.GetValue(), snhh.GetHash(), snhh.GetDec()))

	if err != nil {
		return obitId, err
	}

	m, err := properties.NewManufacturer(manufacturer)

	if err != nil {
		return obitId, err
	}

	pn, err := properties.NewPartNumber(partNumber)

	if err != nil {
		return obitId, err
	}

	obitId, err = properties.NewObitId(snh, m, pn)

	if err != nil {
		return obitId, err
	}

	return obitId, nil
}

func (sdk *Sdk) Debug(message string) {
	if sdk.debug {
		if sdk.logger != nil {
			sdk.logger.Println(message)
		} else {
			log.Println(message)
		}
	}
}


