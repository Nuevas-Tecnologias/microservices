package com.abc.microservice.serviceRegister.service;

import com.abc.microservice.serviceRegister.repository.entity.*;
public interface ISNSPublishService {

	public String snsPublish(String topic, RegistroServicio registro);
	;
}
