package e

type RespCode int32

const (
	// all right
	NONE_CODE RespCode = 0
	// data bind or validate failed
	VALIDATE_FAILED_CODE RespCode = 1

	UNATHENRIZED_CODE RespCode = 2

	// rpc invoke failed, used when receiving rpc err
	RPC_FAILED_CODE RespCode = 3

	INTERNAL_ERROR_CODE RespCode = 4
)
