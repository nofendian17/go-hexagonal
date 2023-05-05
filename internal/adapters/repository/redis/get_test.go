package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_Get(t *testing.T) {
	db, mock := redismock.NewClientMock()

	type fields struct {
		client *redis.Client
		ctx    context.Context
	}
	type args struct {
		key        string
		mockExpect func()
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success - get cache",
			fields: fields{
				client: db,
				ctx:    context.TODO(),
			},
			args: args{
				key: "key",
				mockExpect: func() {
					mock.ExpectGet("key").SetVal("value")
				},
			},
			want:    "value",
			wantErr: false,
		},
		{
			name: "fail - cache error",
			fields: fields{
				client: db,
				ctx:    context.TODO(),
			},
			args: args{
				key: "key",
				mockExpect: func() {
					mock.ExpectGet("key").SetErr(errors.New("error"))
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				client: tt.fields.client,
				ctx:    tt.fields.ctx,
			}

			tt.args.mockExpect()

			got, err := r.Get(tt.args.key)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "Get(%v)", tt.args.key)
		})
	}
}
