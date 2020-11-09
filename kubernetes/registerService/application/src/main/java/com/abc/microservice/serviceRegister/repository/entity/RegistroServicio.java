package com.abc.microservice.serviceRegister.repository.entity;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.OneToMany;

import java.util.List;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
public class RegistroServicio {
	
		public String getPlacaVehiculo() {
		return placaVehiculo;
	}
	public void setPlacaVehiculo(String placaVehiculo) {
		this.placaVehiculo = placaVehiculo;
	}
		@Id
	    @GeneratedValue
	    private int id;
	    private int idOrdenServicio;
	    private int idCentroServicio;
	    private String estado;
	    private String nombreServicio;
	    private String nombreTecnico;
	    private String fechaInicio;
	    private String fechaFin;
	    private String placaVehiculo;
	    
	    @OneToMany(targetEntity = AutoparteIntervenida.class, cascade = CascadeType.ALL)
	    @JoinColumn(name="registro_servicio_id")
	    private List<AutoparteIntervenida> autopartesIntervenidas;
	    
		public Integer getId() {
			
			return id;
		}
	    public int getIdOrdenServicio() {
			return idOrdenServicio;
		}
		public void setIdOrdenServicio(int idOrdenServicio) {
			this.idOrdenServicio = idOrdenServicio;
		}
		public int getIdCentroServicio() {
			return idCentroServicio;
		}
		public void setIdCentroServicio(int idCentroServicio) {
			this.idCentroServicio = idCentroServicio;
		}
		public String getEstado() {
			return estado;
		}
		public void setEstado(String estado) {
			this.estado = estado;
		}
		public String getNombreServicio() {
			return nombreServicio;
		}
		public void setNombreServicio(String nombreServicio) {
			this.nombreServicio = nombreServicio;
		}
		public String getNombreTecnico() {
			return nombreTecnico;
		}
		public void setNombreTecnico(String nombreTecnico) {
			this.nombreTecnico = nombreTecnico;
		}
		public String getFechaInicio() {
			return fechaInicio;
		}
		public void setFechaInicio(String fechaInicio) {
			this.fechaInicio = fechaInicio;
		}
		public String getFechaFin() {
			return fechaFin;
		}
		public void setFechaFin(String fechaFin) {
			this.fechaFin = fechaFin;
		}
		public List<AutoparteIntervenida> getAutopartesIntervenidas() {
			return autopartesIntervenidas;
		}
		public void setAutopartesIntervenidas(List<AutoparteIntervenida> autopartesIntervenidas) {
			this.autopartesIntervenidas = autopartesIntervenidas;
		}
		public void setId(int id) {
			this.id = id;
		}
	    	    
	    

}
