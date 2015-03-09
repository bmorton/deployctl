package client

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bmorton/deployster/server"
)

type Client struct {
	BaseURL    string
	Username   string
	Password   string
	connection *http.Client
}

type Unit server.VersionedUnit
type UnitsResponse struct {
	Units []*Unit `json:"units"`
}

func New(baseURL string, username string, password string, certPath string) *Client {
	transport := &http.Transport{}

	if certPath != "" {
		certPool := x509.NewCertPool()

		customCert, err := ioutil.ReadFile(certPath)
		if err != nil {
			log.Fatal("Could not load CA certificate!")
		}

		done := certPool.AppendCertsFromPEM(customCert)
		if !done {
			log.Fatal("Could not append CA certificate!")
		}

		transport.TLSClientConfig = &tls.Config{RootCAs: certPool, InsecureSkipVerify: false}
	}

	connection := &http.Client{
		Transport: transport,
	}

	return &Client{BaseURL: baseURL, Username: username, Password: password, connection: connection}
}

func (c *Client) CreateDeploy(service string, version string, destroyPrevious bool, instanceCount int) (*http.Response, error) {
	payload := server.DeployRequest{
		Deploy: server.Deploy{
			Version:         version,
			DestroyPrevious: destroyPrevious,
			InstanceCount:   instanceCount,
		},
	}

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.Encode(payload)

	url := fmt.Sprintf("%s/v1/services/%s/deploys", c.BaseURL, service)
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return &http.Response{}, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Content-Type", "application/json")

	return c.connection.Do(req)
}

func (c *Client) CreateTask(service string, version string, command string) (*http.Response, error) {
	payload := server.TaskRequest{
		Task: server.Task{
			Version: version,
			Command: command,
		},
	}

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.Encode(payload)

	url := fmt.Sprintf("%s/v1/services/%s/tasks", c.BaseURL, service)
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return &http.Response{}, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Content-Type", "application/json")

	return c.connection.Do(req)
}

func (c *Client) DestroyDeploy(service string, version string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v1/services/%s/deploys/%s", c.BaseURL, service, version)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return &http.Response{}, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Content-Type", "application/json")

	return c.connection.Do(req)
}

func (c *Client) GetUnits(service string) ([]*Unit, error) {
	url := fmt.Sprintf("%s/v1/services/%s/units", c.BaseURL, service)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []*Unit{}, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.connection.Do(req)
	if err != nil {
		return []*Unit{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var parsedResponse UnitsResponse
	err = decoder.Decode(&parsedResponse)
	if err != nil {
		return []*Unit{}, err
	}

	return parsedResponse.Units, nil
}
