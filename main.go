package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/arizon-dread/sleep-getblocks/models"
)

var text string
var delay int

func main() {

	content, err := ioutil.ReadFile("response.xml")
	strDelay := os.Getenv("DELAY")

	intDelay, err := strconv.Atoi(strDelay)
	if err == nil {
		delay = intDelay
	} else {
		fmt.Printf("error reading env var DELAY as int, using 60s as delay. Error was %v", err)
		delay = 60
	}

	if err != nil {
		fmt.Printf("Could not read file b/c: %v\nWill use static response.\n", err)
	} else {
		fmt.Printf("Will use response.xml as response for every request.")
	}
	text = string(content)

	router := gin.Default()
	router.GET("/healthz", healthz)
	router.POST("/sleep", sleep)
	router.POST("/GetBlocks", getBlocks)
	router.Run(":8080")
}

func healthz(c *gin.Context) {
	c.Request.Response = &http.Response{Status: "OK - Healthy", StatusCode: 200}
}

func sleep(c *gin.Context) {
	jsonData := models.Sleep{}

	if err := c.BindJSON(&jsonData); err == nil {
		if jsonData.Seconds > 0 {
			fmt.Println("got a positive int")
			time.Sleep(time.Duration(jsonData.Seconds) * time.Second)
		} else {
			fmt.Println("got a negative int")
			c.AbortWithError(400, errors.New("Bad request"))
		}
	} else {
		fmt.Printf("got error, %v", err)
		c.AbortWithError(500, err)
	}
}
func getBlocks(c *gin.Context) {
	xmlReq := models.GetBlocksRequest{}
	if err := c.BindXML(&xmlReq); err == nil {
		fmt.Println("got xml request")
		//sleep
		time.Sleep(time.Duration(delay) * time.Second)
		//keep going
		if len(text) > 0 {
			c.String(200, text)
		} else {
			ssn := xmlReq.Envelope.Body.GetBlocks.PatientId.SSNExtension
			careProviderId := xmlReq.Envelope.Body.GetBlocks.CareProviderIds
			c.String(200, `<?xml version=1.0 encoding=UTF-8?>
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
										<ns2:informationCareProviderId>`+careProviderId+`</ns2:informationCareProviderId>
										<ns2:patientId>
											<ns2:root>1.2.752.129.2.1.3.1</ns2:root>
											<ns2:extension>`+ssn+`</ns2:extension>
										</ns2:patientId>
										<ns2:ownerId>SERIALNUMBER=SE5565594230-BLM, CN=ws.sparradmin.inera.se, O=Inera AB, L=Stockholm, C=SE</ns2:ownerId>
									</ns2:blocks>
								<ns2:nextCreatedOnOrAfter>2022-12-16T10:35:52.659+01:00</ns2:nextCreatedOnOrAfter>
								<ns2:latestCancellation>2022-12-14T15:18:51.000+01:00</ns2:latestCancellation>
								</ns3:blockHeader>
							</ns3:GetBlocksResponse>
							</S:Body>
						</S:Envelope>`)
		}
	} else {
		fmt.Printf("Failed to bind xml, %v", err)
		c.AbortWithError(400, err)
	}
}
