# envoy --version

validate:
	envoy --mode validate -c configs/static.yaml --log-path logs/custom.log

run-static: validate
	envoy -c configs/static.yaml --log-path logs/custom.log  --component-log-level upstream:debug,connection:trace,client:info

run-dynamic:
	envoy -c configs/dynamic-filesystem.yaml --log-path logs/custom.log  --log-level info
