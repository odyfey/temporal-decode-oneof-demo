.PHONY: generate
generate:
	rm -rf gen
	buf generate --path proto/yetanothercloud

.PHONY: worker
worker: generate
	go run app/cmd/worker/main.go

.PHONY: starter
starter:
	go run app/cmd/starter/main.go