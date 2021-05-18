package sdk_go

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/obada-protocol/sdk-go/properties"
)

var validate *validator.Validate

type Sdk struct {
	logger   *log.Logger
	debug    bool
	validate *validator.Validate
}

func NewSdk(log *log.Logger, debug bool) (*Sdk, error) {
	return &Sdk{
		logger:   log,
		debug:    debug,
		validate: validator.New(),
	}, nil
}

func (sdk *Sdk) NewObit(dto ObitDto) (Obit, error) {
	var o Obit

	err := sdk.validate.Struct(dto)
	if err != nil {
		return o, err
	}

	serialNumberProp, err := properties.NewStringProperty(dto.GetSerialNumberHash())
	manufacturerProp, err := properties.NewStringProperty(dto.GetManufacturer())
	pnProp, err := properties.NewStringProperty(dto.GetPartNumber())

	if err != nil {
		return o, err
	}

	obitId, err := properties.NewObitId(serialNumberProp, manufacturerProp, pnProp)

	if err != nil {
		return o, err
	}

	o.obitId = obitId

	return o, nil
}

func (sdk *Sdk) NewObitId(dto ObitIdDto) (properties.ObitId, error) {
	var obitId properties.ObitId

	sdk.Debug(fmt.Sprintf("NewObitId(%q, %q, %q)", dto.GetSerialNumberHash(), dto.GetManufacturer(), dto.GetPartNumber()))

	serialNumberProp, err := properties.NewStringProperty(dto.GetSerialNumberHash())
	manufacturerProp, err := properties.NewStringProperty(dto.GetManufacturer())
	pnProp, err := properties.NewStringProperty(dto.GetPartNumber())

	if err != nil {
		return obitId, err
	}

	obitId, err = properties.NewObitId(serialNumberProp, manufacturerProp, pnProp)

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
