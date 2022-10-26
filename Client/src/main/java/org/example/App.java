package org.example;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import ru.zvmkm.grpc.*;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

public class App {

    private static String newDataMessage;
    private static String endMessage;
    private static String mainMessage;

    private static Config createDefaultConfig() {
        List<Property> properties = new ArrayList<>();
        properties.add(Property.newBuilder().setKey("new data message").setValue("New data").build());
        properties.add(Property.newBuilder().setKey("end message").setValue("Application stops").build());
        properties.add(Property.newBuilder().setKey("main thread message").setValue("In main, waiting...").build());
        return Config.newBuilder().setService("Test app").addAllData(properties).build();
    }

    private static Config createUpdatedConfig() {
        List<Property> properties = new ArrayList<>();
        properties.add(Property.newBuilder().setKey("new data message").setValue("NeW DatA ArriveD").build());
        properties.add(Property.newBuilder().setKey("end message").setValue("Exit success").build());
        properties.add(Property.newBuilder().setKey("main thread message").setValue("Still in main, still waiting...").build());
        return Config.newBuilder().setService("Test app").addAllData(properties).build();
    }

    private static void assignMessages(Config config) {
        String tmpDataMessage = "";
        String tmpEndMessage = "";
        String tmpMainMessage = "";
        List<Property> properties = config.getDataList();

        for (Property p: properties) {
            if (p.getKey().equals("new data message")) {
                tmpDataMessage = p.getValue();
            } else if (p.getKey().equals("end message")) {
                tmpEndMessage = p.getValue();
            } else if (p.getKey().equals("main thread message")) {
                tmpMainMessage = p.getValue();
            }
        }
        System.out.println("new data message now is - " + tmpDataMessage);
        System.out.println("end message now is - " + tmpEndMessage);
        System.out.println("main message now is - " + tmpMainMessage);
        newDataMessage = tmpDataMessage;
        endMessage = tmpEndMessage;
        mainMessage = tmpMainMessage;
    }

    public static void main( String[] args ) throws InterruptedException {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 9090)
                .enableRetry().maxRetryAttempts(10).usePlaintext().build();

        ConfigServiceGrpc.ConfigServiceBlockingStub stub = ConfigServiceGrpc.newBlockingStub(channel);

        Config defaultConfig = createDefaultConfig();

        try {
            stub.addConfig(defaultConfig);
        } catch (StatusRuntimeException e) {
            if (e.getStatus() == Status.ALREADY_EXISTS) {
                try {
                    stub.updateConfig(defaultConfig);
                } catch (StatusRuntimeException ex) {
                    System.err.println(e.getMessage());
                    return;
                }
            } else {
                System.err.println(e.getMessage());
                return;
            }
        }

        assignMessages(defaultConfig);

        System.out.println("Subscribing for " + defaultConfig.getService());

//        Thread thread = new Thread(() -> {
//            ConfigNameRequest request = ConfigNameRequest.newBuilder()
//                    .setService(config.getService()).build();
//            try {
//                Iterator<Config> iterator = stub.useConfig(request);
//                while (iterator.hasNext()) {
//                    System.out.println(newDataMessage + "\n----------------------------------------\n");
//                    Config tmp = iterator.next();
//                    assignMessages(tmp);
//                    System.out.println(tmp);
//                    System.out.println("----------------------------------------");
//                }
//                System.out.println("Unsubscribed from " + config.getService());
//            } catch (StatusRuntimeException e) {
//                System.err.println(e.getMessage());
//            }
//        });
//        thread.start();

        ConfigServiceGrpc.ConfigServiceStub nonBlockStub = ConfigServiceGrpc.newStub(channel);
        nonBlockStub.useConfig(ConfigNameRequest.newBuilder().setService(defaultConfig.getService()).build(),
                new StreamObserver<Config>() {
                    @Override
                    public void onNext(Config config) {
                        System.out.println(newDataMessage + "\n----------------------------------------\n");
                        assignMessages(config);
                        System.out.println(config);
                        System.out.println("----------------------------------------");
                    }

                    @Override
                    public void onError(Throwable throwable) {
                        System.err.println(throwable.getMessage());
                        System.exit(1);
                    }

                    @Override
                    public void onCompleted() {
                        System.out.println("Unsubscribed from " + defaultConfig.getService());
                    }
                });

        for (int i = 0; i < 5; ++i) {
            Thread.sleep(500);
            System.out.println(mainMessage);
        }
        Config updated = createUpdatedConfig();
        try {
            stub.updateConfig(updated);
        } catch (StatusRuntimeException e) {
            System.err.println(e.getMessage());
        }
        for (int i = 0; i < 5; ++i) {
            Thread.sleep(500);
            System.out.println(mainMessage);
        }

        System.out.println("Executes stopConfigUseForAll");
        try {
            stub.stopConfigUseForAll(ConfigNameRequest.newBuilder().setService(defaultConfig.getService()).build());
        } catch (StatusRuntimeException ex) {
            System.err.println(ex.getMessage());
        }

//        thread.join();
        System.out.println(endMessage);
    }
}
