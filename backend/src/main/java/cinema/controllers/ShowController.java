package cinema.controllers;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import cinema.services.ShowService;
import cinema.entities.Show;
import cinema.pojos.ShowRequest;

@RestController
public class ShowController{
    @Autowired ShowService showService;

    @GetMapping("/shows")
    public List<Show> getShows(){
        return showService.getShows();
    }

    @PostMapping("/shows")
    public Show addShow(@RequestBody ShowRequest showRequest){
        return showService.saveShow(showRequest);
    }
}
