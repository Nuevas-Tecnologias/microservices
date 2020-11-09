package com.abc.microservice.serviceRegister.service.impl;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.abc.microservice.serviceRegister.repository.RegistroServicioRepository;
import com.abc.microservice.serviceRegister.repository.entity.AutoparteIntervenida;
import com.abc.microservice.serviceRegister.repository.entity.RegistroServicio;
import com.abc.microservice.serviceRegister.service.IRegistroServicioService;

@Service
public class RegistroServicioService implements IRegistroServicioService{
    @Autowired
    RegistroServicioRepository registroServiciorepository;
	
	@Override
	public RegistroServicio saveRegistroServicio(RegistroServicio registroServicio) {
		
		return registroServiciorepository.save(registroServicio);
	}

	@Override
	public List<RegistroServicio> saveRegistrosServivicios(List<RegistroServicio> registrosServicios) {
		
		return registroServiciorepository.saveAll(registrosServicios);
	}

	@Override
	public RegistroServicio getRegistroServicioById(int id) {
		
		return registroServiciorepository.findById(id).orElse(null);
	}

	@Override
	public List<RegistroServicio> getAllRegistroServicio() {
		
		return registroServiciorepository.findAll();
	}

	@Override
	public String delateRegistroServicio(int id) {
		registroServiciorepository.deleteById(id);
		return "registro eliminado: " + id;
	}

	@Override
	public RegistroServicio updateRegistroServicio(RegistroServicio registroServicio) {
		// TODO Auto-generated method stub
		RegistroServicio existingRegistroServicio = registroServiciorepository.findById(registroServicio.getId()).orElse(null);
		existingRegistroServicio.setIdOrdenServicio(registroServicio.getIdOrdenServicio());
		existingRegistroServicio.setIdCentroServicio(registroServicio.getIdCentroServicio());
		existingRegistroServicio.setNombreServicio(registroServicio.getNombreServicio());
		existingRegistroServicio.setEstado(registroServicio.getEstado());
		existingRegistroServicio.setFechaInicio(registroServicio.getFechaInicio());
		existingRegistroServicio.setFechaFin(registroServicio.getFechaFin());
		existingRegistroServicio.setNombreTecnico(registroServicio.getNombreTecnico());
		return registroServiciorepository.save(existingRegistroServicio);
	}

	@Override
	public List<RegistroServicio> findRegistroServicioByIdOrdenServicio(int idOrdenServicio) {
		
		return registroServiciorepository.findByIdOrdenServicio(idOrdenServicio);
	}

	@Override
	public RegistroServicio addAutoparteIntervenidaToResgistroServicio(int idRegistroServicio,
			AutoparteIntervenida autoparteIntervenida) {
		RegistroServicio existingRegistroServicio = registroServiciorepository.findById(idRegistroServicio).orElse(null);
		existingRegistroServicio.getAutopartesIntervenidas().add(autoparteIntervenida);
		return registroServiciorepository.save(existingRegistroServicio);
	}

}
