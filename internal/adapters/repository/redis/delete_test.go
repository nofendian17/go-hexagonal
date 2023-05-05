package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_Delete(t *testing.T) {
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
		want    int64
		wantErr bool
	}{
		{
			name: "success - delete cache",
			fields: fields{
				client: db,
				ctx:    context.TODO(),
			},
			args: args{
				key: "key",
				mockExpect: func() {
					mock.ExpectDel("key").SetVal(int64(1))
				},
			},
			want:    1,
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
					mock.ExpectDel("key").SetErr(errors.New("error"))
				},
			},
			want:    0,
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

			err := r.Delete(tt.args.key)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
