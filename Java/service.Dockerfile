FROM maven:3.8.6-amazoncorretto-8 AS build
COPY Service /home/Service
# RUN rm ~/Service/src/main/resources/application.properties
COPY Service/docker/application.properties /home/Service/src/main/resources/application.properties
RUN mvn --file /home/Service/pom.xml clean package

# FROM amazoncorretto:8-alpine3.13-jre
# COPY --from=build /home/Service/target/ConfigurationService-1.0-jar-with-dependencies.jar /home/Service.jar
# RUN java -jar /home/Service.jar