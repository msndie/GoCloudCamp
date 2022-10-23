package ru.zvmkm.models;

import javafx.util.Pair;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import ru.zvmkm.grpc.Config;
import ru.zvmkm.grpc.Property;

import java.io.Serializable;
import java.util.List;
import java.util.stream.Collectors;

@NoArgsConstructor
@AllArgsConstructor
@Getter
@Setter
public class ConfigEntity implements Serializable {
    private String service;
    private List<Pair<String, String>> data;
    private boolean onUse;

    public static ConfigEntity fromConfig(Config config) {
        return new ConfigEntity(config.getService(),
                config.getDataList()
                        .stream()
                        .map(property -> new Pair<String, String>(property.getKey(), property.getValue()))
                        .collect(Collectors.toList())
                , false);
    }

    public static Config fromConfigEntity(ConfigEntity entity) {
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
