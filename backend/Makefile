.PHONY: run migration

run:
	go run cmd/server/main.go

migration:
	go run cmd/migration/main.go $(filter-out $@,$(MAKECMDGOALS))

	# Prevent Make from treating extra args as targets
%:
	@: