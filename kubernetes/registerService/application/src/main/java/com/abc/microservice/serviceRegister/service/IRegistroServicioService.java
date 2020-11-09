package com.abc.microservice.serviceRegister.service;

import java.util.List;

import com.abc.microservice.serviceRegister.repository.entity.AutoparteIntervenida;
import com.abc.microservice.serviceRegister.repository.entity.RegistroServicio;

public interface IRegistroServicioService {

	public RegistroServicio saveRegistroServicio(RegistroServicio registroServicio);
	public List<RegistroServicio> saveRegistrosServivicios(List<RegistroServicio> registrosServicios);
	public RegistroServicio getRegistroServicioById(int id);
	public List<RegistroServicio> getAllRegistroServicio();
	public String delateRegistroServicio (int id);
	public RegistroServicio updateRegistroServicio (RegistroServicio registroServicio);
	public List<RegistroServicio> findRegistroServicioByIdOrdenServicio(int idOrdenServicio);
	public RegistroServicio addAutoparteIntervenidaToResgistroServicio(int idRegistroServicio, AutoparteIntervenida autoparteIntervenida);
	
	
}
