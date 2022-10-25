package ru.zvmkm.models;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import ru.zvmkm.grpc.Config;
import ru.zvmkm.grpc.Property;

import java.io.Serializable;
import java.time.Instant;
import java.util.List;
import java.util.stream.Collectors;

@NoArgsConstructor
@AllArgsConstructor
@Getter
@Setter
public class ConfigEntity implements Serializable {
    private String service;
    private List<Pair> data;
    private Instant instant;

    public static ConfigEntity fromConfig(Config config) {
        if (config == null) {
            return null;
        }
        return new ConfigEntity(config.getService(),
                config.getDataList()
                        .stream()
                        .map(property -> new Pair(property.getKey(), property.getValue()))
                        .collect(Collectors.toList()),
                Instant.now());
    }

    public static Config fromConfigEntity(ConfigEntity entity) {
        if (entity == null) {
            return null;
        }
        return Config.newBuilder()
                .setService(entity.getService())
                .addAllData(entity.getData()
                        .stream()
                        .map(pair -> {
                            return Property.newBuilder().setKey(pair.getKey()).setValue(pair.getValue()).build();
                        }).collect(Collectors.toList()))
                .build();
    }
}
