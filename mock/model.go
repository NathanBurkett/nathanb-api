package mock

type Model struct {}

func (Model) Table() string {
	return ""
}

func (Model) Fields() []string {
	return []string{}
}
