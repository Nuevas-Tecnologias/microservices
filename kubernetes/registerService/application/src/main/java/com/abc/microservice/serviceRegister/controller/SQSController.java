package com.abc.microservice.serviceRegister.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.aws.messaging.core.QueueMessagingTemplate;
import org.springframework.cloud.aws.messaging.listener.annotation.SqsListener;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.messaging.MessageHeaders;
import org.springframework.messaging.support.MessageBuilder;
import com.abc.microservice.serviceRegister.controller.model.*;
import com.abc.microservice.serviceRegister.repository.entity.RegistroServicio;
import com.abc.microservice.serviceRegister.service.impl.RegistroServicioService;
import com.abc.microservice.serviceRegister.service.impl.SNSPublishService;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.gson.Gson;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@RestController
@RequestMapping("/sqs")
public class SQSController {

	Logger logger= LoggerFactory.getLogger(SQSController.class);
	
	@Autowired
	RegistroServicioService service;
	
	@Autowired
    private SNSPublishService publishSNSService;
	
	@Autowired
	private QueueMessagingTemplate queueMessagingTemplate;
	
	
	
	 @Value("${cloud.aws.end-point.uriReplay}")
	  private String endpointReplay;
	 @Value("${cloud.aws.end-point.arnSNS}")
	  private String snsTopic;
	 
	
	public String sendReplay(CommandCreateRegistroInfo comamadInfo, List<Integer> idsRegistro)
		{
		 ReplayFormatoCommand replay = new ReplayFormatoCommand();
		 
		 replay.setType("TechServiceCreated");
		 replay.setCorrelationId(comamadInfo.getCorrelationId());
		 replay.setTechServiceId(idsRegistro);
		 ObjectMapper Obj = new ObjectMapper();
		 try {
			 
	           
	            String jsonStr = Obj.writeValueAsString(replay);
	           
	            queueMessagingTemplate.send(endpointReplay, MessageBuilder.withPayload(jsonStr).setHeader("message-group-id", "Tech-revision-format").build());
	            
	            return jsonStr;
	        }
	 
	        catch (IOException e) {
	        	logger.info( "Errror convirtiendo json");
	        	 return "message not send";
	        }
		 
		 
		}
	
	  @SqsListener("https://sqs.us-west-2.amazonaws.com/881619806726/create-service-register-bulk-command.fifo")
	    public void loadMessageFromSQS(String message)  {
		  
		  Gson g = new Gson();
		  List<Integer> idsRegistros = new ArrayList<>();
		  CommandCreateRegistroInfo registroInfo = g.fromJson(message, CommandCreateRegistroInfo.class);
	      for (RegistroServicio registroServicio : registroInfo.getRegistrosServicio() )
	      {
	    	  RegistroServicio createdRegistroServicio= service.saveRegistroServicio(registroServicio);
	    	  logger.info("piblish message Topic:" + publishSNSService.snsPublish(snsTopic, createdRegistroServicio));
	    	  idsRegistros.add(createdRegistroServicio.getId());
	      }
	      logger.info("publis replay:" + sendReplay(registroInfo,idsRegistros) );
	    }
	
}
