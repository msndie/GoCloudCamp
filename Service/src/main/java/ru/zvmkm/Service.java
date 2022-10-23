package ru.zvmkm;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import org.springframework.beans.factory.annotation.Autowired;

import java.io.IOException;

@org.springframework.stereotype.Service
public class Service {
    private final ConfigServiceImpl configService;

    @Autowired
    public Service(ConfigServiceImpl configService) {
        this.configService = configService;
    }

    public void runService() throws IOException, InterruptedException {
        Server server = ServerBuilder
                .forPort(8080)
                .addService(configService).build();
        System.out.println("Starting server ...");
        server.start();
        System.out.println("Server started");
        server.awaitTermination();
    }
}
