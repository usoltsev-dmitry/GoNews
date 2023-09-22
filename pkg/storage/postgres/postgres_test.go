package postgres

import (
	"GoNews/pkg/storage"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestStorage_AddAuthor(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		t storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.AddAuthor(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.AddAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.AddAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}
