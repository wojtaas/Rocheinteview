# Rocheinteview

A simple API and GRPC service, that's provides PING service.
## Ports:
API : `40000` - local env `HTTP_PORT`

GRPC: `30000` - local env `GRPC_PORT`

Can be changed by using local env's.

## Run & Build & test

To run type in terminal: 
`make run`

To build type in terminal:
`make build`

To run tests type in terminal:
`make test`

## Generate new swagger:
To generate swagger be sure that You have installed 'swag' - more information: https://github.com/swaggo/swag.

To run generator type in terminal: `make generate-docs`.

## Generate mocks:

Before use be sure that You have installed `mockery`, then in terminal type: `make generate_mocks`.