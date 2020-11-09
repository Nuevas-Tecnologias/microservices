package com.abc.microservice.serviceRegister;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.aws.autoconfigure.context.ContextStackAutoConfiguration;


@SpringBootApplication (exclude = {ContextStackAutoConfiguration.class})
public class ServiceRegisterApplication  {

	public static void main(String[] args) {
		SpringApplication.run(ServiceRegisterApplication.class, args);
	}
}
