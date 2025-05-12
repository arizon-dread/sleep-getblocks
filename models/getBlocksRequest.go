package models

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  Header   `xml:"Header"`
	Body    Body     `xml:"Body"`
}
type GetBlocksRequest struct {
	Envelope Envelope `xml:"Envelope"`
}
type Header struct {
	XMLName        xml.Name `xml:"Header"`
	LogicalAddress string   `xml:"h:LogicalAddress"`
}
type Body struct {
	XMLName   xml.Name  `xml:"Body"`
	GetBlocks GetBlocks `xml:"GetBlocks"`
}
type GetBlocks struct {
	XMLName         xml.Name  `xml:"GetBlocks"`
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
	XMLName      xml.Name `xml:"patientId"`
	SSNRoot      string   `xml:"root"`
	SSNExtension string   `xml:"extension"`
}
