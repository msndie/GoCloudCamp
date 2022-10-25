package ru.zvmkm.services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import ru.zvmkm.grpc.Config;
import ru.zvmkm.models.ConfigEntity;
import ru.zvmkm.repositories.ConfigRepository;

import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

@Service
public class ConfigurationsServiceImpl implements ConfigurationsService {

    private final ConfigRepository repository;

    @Autowired
    public ConfigurationsServiceImpl(ConfigRepository repository) {
        this.repository = repository;
    }

    @Override
    public boolean addConfig(Config config) {
        if (repository.collectionExists(config.getService())) {
            return false;
        }
        repository.createCollection(config.getService());
        repository.insertInCollection(ConfigEntity.fromConfig(config), config.getService());
        return true;
    }

    @Override
    public Optional<Config> deleteConfig(String name) {
        Optional<ConfigEntity> config = repository.findLastEntityInCollection(name);
        if (config.isPresent()) {
            repository.deleteCollection(name);
            return Optional.of(ConfigEntity.fromConfigEntity(config.get()));
        }
        return Optional.empty();
    }

    @Override
    public Optional<Config> findConfig(String name) {
        Optional<ConfigEntity> entity = repository.findLastEntityInCollection(name);
        return entity.map(ConfigEntity::fromConfigEntity);
    }

    @Override
    public List<Config> findAllVersions(String name) {
        return repository.findAllEntitiesInCollection(name)
                .stream()
                .map(ConfigEntity::fromConfigEntity)
                .collect(Collectors.toList());
    }

    @Override
    public boolean updateConfig(Config config) {
        Optional<ConfigEntity> oldEntity = repository.findLastEntityInCollection(config.getService());
        if (!oldEntity.isPresent()) {
            return false;
        }
        ConfigEntity newEntity = ConfigEntity.fromConfig(config);
        if (!oldEntity.get().getData().equals(newEntity.getData())) {
            repository.insertInCollection(newEntity, config.getService());
        }
        return true;
    }

    @Override
    public List<Config> findAllConfigs() {
        return repository.findAllLastEntities().stream()
                .map(ConfigEntity::fromConfigEntity)
                .collect(Collectors.toList());
    }
}
