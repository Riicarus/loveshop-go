package e

type RespMsg string

const (
	// data bind or validate failed
	VALIDATE_FAILED_MSG RespMsg = "VALIDATE_FAILED"
	
	UNATHENRIZED_MSG RespMsg = "UNATHENRIZED"

	// rpc invoke failed, used when receiving rpc err
	RPC_FAILED_MSG RespMsg = "RPC_FAILED"

	INTERNAL_ERROR_MSG RespMsg = "INTERNAL_ERROR"
)