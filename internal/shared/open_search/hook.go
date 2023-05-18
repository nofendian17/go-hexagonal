package open_search

import (
	"context"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"github.com/sirupsen/logrus"
	"strings"
)

type OpenSearchHook struct {
	Client    *opensearch.Client
	IndexName string
}

func (h *OpenSearchHook) Fire(entry *logrus.Entry) error {
	log, err := entry.String()
	if err != nil {
		return err
	}

	go func() {
		document := opensearchapi.IndexRequest{
			Index: h.IndexName,
			Body:  strings.NewReader(log),
		}

		_, err := document.Do(context.Background(), h.Client)
		if err != nil {
			logrus.Errorf("failed to push log to OpenSearch: %v", err)
		}
	}()

	return nil
}

func (h *OpenSearchHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func NewOpenSearchHook(client *opensearch.Client, indexName string) *OpenSearchHook {
	return &OpenSearchHook{
		Client:    client,
		IndexName: indexName,
	}
}
