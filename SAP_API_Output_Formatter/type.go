package sap_api_output_formatter

type BatchMasterRecord struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	Product           string `json:"Product"`
	APISchema         string `json:"api_schema"`
	BatchMasterRecord string `json:"batch_master_record_code"`
	Deleted           string `json:"deleted"`
}

type Batch struct {
	Material                  string `json:"Material"`
	BatchIdentifyingPlant     string `json:"BatchIdentifyingPlant"`
	Batch                     string `json:"Batch"`
	Supplier                  string `json:"Supplier"`
	BatchBySupplier           string `json:"BatchBySupplier"`
	CountryOfOrigin           string `json:"CountryOfOrigin"`
	RegionOfOrigin            string `json:"RegionOfOrigin"`
	MatlBatchAvailabilityDate string `json:"MatlBatchAvailabilityDate"`
	ShelfLifeExpirationDate   string `json:"ShelfLifeExpirationDate"`
	ManufactureDate           string `json:"ManufactureDate"`
	CreationDateTime          string `json:"CreationDateTime"`
	LastChangeDateTime        string `json:"LastChangeDateTime"`
	BatchIsMarkedForDeletion  bool   `json:"BatchIsMarkedForDeletion"`
}