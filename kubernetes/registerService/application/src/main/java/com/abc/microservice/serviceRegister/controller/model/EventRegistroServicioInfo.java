package com.abc.microservice.serviceRegister.controller.model;

import com.abc.microservice.serviceRegister.repository.entity.RegistroServicio;

public class EventRegistroServicioInfo {
	public String getType() {
		return type;
	}
	public void setType(String type) {
		this.type = type;
	}
	public RegistroServicio getRegistroServicio() {
		return registroServicio;
	}
	public void setRegistroServicio(RegistroServicio registroServicio) {
		this.registroServicio = registroServicio;
	}
	String type;
	RegistroServicio registroServicio;
	

}
