package data_object

const FieldId = "id"
const FieldCreatedAt = "created_at"
const FieldUpdatedAt = "updated_at"
const FieldDeletedAt = "deleted_at"

type Model interface {
	Table() string
	Fields() []string
}
