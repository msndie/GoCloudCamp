package ru.zvmkm.repositories;

import ru.zvmkm.models.ConfigEntity;

import java.util.List;
import java.util.Optional;

public interface ConfigRepository {
    boolean collectionExists(String name);
    void createCollection(String name);
    void insertInCollection(ConfigEntity entity, String collection);
    void deleteCollection(String name);
    Optional<ConfigEntity> findLastEntityInCollection(String name);
    List<ConfigEntity> findAllEntitiesInCollection(String name);
    List<ConfigEntity> findAllLastEntities();
}
