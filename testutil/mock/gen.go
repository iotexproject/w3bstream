package mock

//go:generate mockgen -write_generate_directive -package=mock -destination=./mock_output.go -source=../../output/output.go
//go:generate mockgen -write_generate_directive -package=mock -destination=./mock_libp2p_host.go github.com/libp2p/go-libp2p/core/host Host
