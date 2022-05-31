package com.Dislinkt.Dislinkt.Service;

import com.Dislinkt.Dislinkt.Model.Company;

import java.util.List;

public interface CompanyService {
    List<Company> getAll();
    Company getById(String id);
    Company getByName(String name);
    void registerCompany(Company company);
    void acceptCompanyRegistrationRequest(String name);
    void changeCompanyDescription(Company company);
    public Boolean addCommentToCompany(Company company);
}
