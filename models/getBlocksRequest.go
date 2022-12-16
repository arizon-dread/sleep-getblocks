package models

type Envelope struct {
	Header Header `xml:"s:Header"`
	Body   Body   `xml:"s:Body"`
}
type GetBlocksRequest struct {
	Envelope Envelope `xml:"s:Envelope"`
}
type Header struct {
	LogicalAddress string `xml:"h:LogicalAddress"`
}
type Body struct {
	GetBlocks GetBlocks `xml:"GetBlocks"`
}
type GetBlocks struct {
	PatientId       PatientId `xml:"patientId"`
	CareProviderIds string    `xml:"careProviderIds"`
}

type GetBlocksResponse struct {
	ResultCode string   `xml:"resultCode"`
	Blocks     []Blocks `xml:"blocks"`
}
type Blocks struct {
	BlockId                   string `xml:"blockId"`
	BlockType                 string `xml:"blockType"`
	InformationCareUnitId     string `xml:"informationCareUnitId"`
	InformationCareProviderId string `xml:"informationCareProviderId"`
	OwnerId                   string `xml:"ownerId"`
}
type PatientId struct {
	SSNRoot      string `xml:"root"`
	SSNExtension string `xml:"extension"`
}
