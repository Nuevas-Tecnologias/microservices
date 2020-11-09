package com.abc.microservice.serviceRegister.service;

import java.util.List;

import com.abc.microservice.serviceRegister.repository.entity.AutoparteIntervenida;


public interface IAutoparteIntervenidaService {
	public AutoparteIntervenida saveAutoparteIntervenidaService(AutoparteIntervenida autoparteIntervenida);
	public AutoparteIntervenida getAutoparteIntervenidaById(int id);
	public List<AutoparteIntervenida> getAllAutoparteIntervenida();
	public String delateAutoparteIntervenida (int id);
	public AutoparteIntervenida updateAutoparteIntervenida (AutoparteIntervenida autoparteIntervenida);

}
