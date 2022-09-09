package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-batch-master-record-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToBatch(raw []byte, l *logger.Logger) (*Batch, error) {
	pm := &responses.Batch{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Batch. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	batch := &Batch{
		Material:                  data.Material,
		BatchIdentifyingPlant:     data.BatchIdentifyingPlant,
		Batch:                     data.Batch,
		Supplier:                  data.Supplier,
		BatchBySupplier:           data.BatchBySupplier,
		CountryOfOrigin:           data.CountryOfOrigin,
		RegionOfOrigin:            data.RegionOfOrigin,
		MatlBatchAvailabilityDate: data.MatlBatchAvailabilityDate,
		ShelfLifeExpirationDate:   data.ShelfLifeExpirationDate,
		ManufactureDate:           data.ManufactureDate,
		CreationDateTime:          data.CreationDateTime,
		LastChangeDateTime:        data.LastChangeDateTime,
		BatchIsMarkedForDeletion:  data.BatchIsMarkedForDeletion,
	}

	return batch, nil
}