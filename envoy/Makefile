# envoy --version

validate:
	envoy --mode validate -c envoy-demo.yaml --log-path logs/custom.log

run: validate
	envoy -c envoy-demo.yaml --log-path logs/custom.log  --component-log-level upstream:debug,connection:trace,client:info
	# envoy -c envoy-demo.yaml --log-path logs/custom.log  --log-level info
