package ru.zvmkm.services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import ru.zvmkm.grpc.Config;
import ru.zvmkm.models.ConfigEntity;
import ru.zvmkm.repositories.ConfigRepository;
import ru.zvmkm.repositories.ConfigVersionsRepository;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

@Service
public class ConfigurationsServiceImpl implements ConfigurationsService {

    private final ConfigRepository configRepository;
    private final ConfigVersionsRepository configVersionsRepository;

    @Autowired
    public ConfigurationsServiceImpl(ConfigRepository repository,
                                     ConfigVersionsRepository configVersionsRepository) {
        this.configRepository = repository;
        this.configVersionsRepository = configVersionsRepository;
    }

    @Override
    public boolean addConfig(Config config) {
        return configRepository.add(ConfigEntity.fromConfig(config));
    }

    @Override
    public Optional<Config> deleteConfig(String name) {
        ConfigEntity entity = configRepository.findConfig(name);
        if (entity == null) {
            return Optional.empty();
        }
        configRepository.delete(name);
        configVersionsRepository.delete(name);
        return Optional.of(ConfigEntity.fromConfigEntity(entity));
    }

    @Override
    public Optional<Config> findConfig(String name) {
        ConfigEntity entity = configRepository.findConfig(name);
        if (entity == null) {
            return Optional.empty();
        }
        return Optional.of(ConfigEntity.fromConfigEntity(entity));
    }

    @Override
    public List<Config> findAllVersions(String name) {
        ConfigEntity entity = configRepository.findConfig(name);
        List<Config> list = new ArrayList<>();
        if (entity == null) {
            return list;
        }
        list.add(ConfigEntity.fromConfigEntity(entity));
        List<Config> configEntityList = configVersionsRepository
                .findAllConfigsWithKey(entity.getService())
                .stream()
                .map(ConfigEntity::fromConfigEntity)
                .collect(Collectors.toList());
        Collections.reverse(configEntityList);
        list.addAll(configEntityList);
        return list;
    }

    @Override
    public Optional<Config> updateConfig(Config config) {
        ConfigEntity oldEntity = configRepository.findConfig(config.getService());
        if (oldEntity == null) {
            return Optional.empty();
        }
        ConfigEntity newEntity = ConfigEntity.fromConfig(config);
        if (!oldEntity.getData().equals(newEntity.getData())) {
            configRepository.update(newEntity);
            configVersionsRepository.add(config.getService(),
                    config.getService() + " " + LocalDateTime.now(),
                    oldEntity);
        }
        return Optional.of(config);
    }

    @Override
    public List<Config> findAllConfigs() {
        return configRepository.findAllConfigs().stream()
                .map(ConfigEntity::fromConfigEntity)
                .collect(Collectors.toList());
    }
}
