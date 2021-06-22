package sdk_go

import (
	"github.com/obada-foundation/sdk-go/properties"
	"log"
)

type Status string

const StatusFunctional Status = "FUNCTIONAL"
const StatusNonFunctional Status = "NON_FUNCTIONAL"
const StatusDisposed Status = "DISPOSED"
const StatusStolen Status = "STOLEN"
const DisabledByOwner Status = "DISABLED_BY_OWNER"

type ObitIdDto struct {
	SerialNumberHash string `validate:"required"`
	Manufacturer     string `validate:"required"`
	PartNumber       string `validate:"required"`
}

type ObitDto struct {
	ObitIdDto
	OwnerDid       string `validate:"required"`
	ObdDid         string
	Matadata       map[string]string
	StructuredData map[string]string
	Documents      map[string]string
	ModifiedOn     int64 `validate:"min-modified-on"`
	Status         string
}

type Obit struct {
	obitId           properties.ObitId
	serialNumberHash properties.StringProperty
	manufacturer     properties.StringProperty
	partNumber       properties.StringProperty
	ownerDid         properties.StringProperty
	obdDid           properties.StringProperty
	metadata         properties.KvProperty
	structuredData   properties.KvProperty
	documents        properties.KvProperty
	modifiedOn       properties.IntProperty
	status           properties.StatusProperty
	debug            bool
	logger           *log.Logger
}
