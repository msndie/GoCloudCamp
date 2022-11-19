package ru.zvmkm;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import org.springframework.beans.factory.annotation.Autowired;

import java.io.IOException;

@org.springframework.stereotype.Service
public class Service {
    private final ConfigServiceImpl configService;
    private final Logger logger = LoggerFactory.getLogger(Service.class);

    @Autowired
    public Service(ConfigServiceImpl configService) {
        this.configService = configService;
    }

    public void runService() throws IOException, InterruptedException {
        Server server = ServerBuilder
                .forPort(9090)
                .addService(configService).build();
        logger.info("Starting server...");
        server.start();
        logger.info("Server started");
        server.awaitTermination();
    }
}
