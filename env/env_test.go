package env_test

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/nathanburkett/nathanb-api/env"
	"github.com/nathanburkett/nathanb-api/mock"
	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	type args struct {
		key            string
		panicAssertion func(assert.TestingT, assert.PanicTestFunc, ...interface{}) bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Resolves GOPATH",
			args: args{
				key:            "GOPATH",
				panicAssertion: assert.NotPanics,
			},
		},
		{
			name: "Panics on EXAMPLE_TEST",
			args: args{
				key:            "EXAMPLE_TEST",
				panicAssertion: assert.Panics,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.panicAssertion(t, func() {
				env.Must(tt.args.key)
			})
		})
	}
}

func TestNewReader(t *testing.T) {
	type args struct {
		ioReader io.Reader
	}
	tests := []struct {
		name string
		args args
		want env.Reader
	}{
		{
			name: "NewReader ",
			args: args{
				ioReader: mock.Reader{},
			},
			want: env.NewReader(mock.Reader{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := env.NewReader(tt.args.ioReader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_Read(t *testing.T) {
	type fields struct {
		reader io.Reader
	}
	type args struct {
		envPreAffection func()
		panicAssertion func(assert.TestingT, assert.PanicTestFunc, ...interface{}) bool
		assertionFunc func(assert.TestingT)
	}

	pwd, _ := os.Getwd()

	fileA, _ := os.Open(fmt.Sprintf("%s/../.env.test", pwd))
	defer fileA.Close()

	fileB, _ := os.Open(fmt.Sprintf("%s/../.env.test", pwd))
	defer fileB.Close()

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Reads successfully",
			fields: fields{
				reader: fileA,
			},
			args: args{
				envPreAffection: func() {},
				panicAssertion: assert.NotPanics,
				assertionFunc: func(t assert.TestingT) {
					_, exists := os.LookupEnv("DB_DSN")
					assert.True(t, exists)
				},
			},
		},
		{
			name: "Reads successfully. Continues on DB_DSN",
			fields: fields{
				reader: fileB,
			},
			args: args{
				envPreAffection: func() {
					os.Setenv("DB_DSN", "foo")
				},
				panicAssertion: assert.NotPanics,
				assertionFunc: func(t assert.TestingT) {
					_, exists := os.LookupEnv("DB_DSN")
					assert.True(t, exists)
				},
			},
		},
		{
			name: "Panics",
			fields: fields{
				reader: mock.Reader{
					ReadErr: errors.New("fail"),
				},
			},
			args: args{
				envPreAffection: func() {},
				panicAssertion: assert.Panics,
				assertionFunc: func(t assert.TestingT) {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.panicAssertion(t, func() {
				tt.args.envPreAffection()
				r := env.NewReader(tt.fields.reader)
				r.Read()
				tt.args.assertionFunc(t)
			})
		})
	}
}
