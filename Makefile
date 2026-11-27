BINARY_NAME=cobbler
EXECUTOR?=docker
COBBLER_SERVER_URL=http://localhost:25151/
TEST?=$$(go list ./... |grep -v 'vendor')
# Package version for rpm/deb builds. Override via environment or `make VERSION=...`.
# Prereleases must use `~` instead of `-` (e.g. 1.0.0~rc2).
VERSION?=0.0.0

build:
	@echo "building package"
	go build -o ${BINARY_NAME} main.go
	@echo "creating shell completions"
	make shell_completions

build-docker:
	@echo "building docker"
	${EXECUTOR} build -t cobbler/cli:latest -f packaging/docker/production/Dockerfile .

build-rpm-docker: build-rpm-docker-tumbleweed build-rpm-docker-leap16

build-rpm-docker-tumbleweed:
	@docker build -t localhost/cobbler-cli-pkg:opensuse-tumbleweed -f packaging/docker/openSUSE_Tumbleweed/Dockerfile .
	@docker run --rm -e VERSION=$(VERSION) -v $(CURDIR)/rpms/openSUSE_Tumbleweed:/root/rpmbuild/RPMS:Z -v $(CURDIR):/workspace:Z localhost/cobbler-cli-pkg:opensuse-tumbleweed

build-rpm-docker-leap16:
	@docker build -t localhost/cobbler-cli-pkg:opensuse-leap16 -f packaging/docker/openSUSE_Leap16/Dockerfile .
	@docker run --rm -e VERSION=$(VERSION) -v $(CURDIR)/rpms/openSUSE_Leap16:/root/rpmbuild/RPMS:Z -v $(CURDIR):/workspace:Z localhost/cobbler-cli-pkg:opensuse-leap16

build-deb-docker:
	@docker build -t localhost/cobbler-cli-pkg:debian-13 -f packaging/docker/Debian_13/Dockerfile .
	@docker run --rm -e VERSION=$(VERSION) -v $(CURDIR)/debs/Debian_13:/usr/src/cobbler-cli/deb-build:Z -v $(CURDIR):/workspace:Z localhost/cobbler-cli-pkg:debian-13

build-rpm:
	@cp packaging/rpm/cobbler-cli.spec /root/rpmbuild/SPECS/cobbler-cli.spec
	@sed -i "s/^Version:.*/Version:        ${VERSION}/" /root/rpmbuild/SPECS/cobbler-cli.spec
	@cd ..; tar --exclude dist --exclude ".idea" --exclude .git --exclude rpms --exclude debs --exclude vendor --exclude ubuntu-20.04.1-legacy-server-amd64.iso --exclude extracted_iso_image --transform="s/workspace/cobbler-cli-${VERSION}/" -zcvf "cobbler-cli-${VERSION}.tar.gz" /workspace
	@mv ../cobbler-cli-${VERSION}.tar.gz /root/rpmbuild/SOURCES
	@go mod vendor; tar -zcvf "vendor.tar.gz" vendor; mv vendor.tar.gz /root/rpmbuild/SOURCES
	@rpmbuild --define "_topdir /root/rpmbuild" \
         --bb /root/rpmbuild/SPECS/cobbler-cli.spec

build-deb:
	@mkdir debian; cp -r packaging/deb/* debian/
	@cd debian; rename "s/debian\.//" *
	@sed -i -e "s/^Version:.*/Version: ${VERSION}/" -e "s/^DEBTRANSFORM-TAR:.*/DEBTRANSFORM-TAR: cobbler-cli-${VERSION}.tar.gz/" debian/cobbler-cli.dsc
	@sed -i -e "1s/(.*)/(${VERSION})/" -e "3s/.*/  * Package build of version ${VERSION}/" debian/changelog
	@go mod vendor; tar -zcvf "vendor.tar.gz" vendor; mv vendor.tar.gz debian/; rm -r vendor
	@cd debian; tar -xvzf vendor.tar.gz
	@debuild -us -uc
	@rm -r debian
	@cp ../cobbler-cli_* /usr/src/cobbler-cli/deb-build

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

.PHONY: build build-docker build-rpm-docker build-rpm-docker-tumbleweed build-rpm-docker-leap16 build-rpm build-deb-docker build-deb clean cleandoc doc run test shell_completions
