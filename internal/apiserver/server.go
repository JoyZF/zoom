package apiserver

import (
	"context"
	"github.com/JoyZF/zoom/internal/apiserver/config"
	"github.com/JoyZF/zoom/internal/pkg/options"
	"github.com/JoyZF/zoom/internal/pkg/server"
	"github.com/marmotedu/iam/pkg/storage"
)

type apiServer struct {
	// TODO
	//gs               *shutdown.GracefulShutdown
	redisOptions *options.RedisOptions
	//gRPCAPIServer    *grpcAPIServer
	genericAPIServer *server.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	//gs := shutdown.New()
	//gs.AddShutdownManager(posixsignal.NewPosixSignalManager())
	//
	//genericConfig, err := buildGenericConfig(cfg)
	//if err != nil {
	//	return nil, err
	//}
	//
	//extraConfig, err := buildExtraConfig(cfg)
	//if err != nil {
	//	return nil, err
	//}
	//
	//genericServer, err := genericConfig.Complete().New()
	//if err != nil {
	//	return nil, err
	//}
	//extraServer, err := extraConfig.complete().New()
	//if err != nil {
	//	return nil, err
	//}
	//
	server := &apiServer{
		//gs:               gs,
		//redisOptions:     cfg.RedisOptions,
		//genericAPIServer: genericServer,
		//gRPCAPIServer:    extraServer,
	}

	return server, nil
}

func (s *apiServer) PrepareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	s.initRedisStore()

	// TODO
	//s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
	//	mysqlStore, _ := mysql.GetMySQLFactoryOr(nil)
	//	if mysqlStore != nil {
	//		_ = mysqlStore.Close()
	//	}
	//
	//	s.gRPCAPIServer.Close()
	//	s.genericAPIServer.Close()
	//
	//	return nil
	//}))

	return preparedAPIServer{s}
}

func (s *apiServer) initRedisStore() {
	ctx, _ := context.WithCancel(context.Background())
	// TODO
	//s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
	//	cancel()
	//
	//	return nil
	//}))

	config := &storage.Config{
		Host:                  s.redisOptions.Host,
		Port:                  s.redisOptions.Port,
		Addrs:                 s.redisOptions.Addrs,
		MasterName:            s.redisOptions.MasterName,
		Username:              s.redisOptions.Username,
		Password:              s.redisOptions.Password,
		Database:              s.redisOptions.Database,
		MaxIdle:               s.redisOptions.MaxIdle,
		MaxActive:             s.redisOptions.MaxActive,
		Timeout:               s.redisOptions.Timeout,
		EnableCluster:         s.redisOptions.EnableCluster,
		UseSSL:                s.redisOptions.UseSSL,
		SSLInsecureSkipVerify: s.redisOptions.SSLInsecureSkipVerify,
	}

	// try to connect to redis
	go storage.ConnectToRedis(ctx, config)
}

func (s preparedAPIServer) Run() error {
	//go s.gRPCAPIServer.Run()

	// start shutdown managers
	// TODO
	//if err := s.gs.Start(); err != nil {
	//	log.Fatalf("start shutdown manager failed: %s", err.Error())
	//}

	return s.genericAPIServer.Run()
}
