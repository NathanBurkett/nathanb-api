M = $(shell printf "\033[34;1m▶\033[0m")

standard-schema: $(info $(M) Generating standard schema)
	go-bindata -ignore=parser.go -pkg=schema_standard -o=schema_standard/assets.go schema/scalar/... schema/type/... schema/standard/...

admin-schema: $(info $(M) Generating admin schema)
	go-bindata -ignore=parser.go -pkg=schema_admin -o=schema_admin/assets.go schema/scalar/... schema/type/... schema/admin/...

schema: standard-schema admin-schema

html-coverage: $(info $(M) Running tests w/ coverage)
	./test-coverage/run.sh --html

coverage: $(info $(M) Running tests w/ coverage)
	./test-coverage/run.sh
