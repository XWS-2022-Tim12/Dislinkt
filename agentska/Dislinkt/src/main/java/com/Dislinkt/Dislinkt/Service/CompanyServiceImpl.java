package com.Dislinkt.Dislinkt.Service;

import com.Dislinkt.Dislinkt.Model.Company;
import com.Dislinkt.Dislinkt.Model.User;
import com.Dislinkt.Dislinkt.Repository.CompanyRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.xml.stream.events.Comment;
import java.util.List;

@Service
public class CompanyServiceImpl implements CompanyService {
    @Autowired
    private CompanyRepository companyRepository;

    @Override
    public List<Company> getAll() {
        return companyRepository.findAll();
    }

    @Override
    public Company getById(String id) {
        return companyRepository.findById(id)
                .orElseThrow(() -> new IllegalArgumentException("User with id: '" + id + "' doesn't exist!"));
    }

    @Override
    public Company getByName(String name) {
        return companyRepository.getByName(name);
    }

    @Override
    public void registerCompany(Company company) {
        companyRepository.save(company);
    }

    @Override
    public void acceptCompanyRegistrationRequest(String name) {
        Company company = getByName(name);

        if (company != null && !company.isApproved()) {
            company.setApproved(true);
            company.getOwner().setRole(User.Role.agent_owner);

            companyRepository.save(company);
        } else {
            throw new IllegalArgumentException("Company already approved!");
        }
    }

    @Override
    public void changeCompanyDescription(Company company) {
        Company companyFromDatabase = getByName(company.getName());

        if (companyFromDatabase != null && companyFromDatabase.getOwner().getRole().equals(User.Role.agent_owner)) {
            companyFromDatabase.setDescription(company.getDescription());

            companyRepository.save(companyFromDatabase);
        } else {
            throw new IllegalArgumentException("Company doesn't exist or you are not owner!");
        }
    }

    public Boolean addCommentToCompany(Company company) {
        Company foundCompany = companyRepository.getByName(company.getName());

        if(foundCompany != null) {
            companyRepository.save(company);
            return true;
        } else {
            return false;
        }
    }
}
