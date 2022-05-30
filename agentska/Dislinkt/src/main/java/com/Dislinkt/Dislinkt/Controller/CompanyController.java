package com.Dislinkt.Dislinkt.Controller;

import com.Dislinkt.Dislinkt.Model.Company;
import com.Dislinkt.Dislinkt.Service.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@CrossOrigin(origins = "http://localhost:4201")
public class CompanyController {
    @Autowired
    private CompanyService companyService;

    @GetMapping("/company/companies")
    public ResponseEntity<List<Company>> getAllCompanies() {
        return new ResponseEntity<>(companyService.getAll(), HttpStatus.OK);
    }

    @GetMapping("/company/{id}")
    public ResponseEntity<Company> getById(@PathVariable String id) {
        return new ResponseEntity<>(companyService.getById(id), HttpStatus.OK);
    }
}
