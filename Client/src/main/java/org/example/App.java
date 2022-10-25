package org.example;

import com.google.protobuf.Empty;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import ru.zvmkm.grpc.Config;
import ru.zvmkm.grpc.ConfigNameRequest;
import ru.zvmkm.grpc.ConfigServiceGrpc;
import ru.zvmkm.grpc.Configs;

import java.util.Iterator;

public class App {
    public static void main( String[] args ) {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 9090)
                .enableRetry().maxRetryAttempts(10).usePlaintext().build();

        ConfigServiceGrpc.ConfigServiceBlockingStub stub = ConfigServiceGrpc.newBlockingStub(channel);

        Configs configs = stub.getAllConfigs(Empty.getDefaultInstance());
        if (configs.getConfigsList().isEmpty()) {
            System.err.println("Add configs first");
            return;
        }
        Config config = configs.getConfigsList().get(0);

        ConfigNameRequest request = ConfigNameRequest.newBuilder()
                .setService(config.getService()).build();

        Iterator<Config> iterator = stub.useConfig(request);
        while (iterator.hasNext()) {
            System.out.println(iterator.next());
        }
    }
}
