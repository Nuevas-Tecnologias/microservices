package com.abc.microservice.serviceRegister.service.impl;

import java.io.IOException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.abc.microservice.serviceRegister.controller.model.EventRegistroServicioInfo;
import com.abc.microservice.serviceRegister.repository.entity.RegistroServicio;
import com.abc.microservice.serviceRegister.service.ISNSPublishService;
import com.amazonaws.services.sns.AmazonSNSClient;
import com.amazonaws.services.sns.model.PublishRequest;
import com.fasterxml.jackson.databind.ObjectMapper;

@Service
public class SNSPublishService implements ISNSPublishService {

	@Autowired
    private AmazonSNSClient snsClient;
	
	@Override
	public String snsPublish(String topic, RegistroServicio registro)
	{
		try {
			 ObjectMapper Obj = new ObjectMapper();
			EventRegistroServicioInfo eventRegistroInfo = new   EventRegistroServicioInfo();
			eventRegistroInfo.setType("SERVICE_CREATED");
			eventRegistroInfo.setRegistroServicio(registro);
			String jsonStr = Obj.writeValueAsString(eventRegistroInfo);
            PublishRequest publishRequest=new PublishRequest(topic,jsonStr);
   		 	snsClient.publish(publishRequest);    
   		 return  "mesageSend: " + jsonStr;
            
        }
 
        catch (IOException e) {
        	return "Errror convirtiendo json";
        	 
        }
		
	}

}
