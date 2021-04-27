module github.com/feitianlove/multiplePracticeTaking

go 1.15

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/go-redis/redis/v8 v8.4.4
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/kr/text v0.2.0 // indirect
	github.com/moby/moby v20.10.6+incompatible
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	golang.org/x/net v0.0.0-20210316092652-d523dce5a7f4
	google.golang.org/genproto v0.0.0-20210423144448-3a41ef94ed2b
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/olivere/elastic.v3 v3.0.75
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
