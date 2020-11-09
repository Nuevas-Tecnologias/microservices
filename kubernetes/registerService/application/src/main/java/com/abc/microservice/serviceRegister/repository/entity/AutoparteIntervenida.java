package com.abc.microservice.serviceRegister.repository.entity;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;

import com.fasterxml.jackson.annotation.JsonIgnore;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
public class AutoparteIntervenida {

	

	@Id
    @GeneratedValue
    private int id;
	private String idUnicoAutoparte;
	private String nombre;
	private String operacion;
	
    @ManyToOne
    @JsonIgnore	
    private RegistroServicio registroServicio;
	
	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	public String getIdUnicoAutoparte() {
		return idUnicoAutoparte;
	}

	public void setIdUnicoAutoparte(String idUnicoAutoparte) {
		this.idUnicoAutoparte = idUnicoAutoparte;
	}

	public String getNombre() {
		return nombre;
	}

	public void setNombre(String nombre) {
		this.nombre = nombre;
	}

	public String getOperacion() {
		return operacion;
	}

	public void setOperacion(String operacion) {
		this.operacion = operacion;
	}

	public RegistroServicio getResgistroServicio() {
		return registroServicio;
	}

	public void setResgistroServicio(RegistroServicio resgistroServicio) {
		this.registroServicio = resgistroServicio;
	}
}
