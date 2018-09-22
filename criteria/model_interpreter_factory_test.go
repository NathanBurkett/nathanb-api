package criteria_test

import (
	"github.com/nathanburkett/nathanb-api/criteria"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/nathanburkett/nathanb-api/data_object"
)

type mockModel struct {}

func (mockModel) Table() string {
	return ""
}

func TestModelInterpreterBridge_Determine(t *testing.T) {
	type args struct {
		model data_object.Model
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "data_object.Category yields criteria.categoryInterpretation",
			args: args{
				model: data_object.Category{},
			},
			wantErr: false,
		},
		{
			name: "data_object.Classification yields criteria.classificationInterpretation",
			args: args{
				model: data_object.Classification{},
			},
			wantErr: false,
		},
		{
			name: "data_object.ContentBlock yields criteria.contentBlockInterpretation",
			args: args{
				model: data_object.ContentBlock{},
			},
			wantErr: false,
		},
		{
			name: "data_object.Media yields criteria.mediaInterpretation",
			args: args{
				model: data_object.Media{},
			},
			wantErr: false,
		},
		{
			name: "data_object.Profile yields criteria.profileInterpretation",
			args: args{
				model: data_object.Profile{},
			},
			wantErr: false,
		},
		{
			name: "data_object.Publication yields criteria.publicationInterpretation",
			args: args{
				model: data_object.Publication{},
			},
			wantErr: false,
		},
		{
			name: "data_object.User yields criteria.userInterpretation",
			args: args{
				model: data_object.User{},
			},
			wantErr: false,
		},
		{
			name: "non model yields err",
			args: args{
				model: mockModel{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bridge := criteria.ModelInterpreterFactory{}
			got, err := bridge.Create(tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelInterpreterFactory.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Implements(t, (*criteria.ModelInterpreter)(nil), got)
			}
		})
	}
}
