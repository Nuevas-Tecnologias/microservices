package com.abc.microservice.serviceRegister.service.impl;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;

import com.abc.microservice.serviceRegister.repository.AutoparteIntervenidaRepository;
import com.abc.microservice.serviceRegister.repository.entity.AutoparteIntervenida;
import com.abc.microservice.serviceRegister.service.IAutoparteIntervenidaService;

public class AutoparteIntervenidaService implements IAutoparteIntervenidaService{

	@Autowired
	AutoparteIntervenidaRepository autoparteIntervenidaRepository;
	
	@Override
	public AutoparteIntervenida saveAutoparteIntervenidaService(AutoparteIntervenida autoparteIntervenida) {
		
		return autoparteIntervenidaRepository.save(autoparteIntervenida);
	}

	@Override
	public AutoparteIntervenida getAutoparteIntervenidaById(int id) {
		
		return autoparteIntervenidaRepository.findById(id).orElse(null);
	}

	@Override
	public List<AutoparteIntervenida> getAllAutoparteIntervenida() {
		
		return autoparteIntervenidaRepository.findAll() ;
	}

	@Override
	public String delateAutoparteIntervenida(int id) {
		autoparteIntervenidaRepository.deleteById(id);
		return "Autoparte eliminada: " + id;
	}

	@Override
	public AutoparteIntervenida updateAutoparteIntervenida(AutoparteIntervenida autoparteIntervenida) {
		AutoparteIntervenida existingAutoparteIntervenida = autoparteIntervenidaRepository.findById(autoparteIntervenida.getId()).orElse(null);
		existingAutoparteIntervenida.setIdUnicoAutoparte(autoparteIntervenida.getIdUnicoAutoparte());
		existingAutoparteIntervenida.setNombre(autoparteIntervenida.getNombre());
		existingAutoparteIntervenida.setOperacion(autoparteIntervenida.getOperacion());
		existingAutoparteIntervenida.setResgistroServicio(autoparteIntervenida.getResgistroServicio());
		return autoparteIntervenidaRepository.save(existingAutoparteIntervenida);
	}

}
