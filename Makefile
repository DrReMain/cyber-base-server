init_api:
	hz new -module github.com/DrReMain/cyber-base-server \
		 --model_dir biz/hertz_gen -idl idl/template/greet.thrift --unset_omitempty true

update_api:
	hz update --model_dir biz/hertz_gen -idl idl/template/greet.thrift --unset_omitempty true

gorm_migrate:
	go run -tags migrate generate.go
