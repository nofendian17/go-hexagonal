package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"user-svc/internal/shared/config"
)

func TestRepository_Set(t *testing.T) {
	cfg := config.New()
	db, mock := redismock.NewClientMock()
	type fields struct {
		client *redis.Client
		ctx    context.Context
	}
	type args struct {
		key        string
		value      interface{}
		expiration time.Duration
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		mockExpect func()
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "success - set cache",
			fields: fields{
				client: db,
				ctx:    context.TODO(),
			},
			args: args{
				key:        "key",
				value:      "val",
				expiration: time.Duration(cfg.Database.Redis.Lifetime) * time.Second,
			},
			mockExpect: func() {
				mock.ExpectSet("key", "val", time.Duration(cfg.Database.Redis.Lifetime)*time.Second).SetVal("ok")
			},
			wantErr: assert.NoError,
		},
		{
			name: "fail - set cache error",
			fields: fields{
				client: db,
				ctx:    context.TODO(),
			},
			args: args{
				key:        "key",
				value:      "val",
				expiration: time.Duration(cfg.Database.Redis.Lifetime) * time.Second,
			},
			mockExpect: func() {
				mock.ExpectSet("key", "val", time.Duration(cfg.Database.Redis.Lifetime)*time.Second).SetErr(fmt.Errorf("set error"))
			},
			wantErr: assert.Error,
		},
		{
			name: "fail - invalid value type",
			fields: fields{
				client: db,
				ctx:    context.TODO(),
			},
			args: args{
				key:        "key",
				value:      struct{ Name string }{Name: "val"},
				expiration: time.Duration(cfg.Database.Redis.Lifetime) * time.Second,
			},
			mockExpect: func() {
				mock.ExpectSet("key", struct{ Name string }{Name: "val"}, time.Duration(cfg.Database.Redis.Lifetime)*time.Second).SetErr(nil)
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
			tt.wantErr(t, r.Set(tt.args.key, tt.args.value, tt.args.expiration), fmt.Sprintf("Set(%v, %v, %v)", tt.args.key, tt.args.value, tt.args.expiration))
		})
	}
}
