package criteria

import (
	"fmt"
	"github.com/nathanburkett/nathanb-api/data_object"
)

type ModelInterpreter interface {
	handleArgs(AbstractCriteria, interface{})
	handleField(string) (string, bool, error)
	checkDefaultPaginationArgs(PaginationArgs) PaginationArgs
}

type ModelInterpreterFactory struct {}

func (ModelInterpreterFactory) Create(model data_object.Model) (ModelInterpreter, error) {
	var (
		interpreter ModelInterpreter
		err         error
	)

	switch T := model.(type) {
	case data_object.Category:
		interpreter = categoryInterpretation{}
	case data_object.Classification:
		interpreter = classificationInterpretation{}
	case data_object.ContentBlock:
		interpreter = contentBlockInterpretation{}
	case data_object.Media:
		interpreter = mediaInterpretation{}
	case data_object.Profile:
		interpreter = profileInterpretation{}
	case data_object.Publication:
		interpreter = publicationInterpretation{}
	case data_object.User:
		interpreter = userInterpretation{}
	default:
		err = fmt.Errorf("unknown data object type: %s", T)
	}

	return interpreter, err
}
