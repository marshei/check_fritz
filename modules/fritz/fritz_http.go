package fritz

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// DoSoapRequest does two request to authenticate and handle the SOAP request
func DoSoapRequest(soapRequest *SoapData, resps chan<- []byte, errs chan<- error, debug bool) {

	// enable this for debug sessions
	if debug {
		fmt.Printf("---\nSOAP Request:\n")
		fmt.Printf("-> URL    : %v\n", string(soapRequest.URL))
		fmt.Printf("-> Service: %v\n", string(soapRequest.Service))
		fmt.Printf("-> Action : %v\n---\n", string(soapRequest.Action))
	}

	soapClient := createNewSoapClient()

	// prepare first request
	req, err := newSoapRequest(soapRequest, debug)

	if err != nil {
		errs <- err
		return
	}

	// the first request is always unauthenticated due to how digest authentication works
	resp, err := soapClient.Do(req)

	if err != nil {
		errs <- err
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		errs <- err
		return
	}
	resp.Body.Close()

	// enable this for debug sessions
	if debug {
		fmt.Printf("---\nFirst SOAP Response:\n---\n%v\n---\n", string(body))
	}

	// create immediately a new request with authentication
	req, err = newSoapRequest(soapRequest, debug)

	if err != nil {
		errs <- err
		return
	}

	if resp.StatusCode == http.StatusUnauthorized {
		authHeader, err := doDigestAuthentication(resp, soapRequest)

		if err != nil {
			errs <- err
			return
		} else if authHeader == "" {
			errs <- errors.New("Returned authentication header is empty")
			return
		}
		req.Header.Set("Authorization", authHeader)

	} else if resp.StatusCode == http.StatusOK {
		resps <- body
		return
	} else {
		errs <- fmt.Errorf("Unexpected response status code: %d", resp.StatusCode)
	}

	// second request is now authenticated
	resp, err = soapClient.Do(req)

	if err != nil {
		errs <- err
		return
	}

	if resp.StatusCode == http.StatusUnauthorized {
		errs <- errors.New("Unauthorized: wrong username or password")
		return
	} else if resp.StatusCode != http.StatusOK {
		errs <- fmt.Errorf("Unexpected response status code: %d", resp.StatusCode)
		return
	}

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		errs <- err
		return
	}
	resp.Body.Close()

	// enable this for debug sessions

	if debug {
		fmt.Printf("---\nSecond SOAP Response:\n---\n%v\n---\n", string(body))
	}

	resps <- body
}

func createNewSoapClient() *http.Client {
	ht := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true},
	}

	return &http.Client{Transport: ht}
}

func newSoapRequest(soapRequest *SoapData, debug bool) (*http.Request, error) {
	requestBody := newSoapRequestBody(soapRequest, debug)
	req, err := http.NewRequest("POST", soapRequest.URL, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/xml; charset='utf-8'")
	if (strings.HasPrefix(soapRequest.Service, "urn")) {
	  req.Header.Set("SoapAction", soapRequest.Service+"#"+soapRequest.Action)
	} else {
	  req.Header.Set("SoapAction", "urn:dslforum-org:service:"+soapRequest.Service+":1#"+soapRequest.Action)
	}

	return req, nil
}

func newSoapRequestBody(soapRequest *SoapData, debug bool) []byte {
	var request bytes.Buffer

	request.WriteString("<?xml version='1.0?>")
	request.WriteString("<s:Envelope xmlns:s='http://schemas.xmlsoap.org/soap/envelope/' s:encodingStyle='http://schemas.xmlsoap.org/soap/encoding/'>")
	request.WriteString("<s:Body>")
	if (strings.HasPrefix(soapRequest.Service, "urn")) {
	  request.WriteString("<u:" + soapRequest.Action + " xmlns:u='" + soapRequest.Service + "'>")
	} else {
	  request.WriteString("<u:" + soapRequest.Action + " xmlns:u='urn:dslforum-org:service:" + soapRequest.Service + ":1'>")
	}

	if &soapRequest.XMLVariable != nil {
		request.WriteString("<" + soapRequest.XMLVariable.Name + ">")
		request.WriteString(soapRequest.XMLVariable.Value)
		request.WriteString("</" + soapRequest.XMLVariable.Name + ">")
	}

	request.WriteString("</u:" + soapRequest.Action + ">")
	request.WriteString("</s:Body>")
	request.WriteString("</s:Envelope>")

	if debug {
    fmt.Printf("---\nSOAP Request Body:\n---\n%v\n---\n", request.String())
	}
	return request.Bytes()
}
