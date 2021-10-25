BINARY_NAME=cobbler

build:
	@echo "building package"
	go build -o ${BINARY_NAME} main.go
	@echo "creating shell completions"
	make shell_completions

clean:
	go clean
	rm ${BINARY_NAME}
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

shell_completions:
	@mkdir -p config/completions/bash
	@mkdir -p config/completions/fish
	@mkdir -p config/completions/powershell
	@mkdir -p config/completions/zsh
	./${BINARY_NAME} completion bash > config/completions/bash/cobbler
	./${BINARY_NAME} completion fish > config/completions/fish/cobbler
	./${BINARY_NAME} completion powershell > config/completions/powershell/cobbler
	./${BINARY_NAME} completion zsh > config/completions/zsh/cobbler
