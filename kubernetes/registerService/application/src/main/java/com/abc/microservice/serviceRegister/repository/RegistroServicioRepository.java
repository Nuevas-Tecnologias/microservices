package com.abc.microservice.serviceRegister.repository;

import java.util.List;

import org.springframework.data.jpa.repository.JpaRepository;

import com.abc.microservice.serviceRegister.repository.entity.RegistroServicio;


public interface RegistroServicioRepository extends JpaRepository <RegistroServicio,Integer>{

	List<RegistroServicio> findByIdOrdenServicio (int idCentroServicio);
}
