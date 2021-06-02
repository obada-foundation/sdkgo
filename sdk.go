package sdk_go

import (
	"fmt"
	"github.com/obada-protocol/sdk-go/hash"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/obada-protocol/sdk-go/properties"
)

type Sdk struct {
	logger   *log.Logger
	debug    bool
	validate *validator.Validate
}

func NewSdk(log *log.Logger, debug bool) (*Sdk, error) {
	return &Sdk{
		logger:   log,
		debug:    debug,
		validate: initializeValidator(),
	}, nil
}

func initializeValidator() *validator.Validate {
	v := validator.New()

	return v
}

// NewObit s
func (sdk *Sdk) NewObit(dto ObitDto) (Obit, error) {
	var o Obit

	err := sdk.validate.Struct(dto)
	if err != nil {
		return o, err
	}

	sdk.Debug(fmt.Sprintf("NewObit(%v)", dto))

	serialNumberProp, err := properties.NewStringProperty(dto.SerialNumberHash)

	if err != nil {
		return o, err
	}

	sdk.Debug(fmt.Sprintf("Hash(%q) => %q", serialNumberProp.GetValue(), serialNumberProp.GetHash().GetHash()))

	manufacturerProp, err := properties.NewStringProperty(dto.Manufacturer)

	if err != nil {
		return o, err
	}

	sdk.Debug(fmt.Sprintf("Hash(%q) => %q", manufacturerProp.GetValue(), manufacturerProp.GetHash().GetHash()))

	pnProp, err := properties.NewStringProperty(dto.PartNumber)

	if err != nil {
		return o, err
	}

	sdk.Debug(fmt.Sprintf("Hash(%q) => %q", pnProp.GetValue(), pnProp.GetHash().GetHash()))

	obitId, err := properties.NewObitId(serialNumberProp, manufacturerProp, pnProp)

	if err != nil {
		return o, err
	}

	o.obitId = obitId
	o.serialNumberHash = serialNumberProp
	o.manufacturer = manufacturerProp
	o.partNumber = pnProp

	return o, nil
}

func (o Obit) GetRootHash() (hash.Hash, error) {
	var rootHash hash.Hash

	sum := o.obitId.GetHash().GetDec() +
		o.serialNumberHash.GetHash().GetDec() +
		o.manufacturer.GetHash().GetDec() +
		o.partNumber.GetHash().GetDec() +
		o.ownerDid.GetHash().GetDec() +
		o.obdDid.GetHash().GetDec() +
		o.metadata.GetHash().GetDec() +
		o.structureData.GetHash().GetDec() +
		o.documents.GetHash().GetDec() +
		o.modifiedAt.GetHash().GetDec() +
		o.status.GetHash().GetDec()

	rootHash, err := hash.NewHash(fmt.Sprintf("%x", sum))

	if err != nil {
		return rootHash, err
	}

	return rootHash, nil
}

// NewObitId c
func (sdk *Sdk) NewObitId(dto ObitIdDto) (properties.ObitId, error) {
	var obitId properties.ObitId

	if sdk.debug {
		sdk.Debug(fmt.Sprintf("NewObitId(%q, %q, %q)", dto.SerialNumberHash, dto.Manufacturer, dto.PartNumber))
	}

	err := sdk.validate.Struct(dto)
	if err != nil {
		return obitId, err
	}

	snProp, err := properties.NewStringProperty(dto.SerialNumberHash)

	if err != nil {
		return obitId, err
	}

	snPropHash := snProp.GetHash()

	if sdk.debug {
		sdk.Debug(fmt.Sprintf("Hash(%q) -> Hash2Dec(%q) -> %d", snProp.GetValue(), snPropHash.GetHash(), snPropHash.GetDec()))
	}

	mnProp, err := properties.NewStringProperty(dto.Manufacturer)

	if err != nil {
		return obitId, err
	}

	mnPropHash := mnProp.GetHash()

	if sdk.debug {
		sdk.Debug(fmt.Sprintf("Hash(%q) -> Hash2Dec(%q) -> %d", mnProp.GetValue(), mnPropHash.GetHash(), mnPropHash.GetDec()))
	}

	pnProp, err := properties.NewStringProperty(dto.PartNumber)

	if err != nil {
		return obitId, err
	}

	pnPropHash := pnProp.GetHash()

	if sdk.debug {
		sdk.Debug(fmt.Sprintf("Hash(%q) -> Hash2Dec(%q) -> %d", pnProp.GetValue(), pnPropHash.GetHash(), pnPropHash.GetDec()))
	}

	obitId, err = properties.NewObitId(snProp, mnProp, pnProp)

	if err != nil {
		return obitId, err
	}

	if sdk.debug {
		sum := snPropHash.GetDec() + mnPropHash.GetDec() + pnPropHash.GetDec()

		sdk.Debug(
			fmt.Sprintf("(%d + %d + %d) -> %d",
				snPropHash.GetDec(),
				mnPropHash.GetDec(),
				pnPropHash.GetDec(),
				sum,
			))

		dec2Hex := fmt.Sprintf("%x", sum)
		dec2HexHash, _ := hash.NewHash(dec2Hex)

		sdk.Debug(fmt.Sprintf("Dec2Hex(%d) -> %s -> Hash(%s) -> %s", sum, dec2Hex, dec2Hex,dec2HexHash.GetHash()))

		sdk.Debug(fmt.Sprintf("Hash : %s", obitId.GetHash().GetHash()))
		sdk.Debug(fmt.Sprintf("Did : %s", obitId.GetDid()))
		sdk.Debug(fmt.Sprintf("Usn : Base58(%s) -> %s", obitId.GetHash().GetHash(), obitId.GetUsn()))
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
