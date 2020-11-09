package com.abc.microservice.serviceRegister.controller;

import java.io.IOException;
import java.util.List;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.aws.messaging.core.QueueMessagingTemplate;
import org.springframework.messaging.support.MessageBuilder;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.abc.microservice.serviceRegister.service.impl.*;
import com.amazonaws.services.sns.AmazonSNSClient;
import com.amazonaws.services.sns.model.PublishRequest;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.abc.microservice.serviceRegister.repository.entity.RegistroServicio;
import com.abc.microservice.serviceRegister.controller.model.*;


@RestController
@RequestMapping("/registroSerVicio")
public class RegistroServicioController {
  
	Logger logger= LoggerFactory.getLogger(RegistroServicioController.class);
	
	 @Value("${cloud.aws.end-point.arnSNS}")
	  private String snsTopic;
	@Autowired
	RegistroServicioService service;
	
	@Autowired
	private QueueMessagingTemplate queueMessagingTemplate;
	
	@Autowired
    private SNSPublishService publishSNSService;
	
	@PostMapping("/crearRegistroServicio")
    public RegistroServicio crearRegistroServicio (@RequestBody RegistroServicio registroServicio)
    {
		
		RegistroServicio createdRegistroServicio= service.saveRegistroServicio(registroServicio);
		logger.info("piblish message:" + publishSNSService.snsPublish(snsTopic, createdRegistroServicio));
				
		return createdRegistroServicio;
    }
	@GetMapping("/registrosServicio")
	public List<RegistroServicio> registrosServicio ()
	{
		return service.getAllRegistroServicio();
	}
	
	@GetMapping("/registroServicioByIdOrdenServicio/{idOrdenServicio}")
	public List<RegistroServicio> registroServicioByIdOrdenServicio (@PathVariable int idOrdenServicio)
	{
		return service.findRegistroServicioByIdOrdenServicio(idOrdenServicio);
	}
	@PutMapping("/updateRegistroServicio")
    public RegistroServicio updateProduct(@RequestBody RegistroServicio registroServicio) {
        return service.updateRegistroServicio(registroServicio);
    }
	@GetMapping("/registrosServicioIsOn")
	public String registrosServicioIsOn ()
	{
		return "on";
	}
	
	
	
}
