# Check lint, code styling rules. e.g. pylint, phpcs, eslint, style (java) etc ...
.PHONY: style
style:
	@echo "+ $@"
	@golangci-lint run \
		--tests=false \
		--enable-all \
		--disable gocyclo \
		--disable gosec \
		-v \
		./...

# Cyclomatic complexity check (McCabe), radon (python), eslint (js), PHPMD, rules (scala) etc ...
.PHONY: complexity
complexity:
	@echo "+ $@"
	@golangci-lint run \
		--tests=false \
		--disable-all \
		--enable gocyclo \
		-v \
		./...

# Launch static application security testing (SAST). e.g Gosec, bandit, Flawfinder, NodeJSScan, phpcs-security-audit, brakeman etc...
.PHONY: security-sast
security-sast:
	@echo "+ $@"
	@golangci-lint run \
		--tests=false \
		--disable-all \
		--enable gosec \
		-v \
		./...

# Format code. e.g Prettier (js), format (golang)
.PHONY: format
format:
	@echo "+ $@"
	@go fmt ./...

# Shortcut to launch all the test tasks (unit, functional and integration).
.PHONY: test
test: test-unit
	@echo "+ $@"

# Launch unit tests. e.g. pytest, jest (js), phpunit, JUnit (java) etc ...
.PHONY: test-unit
test-unit:
	@echo "+ $@"
	@go test \
		-race \
		-v \
		-cover \
		-coverprofile \
		coverage.out
