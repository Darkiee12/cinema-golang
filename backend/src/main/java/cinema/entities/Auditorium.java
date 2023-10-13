package cinema.entities;

import lombok.Getter;
import lombok.Setter;

import java.util.Set;

import com.fasterxml.jackson.annotation.JsonIgnore;

import jakarta.persistence.*;

@Entity
@Table(name = "auditoriums")
@Getter
@Setter
public class Auditorium {
    @JsonIgnore
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "auditorium_id")
    private Long auditoriumId;

    @Column(name = "name", length = 255)
    private String name;

    @Column(name = "seats")
    private Integer seats;

    @ManyToOne
    @JoinColumn(name = "cinema_id", referencedColumnName = "cinema_id")
    private Cinema cinema;

    @JsonIgnore
    @OneToMany
    @JoinColumn(name = "auditorium_id", referencedColumnName = "auditorium_id")
    private Set<AuditoriumSeat> auditoriumSeats;

    @JsonIgnore
    @OneToMany
    @JoinColumn(name = "auditorium_id", referencedColumnName = "auditorium_id")
    private Set<Show> shows;
}
