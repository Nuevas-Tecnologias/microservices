FROM maven:3.6.0-jdk-11-slim AS build

WORKDIR /app
COPY ./ /app

RUN mvn package

FROM openjdk:11-jre-slim
COPY --from=build /app/target/serviceRegister-0.0.1-SNAPSHOT.jar /app/serviceRegister.jar
EXPOSE 8080

CMD ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/app/serviceRegister.jar"]