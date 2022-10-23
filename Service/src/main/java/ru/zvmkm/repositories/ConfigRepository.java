package ru.zvmkm.repositories;

import ru.zvmkm.models.ConfigEntity;

import java.util.List;

public interface ConfigRepository {
    List<ConfigEntity> findAllConfigs();
    boolean add(ConfigEntity entity);
    void delete(String name);
    ConfigEntity findConfig(String name);
    void update(ConfigEntity entity);
}
