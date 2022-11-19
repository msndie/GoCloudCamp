package ru.zvmkm.repositories;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Sort;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.stereotype.Repository;
import ru.zvmkm.models.ConfigEntity;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.Set;

@Repository
public class ConfigRepositoryImp implements ConfigRepository{

    private final MongoTemplate template;

    @Autowired
    public ConfigRepositoryImp(MongoTemplate template) {
        this.template = template;
    }

    @Override
    public boolean collectionExists(String name) {
        return template.collectionExists(name);
    }

    @Override
    public void createCollection(String name) {
        template.createCollection(name);
    }

    @Override
    public void insertInCollection(ConfigEntity entity, String collection) {
        template.insert(entity, collection);
    }

    @Override
    public void deleteCollection(String name) {
        template.dropCollection(name);
    }

    @Override
    public Optional<ConfigEntity> findLastEntityInCollection(String name) {
        Query query = new Query();
        query.with(Sort.by(Sort.Direction.DESC, "instant"));
        ConfigEntity entity = template.findOne(query, ConfigEntity.class, name);
        return Optional.ofNullable(entity);
    }

    @Override
    public List<ConfigEntity> findAllEntitiesInCollection(String name) {
        Query query = new Query();
        query.with(Sort.by(Sort.Direction.DESC, "instant"));
        return template.find(query, ConfigEntity.class, name);
    }

    @Override
    public List<ConfigEntity> findAllLastEntities() {
        Set<String> set = template.getCollectionNames();
        List<ConfigEntity> list = new ArrayList<>();
        set.forEach(s -> findLastEntityInCollection(s).ifPresent(list::add));
        return list;
    }
}
