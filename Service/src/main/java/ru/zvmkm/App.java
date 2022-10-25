package ru.zvmkm;

import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import ru.zvmkm.config.SpringMongoConfig;

import java.io.IOException;

public class App {

    public static void main( String[] args ) throws IOException, InterruptedException {
        ApplicationContext context = new AnnotationConfigApplicationContext(SpringMongoConfig.class);
        Service service = context.getBean(Service.class);
        service.runService();
    }
}
