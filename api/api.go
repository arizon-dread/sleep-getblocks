package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"gitlab.com/arizon-dread/sleep-getblocks/models"
)

var delay int
var text string

func SetDelay(i int) {
	delay = i
}
func SetText(str string) {
	text = str
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy"))
}

func Sleep(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not read body"))
		return
	}

	jsonData := &models.Sleep{}
	if err = json.Unmarshal(body, jsonData); err == nil {
		if jsonData.Seconds >= 0 {
			log.Println("got a positive int")
			delay = jsonData.Seconds
			w.Write([]byte(fmt.Sprintf("setting delay to %vs", delay)))
		} else {
			log.Println("got a negative int")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("expected a positive int"))
			return
		}
	} else {
		log.Printf("got error, %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to marshal json into go struct"))
		return
	}
}
func GetBlocks(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	xmlReq := &models.Envelope{}
	if err := xml.Unmarshal(body, xmlReq); err == nil {
		log.Println("got xml request")
		//sleep
		time.Sleep(time.Duration(delay) * time.Second)
		//keep going
		if len(text) > 0 {
			w.Write([]byte(text))
		} else {
			//Set nextCreatedOnOrAfter timestamp + latestCancellation timestamp for each request
			datetime := time.Now()
			tz, _ := time.LoadLocation("Europe/Stockholm")
			currentTime := datetime.In(tz).Format("2006-01-02T03:04:05-07:00")
			latestCancellation := datetime.Add(-48 * time.Hour).In(tz).Format("2006-01-02T03:04:05-07:00")

			ssn := xmlReq.Body.GetBlocks.PatientId.SSNExtension
			careProviderId := xmlReq.Body.GetBlocks.CareProviderIds
			w.Write([]byte(`<?xml version=1.0 encoding=UTF-8?>
	<S:Envelope xmlns:S=http://schemas.xmlsoap.org/soap/envelope/>
	<S:Body>
		<ns3:GetBlocksResponse xmlns:ns2=urn:riv:informationsecurity:authorization:blocking:4 xmlns:ns3=urn:riv:informationsecurity:authorization:blocking:GetBlocksResponder:4 xmlns:ns4=urn:riv:itintegration:registry:1>
		<ns3:blockHeader>
			<ns2:result>
				<ns2:resultCode>OK</ns2:resultCode>
			</ns2:result>
			<ns2:blocks>
				<ns2:blockId>d34bb78a-7d3c-11ed-a0a1-44af280ae852</ns2:blockId>
				<ns2:blockType>Outer</ns2:blockType>
				<ns2:informationCareProviderId>` + careProviderId + `</ns2:informationCareProviderId>
				<ns2:patientId>
					<ns2:root>1.2.752.129.2.1.3.1</ns2:root>
					<ns2:extension>` + ssn + `</ns2:extension>
				</ns2:patientId>
				<ns2:ownerId>SERIALNUMBER=SE5565594230-BLM, CN=ws.sparradmin.inera.se, O=Inera AB, L=Stockholm, C=SE</ns2:ownerId>
			</ns2:blocks>
		<ns2:nextCreatedOnOrAfter>` + currentTime + `</ns2:nextCreatedOnOrAfter>
		<ns2:latestCancellation>` + latestCancellation + `</ns2:latestCancellation>
		</ns3:blockHeader>
	</ns3:GetBlocksResponse>
	</S:Body>
</S:Envelope>`))
		}
	} else {
		log.Printf("Failed to bind xml, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to marshal body into xml"))
	}
}
