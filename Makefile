BINARY_NAME=cobbler
EXECUTOR?=docker
COBBLER_SERVER_URL=http://localhost:8081/cobbler_api
TEST?=$$(go list ./... |grep -v 'vendor')
VERSION=0.0.1

build:
	@echo "building package"
	go build -o ${BINARY_NAME} main.go
	@echo "creating shell completions"
	make shell_completions

build-docker:
	@echo "building docker"
	${EXECUTOR} build -t cobbler/cli:latest -f packaging/docker/production/Dockerfile .

build-rpm-docker:
	@docker build -t localhost/cobbler-cli-pkg:opensuse-tumblewed -f packaging/docker/openSUSE_Tumbleweed/Dockerfile .
	@docker run --rm -v $(CURDIR)/rpms/openSUSE_Tumbleweed:/root/rpmbuild/RPMS -v $(CURDIR):/workspace localhost/cobbler-cli-pkg:opensuse-tumblewed

build-rpm:
	@cp packaging/rpm/cobbler-cli.spec /root/rpmbuild/SPECS/cobbler-cli.spec
	@cd ..; tar --exclude dist --exclude ".idea" --exclude ubuntu-20.04.1-legacy-server-amd64.iso --exclude extracted_iso_image --transform="s/workspace/cobbler-cli-${VERSION}/" -zcvf "cobbler-cli-${VERSION}.tar.gz" /workspace
	@mv ../cobbler-cli-${VERSION}.tar.gz /root/rpmbuild/SOURCES
	@go mod vendor; tar -zcvf "vendor.tar.gz" vendor; mv vendor.tar.gz /root/rpmbuild/SOURCES
	@rpmbuild --define "_topdir /root/rpmbuild" \
         --bb /root/rpmbuild/SPECS/cobbler-cli.spec

clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -rf config/completions/*

cleandoc: ## Cleans the docs directory.
	@echo "cleaning documentation"
	@cd docs; make clean > /dev/null 2>&1

doc:
	@echo "creating: documentation"
	@cd docs; make html

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

test:
	@./testing/start.sh ${COBBLER_SERVER_URL}
	go test -v -coverprofile="coverage.out" -covermode="atomic" $(TEST)

shell_completions:
	@mkdir -p config/completions/bash
	@mkdir -p config/completions/fish
	@mkdir -p config/completions/powershell
	@mkdir -p config/completions/zsh
	./${BINARY_NAME} completion bash > config/completions/bash/cobbler
	./${BINARY_NAME} completion fish > config/completions/fish/cobbler
	./${BINARY_NAME} completion powershell > config/completions/powershell/cobbler
	./${BINARY_NAME} completion zsh > config/completions/zsh/cobbler

.PHONY: build build-docker build-rpm-docker build-rpm clean cleandoc doc run test shell_completions
