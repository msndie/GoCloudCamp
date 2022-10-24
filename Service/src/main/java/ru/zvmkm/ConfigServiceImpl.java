package ru.zvmkm;

import com.google.protobuf.Empty;
import io.grpc.Context;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import ru.zvmkm.grpc.*;
import ru.zvmkm.services.ConfigurationsService;

import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

@Component
public class ConfigServiceImpl extends ConfigServiceGrpc.ConfigServiceImplBase {

    private static final Logger logger = LoggerFactory.getLogger(ConfigServiceImpl.class);
    private final ConfigurationsService service;
    private static final Map<String, Set<StreamObserver<Config>>> observers = new ConcurrentHashMap<>();

    @Autowired
    public ConfigServiceImpl(ConfigurationsService service) {
        this.service = service;
    }

    @Override
    public void addConfig(Config request, StreamObserver<Config> responseObserver) {
        if (request.getService().trim().isEmpty() || request.getDataList().isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
        } else if (!service.addConfig(request)) {
            responseObserver.onError(new StatusRuntimeException(Status.ALREADY_EXISTS));
        } else {
            responseObserver.onNext(request);
            responseObserver.onCompleted();
        }
        logger.info("New config for " + request.getService() + " has been added");
    }

    @Override
    public void getConfig(ConfigNameRequest request, StreamObserver<Config> responseObserver) {
        if (request.getService().trim().isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
        } else {
            Optional<Config> config = service.findConfig(request.getService());
            if (!config.isPresent()) {
                responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
            } else {
                responseObserver.onNext(config.get());
                responseObserver.onCompleted();
            }
        }
    }

    @Override
    public void getAllVersionsOfConfig(ConfigNameRequest request, StreamObserver<Configs> responseObserver) {
        if (request.getService().trim().isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
        } else {
            List<Config> list = service.findAllVersions(request.getService());
            if (list.isEmpty()) {
                responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
            } else {
                responseObserver.onNext(Configs.newBuilder().addAllConfigs(list).build());
                responseObserver.onCompleted();
            }
        }
    }

    @Override
    public void getAllConfigs(Empty request, StreamObserver<Configs> responseObserver) {
        responseObserver.onNext(Configs.newBuilder().addAllConfigs(service.findAllConfigs()).build());
        responseObserver.onCompleted();
    }

    @Override
    public void updateConfig(Config request, StreamObserver<Config> responseObserver) {
        if (request.getService().trim().isEmpty() || request.getDataList().isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
            return;
        }
        Optional<Config> config = service.updateConfig(request);
        if (!config.isPresent()) {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        } else {
            Set<StreamObserver<Config>> observersSet = observers.get(config.get().getService());
            if (observersSet != null) {
                observersSet.forEach(observer -> observer.onNext(config.get()));
            }
            responseObserver.onNext(config.get());
            responseObserver.onCompleted();
            logger.info(request.getService() + " has been updated");
        }
    }

    @Override
    public void deleteConfig(ConfigNameRequest request, StreamObserver<Config> responseObserver) {
        if (request.getService().trim().isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
            return;
        } else if (observers.containsKey(request.getService())) {
            responseObserver.onError(new StatusRuntimeException(Status.FAILED_PRECONDITION));
            return;
        }
        Optional<Config> config = service.deleteConfig(request.getService());
        if (!config.isPresent()) {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        } else {
            responseObserver.onNext(config.get());
            responseObserver.onCompleted();
            logger.info("Config for " + request.getService() + " has been deleted");
        }
    }

    @Override
    public void useConfig(ConfigNameRequest request, StreamObserver<Config> responseObserver) {
        if (request.getService().trim().isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
            return;
        }
        Optional<Config> config = service.findConfig(request.getService());
        if (!config.isPresent()) {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        } else {
            Set<StreamObserver<Config>> set = observers.computeIfAbsent(request.getService(), s -> ConcurrentHashMap.newKeySet());
            Context.current().addListener(context -> {},
                    command -> {
                        Set<StreamObserver<Config>> observersSet = observers.get(request.getService());
                        if (observersSet != null) {
                            observersSet.remove(responseObserver);
                            if (observersSet.isEmpty()) {
                                observers.remove(request.getService());
                                logger.info("No subscribers for " + request.getService() + " left");
                            }
                        }
            });
            set.add(responseObserver);
            responseObserver.onNext(config.get());
            logger.info("New subscriber for " + request.getService());
        }
    }

    @Override
    public void stopConfigUseForAll(ConfigNameRequest request, StreamObserver<Empty> responseObserver) {
        if (request.getService().trim().isEmpty() || !observers.containsKey(request.getService())) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
        } else {
            Set<StreamObserver<Config>> observersSet = observers.get(request.getService());
            observersSet.forEach(StreamObserver::onCompleted);
            observers.remove(request.getService());
            responseObserver.onNext(Empty.getDefaultInstance());
            responseObserver.onCompleted();
            logger.info("All subscribers for " + request.getService() + " has been unsubscribed");
        }
    }
}
