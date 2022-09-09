package sap_api_input_reader

import (
	"sap-api-integrations-batch-master-record-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToBatch() *requests.Batch {
	data := sdc.BatchMasterRecord
	return &requests.Batch{
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
}