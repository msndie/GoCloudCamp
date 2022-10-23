package ru.zvmkm;

import com.google.protobuf.Empty;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import ru.zvmkm.grpc.*;
import ru.zvmkm.services.ConfigurationsService;

import java.util.List;
import java.util.Optional;

@Component
public class ConfigServiceImpl extends ConfigServiceGrpc.ConfigServiceImplBase {

    private final ConfigurationsService service;

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
            responseObserver.onNext(service.findConfig(request.getService()).get());
            responseObserver.onCompleted();
        }
    }

    @Override
    public void getConfig(ConfigNameRequest request, StreamObserver<Config> responseObserver) {
        if (request.getService().trim().isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
            return;
        }
        Optional<Config> config = service.findConfig(request.getService());
        if (!config.isPresent()) {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        } else {
            responseObserver.onNext(config.get());
            responseObserver.onCompleted();
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
            responseObserver.onNext(config.get());
            responseObserver.onCompleted();
        }
    }

    @Override
    public void deleteConfig(ConfigNameRequest request, StreamObserver<Config> responseObserver) {
        if (request.getService().trim().isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
            return;
        }
        Optional<Config> config = service.deleteConfig(request.getService());
        if (!config.isPresent()) {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        } else {
            responseObserver.onNext(config.get());
            responseObserver.onCompleted();
        }
    }

    @Override
    public void useConfig(ConfigNameRequest request, StreamObserver<Config> responseObserver) {
        if (request.getService().trim().isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
            return;
        }
        Optional<Config> config = service.markAsOnUse(request.getService());
        if (!config.isPresent()) {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        } else {
            responseObserver.onNext(config.get());
            responseObserver.onCompleted();
        }
    }

    @Override
    public void stopConfigUse(ConfigNameRequest request, StreamObserver<Empty> responseObserver) {
        if (request.getService().trim().isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
            return;
        }
        if (service.unmarkAsOnUse(request.getService())) {
            responseObserver.onNext(Empty.getDefaultInstance());
            responseObserver.onCompleted();
        } else {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
        }
    }

    @Override
    public void getAllVersionsOfConfig(ConfigNameRequest request, StreamObserver<Configs> responseObserver) {
        if (request.getService().trim().isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.INVALID_ARGUMENT));
            return;
        }
        List<Config> list = service.findAllVersions(request.getService());
        if (list.isEmpty()) {
            responseObserver.onError(new StatusRuntimeException(Status.NOT_FOUND));
        } else {
            responseObserver.onNext(Configs.newBuilder().addAllConfigs(list).build());
            responseObserver.onCompleted();
        }
    }
}
