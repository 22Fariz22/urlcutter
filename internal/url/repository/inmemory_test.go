package repository

import (
	"context"
	"sync"
	"testing"
)

func Test_memoryStorage_Save(t *testing.T) {
	type fields struct {
		storage map[string]string
		mutex   sync.RWMutex
	}
	newMem := NewMemory()

	type args struct {
		ctx   context.Context
		long  string
		short string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok save",
			fields: fields{
				storage: newMem.storage,
				mutex:   newMem.mutex,
			},
			args: args{
				ctx:   context.Background(),
				long:  "mock_long",
				short: "",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memoryStorage{
				storage: tt.fields.storage,
			}
			got, err := m.Save(tt.args.ctx, tt.args.long, tt.args.short)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryStorage_Get(t *testing.T) {
	type fields struct {
		storage map[string]string
		mutex   sync.RWMutex
	}

	newMem := NewMemory()

	type args struct {
		ctx   context.Context
		short string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok get",
			fields: fields{
				storage: newMem.storage,
				mutex:   newMem.mutex,
			},
			args: args{
				ctx:   context.Background(),
				short: "mock_short",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memoryStorage{
				storage: tt.fields.storage,
			}
			got, err := m.Get(tt.args.ctx, tt.args.short)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
