run:
	go run cmd/scrobble-status/main.go

ssl-proxy:
	npx local-ssl-proxy --source 7188 --target 8564