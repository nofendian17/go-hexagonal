package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_Ping(t *testing.T) {
	db, mock := redismock.NewClientMock()
	type fields struct {
		client *redis.Client
		ctx    context.Context
	}
	tests := []struct {
		name       string
		fields     fields
		mockExpect func()
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "success - ping",
			fields: fields{
				client: db,
				ctx:    context.TODO(),
			},
			mockExpect: func() {
				mock.ExpectPing().SetVal("pong")
			},
			wantErr: assert.NoError,
		},
		{
			name: "fail - ping",
			fields: fields{
				client: db,
				ctx:    context.TODO(),
			},
			mockExpect: func() {
				mock.ExpectPing().SetErr(errors.New("error"))
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				client: tt.fields.client,
				ctx:    tt.fields.ctx,
			}
			tt.mockExpect()
			tt.wantErr(t, r.Ping(), fmt.Sprintf("Ping()"))
		})
	}
}
