package com.abc.microservice.serviceRegister.config;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;

import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.sns.AmazonSNSClient;
import com.amazonaws.services.sns.AmazonSNSClientBuilder;

@Configuration
public class AWSSNSConfig {

	
	  @Value("${cloud.aws.region.static}")
	    private String region;

	    @Value("${cloud.aws.credentials.access-key}")
	    private String awsAccessKey;

	    @Value("${cloud.aws.credentials.secret-key}")
	    private String awsSecretKey;
	    
	    @Primary
	    @Bean
	    public AmazonSNSClient getSnsClient() {
	        return (AmazonSNSClient) AmazonSNSClientBuilder.standard().withRegion(Regions.US_WEST_2)
	                .withCredentials(new AWSStaticCredentialsProvider(new BasicAWSCredentials(awsAccessKey,
	                		awsSecretKey)))
	                .build();
	    }
}
