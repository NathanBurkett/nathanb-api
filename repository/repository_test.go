package repository_test

import (
	"database/sql"
	"errors"
	"github.com/nathanburkett/nathanb-api/mock"
	"github.com/nathanburkett/nathanb-api/repository"
	"reflect"
	"testing"

	"github.com/nathanburkett/nathanb-api/criteria"
	"github.com/nathanburkett/nathanb-api/data"
)

func TestNew(t *testing.T) {
	type args struct {
		db data.Database
	}
	tests := []struct {
		name string
		args args
		want repository.Repository
	}{
		{
			name: "Valid Repository",
			args: args{
				db: &mock.Database{},
			},
			want: repository.Repository{
				DB: &mock.Database{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_First(t *testing.T) {
	type fields struct {
		DB data.Database
	}
	type args struct {
		dest interface{}
		cri  criteria.AbstractCriteria
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Returns nil",
			fields: fields{
				DB: &mock.Database{},
			},
			args: args{
				dest: "",
				cri: &mock.Criteria{},
			},
			wantErr: false,
		},
		{
			name: "Returns criteria.AbstractCriteria.ToSql error",
			fields: fields{
				DB: &mock.Database{},
			},
			args: args{
				dest: "",
				cri: &mock.Criteria{
					ToSqlErr: errors.New("foo"),
				},
			},
			wantErr: true,
		},
		{
			name: "Returns data.Database.Get error",
			fields: fields{
				DB: &mock.Database{
					GetErr: errors.New("foo"),
				},
			},
			args: args{
				dest: "",
				cri: &mock.Criteria{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.Repository{
				DB: tt.fields.DB,
			}
			if err := r.First(tt.args.dest, tt.args.cri); (err != nil) != tt.wantErr {
				t.Errorf("Repository.First() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_Find(t *testing.T) {
	type fields struct {
		DB data.Database
	}
	type args struct {
		dest interface{}
		cri  criteria.AbstractCriteria
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Returns nil",
			fields: fields{
				DB: &mock.Database{},
			},
			args: args{
				dest: "",
				cri: &mock.Criteria{},
			},
			wantErr: false,
		},
		{
			name: "Returns criteria.AbstractCriteria.ToSql error",
			fields: fields{
				DB: &mock.Database{},
			},
			args: args{
				dest: "",
				cri: &mock.Criteria{
					ToSqlErr: errors.New("foo"),
				},
			},
			wantErr: true,
		},
		{
			name: "Returns data.Database.Get error",
			fields: fields{
				DB: &mock.Database{
					GetErr: errors.New("foo"),
				},
			},
			args: args{
				dest: "",
				cri: &mock.Criteria{},
			},
			wantErr: true,
		},
		{
			name: "Doesn't return error on sql.ErrNoRows",
			fields: fields{
				DB: &mock.Database{
					GetErr: sql.ErrNoRows,
				},
			},
			args: args{
				dest: "",
				cri: &mock.Criteria{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.Repository{
				DB: tt.fields.DB,
			}
			if err := r.Find(tt.args.dest, tt.args.cri); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Find() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_All(t *testing.T) {
	type fields struct {
		DB data.Database
	}
	type args struct {
		dest interface{}
		cri  criteria.AbstractCriteria
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Returns nil",
			fields: fields{
				DB: &mock.Database{},
			},
			args: args{
				dest: "",
				cri: &mock.Criteria{},
			},
			wantErr: false,
		},
		{
			name: "Returns criteria.AbstractCriteria.ToSql error",
			fields: fields{
				DB: &mock.Database{},
			},
			args: args{
				dest: "",
				cri: &mock.Criteria{
					ToSqlErr: errors.New("foo"),
				},
			},
			wantErr: true,
		},
		{
			name: "Returns data.Database.Select error",
			fields: fields{
				DB: &mock.Database{
					SelectErr: errors.New("foo"),
				},
			},
			args: args{
				dest: "",
				cri: &mock.Criteria{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.Repository{
				DB: tt.fields.DB,
			}
			if err := r.All(tt.args.dest, tt.args.cri); (err != nil) != tt.wantErr {
				t.Errorf("Repository.All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
