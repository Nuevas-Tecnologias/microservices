FROM openjdk:8-jdk-alpine
VOLUME /tmp
ARG JAR_FILE=src/*.jar
COPY ${JAR_FILE} serviceRegiterMS.jar
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/serviceRegiterMS.jar"]