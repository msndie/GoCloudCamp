package ru.zvmkm.repositories;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.HashOperations;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Repository;
import ru.zvmkm.models.ConfigEntity;

import javax.annotation.PostConstruct;
import java.util.List;

@Repository
public class ConfigRepositoryImpl implements ConfigRepository {
    private static final String KEY = "Config";
    private final RedisTemplate<String, Object> template;
    private HashOperations<String, String, ConfigEntity> hashOperations;

    @Autowired
    public ConfigRepositoryImpl(RedisTemplate<String, Object> template) {
        this.template = template;
    }

    @PostConstruct
    private void init() {
        hashOperations = template.opsForHash();
    }

    @Override
    public List<ConfigEntity> findAllConfigs() {
        return hashOperations.values(KEY);
    }

    @Override
    public boolean add(ConfigEntity entity) {
        return hashOperations.putIfAbsent(KEY, entity.getService(), entity);
    }

    @Override
    public void delete(String name) {
        hashOperations.delete(KEY, name);
    }

    @Override
    public ConfigEntity findConfig(String name) {
        return hashOperations.get(KEY, name);
    }

    @Override
    public void update(ConfigEntity entity) {
        hashOperations.put(KEY, entity.getService(), entity);
    }
}
