package open_search

import (
	"context"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

// OpenSearchHook is a Logrus hook for sending logs to OpenSearch.
type OpenSearchHook struct {
	Client    *opensearch.Client
	IndexName string
}

// Fire sends the log entry to OpenSearch.
func (h *OpenSearchHook) Fire(entry *logrus.Entry) error {
	log, err := entry.String()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		document := opensearchapi.IndexRequest{
			Index: h.IndexName,
			Body:  strings.NewReader(log),
		}

		res, err := document.Do(context.Background(), h.Client)
		if err != nil {
			logrus.Errorf("failed to push log to OpenSearch: %v", err)
			return
		}

		err = res.Body.Close()
		if err != nil {
			logrus.Errorf("failed to close body to OpenSearch: %v", err)
			return
		}
	}()

	wg.Wait()
	return nil
}

// Levels returns all log levels for which this hook should be fired.
func (h *OpenSearchHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// NewOpenSearchHook creates a new OpenSearchHook.
func NewOpenSearchHook(client *opensearch.Client, indexName string) *OpenSearchHook {
	return &OpenSearchHook{
		Client:    client,
		IndexName: indexName,
	}
}
