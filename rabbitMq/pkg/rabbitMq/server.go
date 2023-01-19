package rabbitMq

import (
	"github.com/houseofcat/turbocookedrabbit/v2/pkg/tcr"
	logger "github.com/jonny-mark/package-demo/rabbitMq/pkg/log"
)

var RabbitService *tcr.RabbitService

func NewRabbitService() {
	conf := GetConfig()
	//connectPool, err := tcr.NewConnectionPoolWithErrorHandler(conf.PoolConfig, func(err error) {
	//	print(err.Error())
	//})
	//if err != nil {
	//	print(err.Error())
	//}
	//rabbitService, err := tcr.NewRabbitServiceWithConnectionPool(connectPool, conf, "", "", nil, nil)
	//if err != nil {
	//	print(err.Error())
	//}
	rabbitService, err := tcr.NewRabbitService(conf, "", "", nil, nil)
	if err != nil {
		logger.Error(err)
	}
	RabbitService = rabbitService
	return
}
