FROM openjdk:8-jdk-alpine
VOLUME /tmp
ARG JAR_FILE=src/*.jar
COPY ${JAR_FILE} serviceRegisterMS.jar
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/serviceRegisterMS.jar"]