package ru.zvmkm.repositories;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.HashOperations;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Repository;
import ru.zvmkm.models.ConfigEntity;

import javax.annotation.PostConstruct;
import java.util.List;

@Repository
public class ConfigVersionsRepositoryImpl implements ConfigVersionsRepository {
    private final RedisTemplate<String, Object> template;
    private HashOperations<String, String, ConfigEntity> hashOperations;

    @Autowired
    public ConfigVersionsRepositoryImpl(RedisTemplate<String, Object> template) {
        this.template = template;
    }

    @PostConstruct
    private void init() {
        hashOperations = template.opsForHash();
    }

    @Override
    public void add(String key, String hashKey, ConfigEntity entity) {
        hashOperations.put(key, hashKey, entity);
    }

    @Override
    public List<ConfigEntity> findAllConfigsWithKey(String key) {
        return hashOperations.values(key);
    }

    @Override
    public void delete(String key) {
        template.delete(key);
    }
}
