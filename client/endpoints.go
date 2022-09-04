package client

import "net/url"

var DefaultEndpoints = map[string]string{
	"security": "http://0.0.0.0/di/v1/products/0/security",
}

func GenerateEndpoints(host string) (map[string]string, error) {
	endpoints := map[string]string{}
	for endpoint, url := range DefaultEndpoints {
		newUrl, err := UpdateEndpointHost(host, url)
		if err != nil {
			return nil, err
		}
		endpoints[endpoint] = newUrl
	}
	return endpoints, nil
}

func UpdateEndpointHost(host, endpoint string) (string, error) {
	e, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	e.Host = host
	return e.String(), nil
}
