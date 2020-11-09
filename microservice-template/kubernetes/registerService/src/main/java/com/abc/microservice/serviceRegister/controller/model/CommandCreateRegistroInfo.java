package com.abc.microservice.serviceRegister.controller.model;

import java.util.List;

import com.abc.microservice.serviceRegister.repository.entity.*;
public class CommandCreateRegistroInfo {
	public CommandCreateRegistroInfo( int correlationId, List<RegistroServicio> registrosServicio) {
		super();
		
		this.correlationId = correlationId;
		this.registrosServicio = registrosServicio;
	}
	
	public int getCorrelationId() {
		return correlationId;
	}
	public void setCorrelationId(int correlationId) {
		this.correlationId = correlationId;
	}
	public List<RegistroServicio> getRegistrosServicio() {
		return registrosServicio;
	}
	public void setRegistrosServicio(List<RegistroServicio> registrosServicio) {
		this.registrosServicio = registrosServicio;
	}
	
private int correlationId;
 private List<RegistroServicio> registrosServicio;
 
}
