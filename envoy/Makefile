# envoy --version

validate:
	envoy --mode validate -c static.yaml --log-path logs/custom.log

run: validate
	envoy -c static.yaml --log-path logs/custom.log  --component-log-level upstream:debug,connection:trace,client:info
	# envoy -c dynamic-filesystem.yaml --log-path logs/custom.log  --log-level info
