package env_test

import (
	"github.com/nathanburkett/nathanb-api/env"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMust(t *testing.T) {
	type args struct {
		key string
		panicAssertion func(assert.TestingT, assert.PanicTestFunc, ...interface{}) bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Resolves GOPATH",
			args: args{
				key: "GOPATH",
				panicAssertion: assert.NotPanics,
			},
		},
		{
			name: "Panics on EXAMPLE_TEST",
			args: args{
				key: "EXAMPLE_TEST",
				panicAssertion: assert.Panics,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.panicAssertion(t, func () {
				env.Must(tt.args.key)
			})
		})
	}
}
