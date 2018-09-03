M = $(shell printf "\033[34;1m▶\033[0m")

standard-schema: $(info $(M) Generating standard schema)
	go-bindata -ignore=parser\.go -pkg=schema_standard -o=schema_standard/assets.go schema/scalar/... schema/type/... schema/standard/...

admin-schema: $(info $(M) Generating admin schema)
	go-bindata -ignore=parser.go -pkg=schema_admin -o=schema_admin/assets.go schema/scalar/... schema/type/... schema/admin/...

schema: standard-schema admin-schema

run: schema ; $(info $(M) Running the application...)
	APP_HOST=localhost APP_PORT=8080 DB_DSN=nathanb:password@tcp\(192.168.88.100:3306\)/nathanb?parseTime=true go run main.go
