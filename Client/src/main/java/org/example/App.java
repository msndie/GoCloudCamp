package org.example;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import ru.zvmkm.grpc.Config;
import ru.zvmkm.grpc.ConfigNameRequest;
import ru.zvmkm.grpc.ConfigServiceGrpc;

import java.util.Iterator;

public class App {
    public static void main( String[] args ) {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 9090).usePlaintext().build();

        ConfigServiceGrpc.ConfigServiceBlockingStub stub = ConfigServiceGrpc.newBlockingStub(channel);

        ConfigNameRequest request = ConfigNameRequest.newBuilder().setService("Ut").build();

        Iterator<Config> iterator = stub.useConfig(request);
        while (iterator.hasNext()) {
            System.out.println(iterator.next());
        }
    }
}
