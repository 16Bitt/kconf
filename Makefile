kconf: dockercompose
	cd pkg/kconf && go build
	cd cmd/kconf && go build

dockercompose:
	cd internal/dockercompose && go build