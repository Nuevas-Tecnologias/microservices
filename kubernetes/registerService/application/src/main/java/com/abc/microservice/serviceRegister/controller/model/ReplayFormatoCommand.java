package com.abc.microservice.serviceRegister.controller.model;

import java.util.List;

import lombok.Data;

@Data
public class ReplayFormatoCommand {
	public String getType() {
		return type;
	}
	public void setType(String type) {
		this.type = type;
	}
	public int getCorrelationId() {
		return correlationId;
	}
	public void setCorrelationId(int correlationId) {
		this.correlationId = correlationId;
	}
	public List<Integer> getTechServiceId() {
		return techServiceId;
	}
	public void setTechServiceId(List<Integer> techServiceId) {
		this.techServiceId = techServiceId;
	}
	String type;
	int correlationId;
	List<Integer> techServiceId;
		
}
