package ru.zvmkm;

import com.google.protobuf.Empty;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.internal.testing.StreamRecorder;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import ru.zvmkm.grpc.Config;
import ru.zvmkm.grpc.ConfigNameRequest;
import ru.zvmkm.grpc.Configs;
import ru.zvmkm.grpc.Property;
import ru.zvmkm.services.ConfigurationsService;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.TimeUnit;

import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

public class AppTest {
    private static ConfigurationsService service;
    private static ConfigServiceImpl serviceImpl;
    private static Config emptyConfig;
    private static Config emptyDataConfig;
    private static Config emptyServiceNameConfig;
    private static Config validConfig;
    private static Config validConfigUpdated;
    private static ConfigNameRequest nameRequestEmpty;
    private static ConfigNameRequest nameRequestInvalid;
    private static ConfigNameRequest nameRequestValid;
    private static List<Config> allVersionsList;
    private static List<Config> allConfigs;

    @BeforeAll
    static void setup() {
        service = mock(ConfigurationsService.class);
        serviceImpl = new ConfigServiceImpl(service);
        emptyConfig = Config.newBuilder().build();
        emptyDataConfig = Config.newBuilder().setService("test").build();
        emptyServiceNameConfig = Config.newBuilder()
                .addData(Property.newBuilder().setKey("test").setValue("test").build())
                .build();
        validConfig = Config.newBuilder()
                .addData(Property.newBuilder().setKey("test").setValue("test").build())
                .setService("test")
                .build();
        validConfigUpdated = Config.newBuilder()
                .addData(Property.newBuilder().setKey("test updated").setValue("test updated").build())
                .setService("test")
                .build();
        nameRequestInvalid = ConfigNameRequest.newBuilder().setService("    ").build();
        nameRequestEmpty = ConfigNameRequest.newBuilder().build();
        nameRequestValid = ConfigNameRequest.newBuilder().setService("test").build();
        allConfigs = new ArrayList<>();
        allVersionsList = new ArrayList<>();
        allVersionsList.add(validConfig);
        allConfigs.add(validConfig);
        allConfigs.add(Config.newBuilder()
                .setService("test 2")
                .addData(Property.newBuilder().setKey("test 2").setValue("test 2").build())
                .build());
        allVersionsList.add(Config.newBuilder()
                .setService("test")
                .addData(Property.newBuilder().setKey("test old").setValue("test old").build())
                .build());
    }

    @Test
    public void shouldReturnConfig_addConfig() throws Exception {
        StreamRecorder<Config> observerConfig = StreamRecorder.create();

        when(service.addConfig(validConfig)).thenReturn(true);
        serviceImpl.addConfig(validConfig, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() == null
                && !observerConfig.getValues().isEmpty()
                && observerConfig.getValues().get(0).equals(validConfig));
    }

    @Test
    public void shouldReturnConfig_getConfig() throws Exception {
        StreamRecorder<Config> observer = StreamRecorder.create();
        when(service.findConfig(nameRequestValid.getService())).thenReturn(Optional.of(validConfig));
        serviceImpl.getConfig(nameRequestValid, observer);
        if (!observer.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observer.getValues().size() == 1
                && observer.getValues().get(0) == validConfig);
    }

    @Test
    public void shouldReturnAllVersionsOfConfig_getAllVersionsOfConfig() throws Exception {
        StreamRecorder<Configs> observer = StreamRecorder.create();
        when(service.findAllVersions(nameRequestValid.getService())).thenReturn(allVersionsList);
        serviceImpl.getAllVersionsOfConfig(nameRequestValid, observer);
        if (!observer.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observer.getValues().size() == 1
                && observer.getValues().get(0).getConfigsList().equals(allVersionsList));
    }

    @Test
    public void shouldReturnAllConfigs_getAllConfigs() throws Exception {
        StreamRecorder<Configs> observer = StreamRecorder.create();
        StreamRecorder<Configs> observerForEmpty = StreamRecorder.create();
        when(service.findAllConfigs()).thenReturn(allConfigs);
        serviceImpl.getAllConfigs(Empty.getDefaultInstance(), observer);
        if (!observer.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observer.getValues().size() == 1
                && observer.getValues().get(0).getConfigsList().equals(allConfigs));

        when(service.findAllConfigs()).thenReturn(Collections.emptyList());
        serviceImpl.getAllConfigs(Empty.getDefaultInstance(), observerForEmpty);
        if (!observerForEmpty.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerForEmpty.getValues().size() == 1
                && observerForEmpty.getValues().get(0).getConfigsList().isEmpty());
    }

    @Test
    public void multiTestForUpdateDeleteUseAndStopUse() throws Exception {
        StreamRecorder<Config> observerForDelete = StreamRecorder.create();
        StreamRecorder<Empty> observerForStop = StreamRecorder.create();
        StreamRecorder<Config> observerForUpdate = StreamRecorder.create();
        StreamRecorder<Config> configUser = StreamRecorder.create();

        when(service.findConfig(nameRequestValid.getService())).thenReturn(Optional.of(validConfig));
        serviceImpl.useConfig(nameRequestValid, configUser);
        serviceImpl.deleteConfig(nameRequestValid, observerForDelete);

        if (!observerForDelete.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerForDelete.getError() != null
                && observerForDelete.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerForDelete.getError()).getStatus() == Status.FAILED_PRECONDITION);

        when(service.updateConfig(validConfigUpdated)).thenReturn(true);
        serviceImpl.updateConfig(validConfigUpdated, observerForUpdate);
        if (!observerForUpdate.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(!observerForUpdate.getValues().isEmpty()
                && observerForUpdate.getValues().get(0).equals(validConfigUpdated));

        serviceImpl.stopConfigUseForAll(nameRequestValid, observerForStop);
        if (!observerForStop.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(!observerForStop.getValues().isEmpty()
                && observerForStop.getValues().get(0).equals(Empty.getDefaultInstance()));


        observerForDelete = StreamRecorder.create();
        when(service.deleteConfig(nameRequestValid.getService())).thenReturn(Optional.of(validConfigUpdated));
        serviceImpl.deleteConfig(nameRequestValid, observerForDelete);

        if (!observerForDelete.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(!observerForDelete.getValues().isEmpty()
                && observerForDelete.getValues().get(0).equals(validConfigUpdated));

        Assertions.assertEquals(2, configUser.getValues().size());

    }

    @Test
    public void shouldReturnError_addConfig() throws Exception {
        StreamRecorder<Config> observerConfig = StreamRecorder.create();
        serviceImpl.addConfig(emptyConfig, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        serviceImpl.addConfig(emptyServiceNameConfig, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        serviceImpl.addConfig(emptyDataConfig, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        when(service.addConfig(validConfig)).thenReturn(false);
        serviceImpl.addConfig(validConfig, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.ALREADY_EXISTS);
    }

    @Test
    public void shouldReturnError_getConfig() throws Exception {
        StreamRecorder<Config> observerConfig = StreamRecorder.create();

        serviceImpl.getConfig(nameRequestInvalid, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        serviceImpl.getConfig(nameRequestEmpty, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        when(service.findConfig(nameRequestValid.getService())).thenReturn(Optional.empty());
        serviceImpl.getConfig(nameRequestValid, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.NOT_FOUND);
    }

    @Test
    public void shouldReturnError_getAllVersionsOfConfig() throws Exception {
        StreamRecorder<Configs> observerConfigs = StreamRecorder.create();

        serviceImpl.getAllVersionsOfConfig(nameRequestInvalid, observerConfigs);
        if (!observerConfigs.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfigs.getError() != null
                && observerConfigs.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfigs.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfigs = StreamRecorder.create();
        serviceImpl.getAllVersionsOfConfig(nameRequestEmpty, observerConfigs);
        if (!observerConfigs.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfigs.getError() != null
                && observerConfigs.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfigs.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfigs = StreamRecorder.create();
        when(service.findAllVersions(nameRequestValid.getService())).thenReturn(Collections.emptyList());
        serviceImpl.getAllVersionsOfConfig(nameRequestValid, observerConfigs);
        if (!observerConfigs.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfigs.getError() != null
                && observerConfigs.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfigs.getError()).getStatus() == Status.NOT_FOUND);
    }

    @Test
    public void shouldReturnError_updateConfig() throws Exception {
        StreamRecorder<Config> observerConfig = StreamRecorder.create();

        serviceImpl.updateConfig(emptyConfig, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        serviceImpl.updateConfig(emptyServiceNameConfig, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        serviceImpl.updateConfig(emptyDataConfig, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        when(service.updateConfig(validConfig)).thenReturn(false);
        serviceImpl.updateConfig(validConfig, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.NOT_FOUND);
    }

    @Test
    public void shouldReturnError_deleteConfig() throws Exception {
        StreamRecorder<Config> observerConfig = StreamRecorder.create();

        serviceImpl.deleteConfig(nameRequestInvalid, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        serviceImpl.deleteConfig(nameRequestEmpty, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        when(service.deleteConfig(nameRequestValid.getService())).thenReturn(Optional.empty());
        serviceImpl.deleteConfig(nameRequestValid, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.NOT_FOUND);
    }

    @Test
    public void shouldReturnError_useConfig() throws Exception {
        StreamRecorder<Config> observerConfig = StreamRecorder.create();

        serviceImpl.useConfig(nameRequestInvalid, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        serviceImpl.useConfig(nameRequestEmpty, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerConfig = StreamRecorder.create();
        when(service.findConfig(nameRequestValid.getService())).thenReturn(Optional.empty());
        serviceImpl.useConfig(nameRequestValid, observerConfig);
        if (!observerConfig.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerConfig.getError() != null
                && observerConfig.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerConfig.getError()).getStatus() == Status.NOT_FOUND);
    }

    @Test
    public void shouldReturnError_allMethods() throws Exception {
        StreamRecorder<Empty> observerEmpty = StreamRecorder.create();

        serviceImpl.stopConfigUseForAll(nameRequestInvalid, observerEmpty);
        if (!observerEmpty.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerEmpty.getError() != null
                && observerEmpty.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerEmpty.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerEmpty = StreamRecorder.create();
        serviceImpl.stopConfigUseForAll(nameRequestEmpty, observerEmpty);
        if (!observerEmpty.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerEmpty.getError() != null
                && observerEmpty.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerEmpty.getError()).getStatus() == Status.INVALID_ARGUMENT);
        observerEmpty = StreamRecorder.create();
        when(service.findConfig(nameRequestValid.getService())).thenReturn(Optional.empty());
        serviceImpl.stopConfigUseForAll(nameRequestValid, observerEmpty);
        if (!observerEmpty.awaitCompletion(1, TimeUnit.SECONDS)) {
            Assertions.fail("Call did not terminate in time");
        }
        Assertions.assertTrue(observerEmpty.getError() != null
                && observerEmpty.getError() instanceof StatusRuntimeException
                && ((StatusRuntimeException) observerEmpty.getError()).getStatus() == Status.NOT_FOUND);
    }
}
