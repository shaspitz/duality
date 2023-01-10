module github.com/NicholasDotSol/duality

go 1.16

require (
	github.com/cosmos/admin-module v0.0.0
	github.com/cosmos/cosmos-sdk v0.45.7-0.20221104161803-456ca5663c5e
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/cosmos/interchain-security v0.0.0-20221104204906-6c0d718d8c59
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.14.0 // indirect
	github.com/ignite-hq/cli v0.22.0
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.5.0
	github.com/strangelove-ventures/packet-forward-middleware/v3 v3.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.0
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.7
	google.golang.org/genproto v0.0.0-20221114212237-e4508ebdbee1
	google.golang.org/grpc v1.50.1
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/cosmos/admin-module => github.com/jtieri/admin-module v0.0.0-20221116191954-1d63d5fc9608
	github.com/cosmos/ibc-go/v3 => github.com/jtieri/ibc-go/v3 v3.0.0-beta1.0.20221116191630-01c53c7f66f3
	github.com/cosmos/interchain-security v0.0.0-20221102103028-d7f8d448be65 => github.com/jtieri/interchain-security v0.0.0-20221116194529-59bf07eb134f
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	github.com/strangelove-ventures/packet-forward-middleware/v3 => github.com/strangelove-ventures/packet-forward-middleware/v3 v3.0.1-0.20230109224123-feac15ea2cb3
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
