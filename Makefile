include .env

migstatus:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DBSTRING) goose -dir=./migrations/ status

migup:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DBSTRING) goose -dir=./migrations/ up


migdown:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DBSTRING) goose -dir=./migrations/ down
