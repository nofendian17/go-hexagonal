package open_search

import (
	"crypto/tls"
	"fmt"
	"github.com/opensearch-project/opensearch-go"
	"net/http"
	"user-svc/internal/shared/config"
)

const (
	protocolHttps = "https://"
	protocolHttp  = "http://"
)

type OpenSearchClient struct {
	Client *opensearch.Client
}

func NewClient(config *config.Config) *OpenSearchClient {
	protocol := protocolHttp
	if config.Log.OpenSearch.HttpSecure {
		protocol = protocolHttps
	}

	addr := fmt.Sprintf("%s%s:%d", protocol, config.Log.OpenSearch.Host, config.Log.OpenSearch.Port)
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{
			addr,
		},
		Username: config.Log.OpenSearch.Username,
		Password: config.Log.OpenSearch.Password,
	})

	if config.App.Debug {
		fmt.Println(fmt.Sprintf("Trying connect openSearch with %s", addr))
	}

	if err != nil {
		fmt.Println("failed to open connection to openSearch", err.Error())
	}

	_, err = client.Ping()
	if err != nil {
		fmt.Println("failed to ping connection to openSearch", err.Error())
	}

	return &OpenSearchClient{
		Client: client,
	}
}
