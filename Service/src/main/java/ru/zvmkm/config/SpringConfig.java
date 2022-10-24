package ru.zvmkm.config;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.PropertySource;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.connection.RedisStandaloneConfiguration;
import org.springframework.data.redis.connection.lettuce.LettuceConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;

@Configuration
@PropertySource("classpath:application.properties")
@ComponentScan("ru.zvmkm")
public class SpringConfig {

    @Value("${redis.port}")
    private String port;

    @Value("${redis.host}")
    private String host;

    @Bean
    public RedisConnectionFactory redisConnectionFactory() {
        int portInt = Integer.parseInt(port);
        return new LettuceConnectionFactory(new RedisStandaloneConfiguration(host, portInt));
    }

    @Bean
    public RedisTemplate<String, Object> redisTemplate() {
        RedisTemplate<String, Object> template = new RedisTemplate<>();
        template.setConnectionFactory(redisConnectionFactory());
        template.getConnectionFactory().getConnection().close();
        return template;
    }
}
