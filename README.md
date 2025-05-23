# Testapplikation för att simulera delay mot Säkerhetstjänster spärr.

* Exempel på att starta container:
```
docker run -p 8080:8080 sleep-endpoint:v0.2
```

Exempel på att curla den:
```
curl http://localhost:8080/getblocks -d '<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
	<s:Header>
		<h:LogicalAddress xmlns:h="urn:riv:itintegration:registry:1" xmlns="urn:riv:itintegration:registry:1">SE165565594230-1000</h:LogicalAddress>
	</s:Header>
	<s:Body xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
		<GetBlocks xmlns="urn:riv:informationsecurity:authorization:blocking:GetBlocksResponder:4">
			<patientId>
				<root xmlns="urn:riv:informationsecurity:authorization:blocking:4">1.2.752.129.2.1.3.1</root>
				<extension xmlns="urn:riv:informationsecurity:authorization:blocking:4">191212121212</extension>
			</patientId>
			<careProviderIds>SE2321000164-7381037590003</careProviderIds>
		</GetBlocks>
	</s:Body>
</s:Envelope>'
```
För att ändra hur lång timeout den ska ha kan man posta ett request till den:  
`curl http://localhost:8080/sleep -d '{"seconds": 120}'`  

Applikationen kan behöva en reverse-proxy framför som validerar mTLS-cert från klient-applikationen och sedan skickar trafiken bakåt.