package ru.zvmkm.services;

import ru.zvmkm.grpc.Config;

import java.util.List;
import java.util.Optional;

public interface ConfigurationsService {
    boolean addConfig(Config config);
    Optional<Config> deleteConfig(String name);
    Optional<Config> findConfig(String name);
    List<Config> findAllVersions(String name);
    Optional<Config> updateConfig(Config config);
    List<Config> findAllConfigs();
}
