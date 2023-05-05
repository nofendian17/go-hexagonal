package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_Exists(t *testing.T) {
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
		want    bool
		wantErr bool
	}{
		{
			name: "success - exist cache",
			fields: fields{
				client: db,
				ctx:    context.TODO(),
			},
			args: args{
				key: "key",
				mockExpect: func() {
					mock.ExpectExists("key").SetVal(1)
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "fail - exist cache error",
			fields: fields{
				client: db,
				ctx:    context.TODO(),
			},
			args: args{
				key: "key",
				mockExpect: func() {
					mock.ExpectExists("key").SetErr(errors.New("error"))
				},
			},
			want:    false,
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

			got, err := r.Exists(tt.args.key)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "Exist(%v)", tt.args.key)
		})
	}
}
