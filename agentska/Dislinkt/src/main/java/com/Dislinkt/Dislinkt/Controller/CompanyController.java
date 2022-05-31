package com.Dislinkt.Dislinkt.Controller;

import com.Dislinkt.Dislinkt.Model.Company;
import com.Dislinkt.Dislinkt.Model.User;
import com.Dislinkt.Dislinkt.Service.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
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

    @PutMapping("/company/edit")
    public ResponseEntity<Boolean> addCommentToCompany(@RequestBody Company company) {
        if (companyService.addCommentToCompany(company)) {
            return new ResponseEntity<Boolean>(true, HttpStatus.OK);
        } else {
            return new ResponseEntity<Boolean>(false, HttpStatus.NOT_ACCEPTABLE);
        }
    }

    @GetMapping("/company/companyByName")
    public ResponseEntity<Company> getCompanyByName(@RequestParam("name") String name) {
        return new ResponseEntity<Company>(companyService.getByName(name), HttpStatus.OK);
    }
}
