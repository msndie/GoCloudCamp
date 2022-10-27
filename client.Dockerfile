FROM maven:3.8.6-amazoncorretto-8 AS build
COPY ./Client /home/Client
RUN mvn --file /home/Client/pom.xml clean package
CMD java -jar /home/Client/target/Client-1.0-jar-with-dependencies.jar