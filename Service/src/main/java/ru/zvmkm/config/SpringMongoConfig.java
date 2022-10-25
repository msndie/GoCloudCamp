package ru.zvmkm.config;

import com.mongodb.ConnectionString;
import com.mongodb.MongoClientSettings;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.PropertySource;
import org.springframework.data.mongodb.config.AbstractMongoClientConfiguration;
import org.springframework.util.Assert;

import javax.annotation.Nonnull;
import java.util.Collection;
import java.util.Collections;

@Configuration
@PropertySource("classpath:application.properties")
@ComponentScan("ru.zvmkm")
public class SpringMongoConfig extends AbstractMongoClientConfiguration {

    @Value("${mongo.connection}")
    private String connection;

    @Value("${mongo.database}")
    private String db;

    @Override
    @Nonnull
    protected String getDatabaseName() {
        Assert.isTrue(db != null, "Database name must be specified in application.properties file");
        return db;
    }

    @Override
    @Nonnull
    public MongoClient mongoClient() {
        Assert.isTrue(connection != null, "Mongo URL must be specified in application.properties file");
        ConnectionString connectionString = new ConnectionString(connection);
        MongoClientSettings settings = MongoClientSettings.builder()
                .applyConnectionString(connectionString)
                .build();
        return MongoClients.create(settings);
    }

    @Override
    @Nonnull
    protected Collection<String> getMappingBasePackages() {
        return Collections.singleton("ru.zvmkm");
    }
}
