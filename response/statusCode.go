package response

type StatusCode int

const (
	StatusSuccess StatusCode = 20000

	StatusParameterInvalid    StatusCode = 40000
	StatusNotLogin            StatusCode = 40100
	StatusFrequently          StatusCode = 40200
	StatusForbidden           StatusCode = 40300
	StatusNotFound            StatusCode = 40400
	StatusTokenInvalid        StatusCode = 40500
	StatusOtherClientLoggedIn StatusCode = 40600
	StatusSystemInternal      StatusCode = 50000
	StatusSystemTimeout       StatusCode = 50100
	StatusEtcdConnFailed      StatusCode = 52379
	StatusMysqlConnFailed     StatusCode = 53306
	StatusRedisConnFailed     StatusCode = 56379
)
