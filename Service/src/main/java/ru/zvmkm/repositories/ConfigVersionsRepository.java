package ru.zvmkm.repositories;

import ru.zvmkm.models.ConfigEntity;

import java.util.List;

public interface ConfigVersionsRepository {
    void add(String key, String hashKey, ConfigEntity entity);
    List<ConfigEntity> findAllConfigsStartsWith(String key);
    void delete(String key);
}
