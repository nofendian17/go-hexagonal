package redis

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Close(t *testing.T) {
	db, _ := redismock.NewClientMock()

	type fields struct {
		client *redis.Client
		ctx    context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success - close client",
			fields: fields{
				client: db,
				ctx:    context.TODO(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				client: tt.fields.client,
				ctx:    tt.fields.ctx,
			}

			err := r.Close()

			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
