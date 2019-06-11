kconf: dockercompose kubernetes
	cd pkg/kconf && go build
	cd cmd/kconf && go build

kubernetes:
	cd internal/kubernetes && go build

dockercompose:
	cd internal/dockercompose && go build
	
run: kconf
	cmd/kconf/kconf

install: kconf
	cd cmd/kconf && go install
	
