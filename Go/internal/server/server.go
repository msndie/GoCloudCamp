package server

import (
	"cmd/internal/service"
	pb "cmd/proto"
	"container/list"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"sync"
)

type Server struct {
	pb.UnimplementedConfigServiceServer
	service *service.DistributedConfigService
	subs    map[string]*list.List
	lock    *sync.RWMutex
}

func NewServer(service *service.DistributedConfigService) *Server {
	var s Server
	s.subs = make(map[string]*list.List)
	s.service = service
	s.lock = &sync.RWMutex{}
	return &s
}

type sub struct {
	stream pb.ConfigService_UseConfigServer
	cfg    chan<- *pb.Config
}

func removeFromSubs(s *Server, e *list.Element, service string) {
	s.lock.Lock()
	if l, ok := s.subs[service]; ok {
		l.Remove(e)
		if l.Len() == 0 {
			delete(s.subs, service)
		}
	}
	s.lock.Unlock()
}

func broadcastConfig(s *Server, in *pb.Config) {
	s.lock.RLock()
	if l, ok := s.subs[in.Service]; ok {
		for e := l.Front(); e != nil; e = e.Next() {
			e.Value.(sub).cfg <- in
		}
	}
	s.lock.RUnlock()
}

func (s *Server) AddConfig(ctx context.Context, in *pb.Config) (*pb.Config, error) {
	b, err := s.service.AddConfig(in)
	if err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.Internal, "Error occurred")
	}
	if b {
		return in, nil
	}
	return nil, status.Errorf(codes.AlreadyExists, "Config already exists")
}

func (s *Server) GetConfig(ctx context.Context, in *pb.ConfigNameRequest) (*pb.Config, error) {
	c, err := s.service.FindConfig(in.GetService())
	if err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.Internal, "Error occurred")
	}
	if c != nil {
		return c, nil
	}
	return nil, status.Errorf(codes.NotFound, "Config not found")
}

func (s *Server) GetAllVersionsOfConfig(ctx context.Context, in *pb.ConfigNameRequest) (*pb.Configs, error) {
	configs, err := s.service.GetAllVersionsOfConfig(in.Service)
	if err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.Internal, "Error occurred")
	}
	if configs != nil {
		return &pb.Configs{Configs: configs}, nil
	}
	return nil, status.Errorf(codes.NotFound, "Config not found")
}

func (s *Server) GetAllConfigs(ctx context.Context, in *emptypb.Empty) (*pb.Configs, error) {
	configs, err := s.service.GetAllConfigs()
	if err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.Internal, "Error occurred")
	}
	return &pb.Configs{Configs: configs}, nil
}

func (s *Server) UpdateConfig(ctx context.Context, in *pb.Config) (*pb.Config, error) {
	b, err := s.service.IsExistsByName(in.Service)
	if err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.Internal, "Error occurred")
	} else if !b {
		return nil, status.Errorf(codes.NotFound, "Config not found")
	}

	if b, err = s.service.AddConfig(in); err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.Internal, "Error occurred")
	} else if b {
		broadcastConfig(s, in)
	}
	return in, nil
}

func (s *Server) DeleteConfig(ctx context.Context, in *pb.ConfigNameRequest) (*pb.Config, error) {
	_, ok := s.subs[in.Service]
	if ok {
		return nil, status.Errorf(codes.FailedPrecondition, "config cannot be deleted")
	} else {
		config, err := s.service.FindConfig(in.Service)
		if err != nil {
			log.Print(err.Error())
			return nil, status.Errorf(codes.Internal, "error occurred")
		} else if config != nil {
			if _, err = s.service.DeleteConfig(config.Service); err != nil {
				log.Print(err.Error())
				return nil, status.Errorf(codes.Internal, "error occurred")
			}
			return config, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "no such config")
}

func (s *Server) UseConfig(in *pb.ConfigNameRequest, stream pb.ConfigService_UseConfigServer) error {
	cfg, err := s.service.FindConfig(in.GetService())
	if err != nil {
		log.Print(err.Error())
		return status.Errorf(codes.Internal, "Error occurred")
	} else if cfg == nil {
		return status.Errorf(codes.NotFound, "Config not found")
	}

	s.lock.Lock()
	var subscriber sub
	subscriber.stream = stream
	ch := make(chan *pb.Config)
	subscriber.cfg = ch
	v, ok := s.subs[in.Service]
	if !ok {
		v = list.New()
		s.subs[in.Service] = v
	}
	e := v.PushBack(subscriber)
	s.lock.Unlock()
	ctx := stream.Context()

	if err = stream.Send(cfg); err != nil {
		log.Print(err.Error())
		removeFromSubs(s, e, in.Service)
		return status.Errorf(codes.Internal, "Error occurred")
	}

	for {
		select {
		case <-ctx.Done():
			log.Print("Client disconnected. Reason: ", ctx.Err().Error())
			removeFromSubs(s, e, in.Service)
			return nil
		case res := <-ch:
			if res == nil {
				log.Print("Client forcefully disconnected")
				removeFromSubs(s, e, in.Service)
				return nil
			}
			err = stream.Send(res)
			if err != nil {
				removeFromSubs(s, e, in.Service)
				log.Print(err.Error())
				return status.Errorf(codes.Internal, "Error occurred")
			}
		}
	}
}

func (s *Server) StopConfigUseForAll(ctx context.Context, in *pb.ConfigNameRequest) (*emptypb.Empty, error) {
	v, ok := s.subs[in.Service]
	if ok {
		s.lock.Lock()
		for e := v.Front(); e != nil; e = e.Next() {
			e.Value.(sub).cfg <- nil
		}
		v.Init()
		delete(s.subs, in.Service)
		s.lock.Unlock()
	}
	return &emptypb.Empty{}, nil
}
