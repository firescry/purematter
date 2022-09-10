package client

import (
	"reflect"
	"testing"
)

func TestUpdateEndpointHost(t *testing.T) {
	host := "255.255.255.255"
	endpoint := "http://0.0.0.0/api/v1/test"
	res, err := UpdateEndpointHost(host, endpoint)
	if err != nil {
		t.Errorf("UpdateEndpointHost returned an error: %s\n", err)
	}
	exp := "http://255.255.255.255/api/v1/test"
	if exp != res {
		t.Errorf("Expected '%s', got '%s'\n", exp, res)
	}
}

func TestUpdateEndpointHostWithBrokenURL(t *testing.T) {
	host := "255.255.255.255"
	endpoint := "::http://0.0.0.0/api/v1/test"
	_, err := UpdateEndpointHost(host, endpoint)
	if err == nil {
		t.Errorf("UpdateEndpointHost did not return an error for broken URL: %s\n", endpoint)
	}
}

func TestGenerateEndpoints(t *testing.T) {
	host := "255.255.255.255"
	res, err := GenerateEndpoints(host)
	if err != nil {
		t.Errorf("GenerateEndpoints returned an error: %s\n", err)
	}
	exp := map[string]string{
		"security": "http://255.255.255.255/di/v1/products/0/security",
	}
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("Expected %v, got %v\n", exp, res)
	}
}
