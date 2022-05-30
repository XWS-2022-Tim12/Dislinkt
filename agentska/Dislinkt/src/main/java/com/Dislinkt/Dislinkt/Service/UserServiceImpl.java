package com.Dislinkt.Dislinkt.Service;


import com.Dislinkt.Dislinkt.Model.Company;
import com.Dislinkt.Dislinkt.Model.User;
import com.Dislinkt.Dislinkt.Repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;


@Service
public class UserServiceImpl implements UserService {
    @Autowired
    private UserRepository userRepository;
    @Autowired
    private CompanyService companyService;

    public Boolean register(User user) {
        User foundUser = userRepository.findByEmail(user.getEmail());
        if(foundUser != null){
            return false;
        }
        foundUser = userRepository.findByUsername(user.getUsername());
        if(foundUser != null){
            return false;
        }

        user.setRole(User.Role.agent_user);
        user.setPassword(user.getPassword());
        userRepository.save(user);
        return true;
    }

    public Boolean login(User user) {
        User foundUser = userRepository.findByUsernameAndPassword(user.getUsername(),user.getPassword());
        if (foundUser != null){
            return true;
        }
        else{
            return false;
        }
    }

    public List<User> getAll() {
        return userRepository.findAll();
    }

    public String getId(User user) {
        User foundUser = userRepository.findByUsernameAndPassword(user.getUsername(),user.getPassword());
        if (foundUser != null){
            return foundUser.getId();
        }
        else{
            return "";
        }
    }

    public User.Role getRole(User user) {
        User foundUser = userRepository.findByUsernameAndPassword(user.getUsername(),user.getPassword());
        if (foundUser != null){
            return foundUser.getRole();
        }
        else{
            return null;
        }
    }

    @Override
    public Boolean checkCompanyForRegistration(Company company, String loggedUserUsername) {
        if (company.getName() == null || company.getName().equals("")) {
            throw new IllegalArgumentException("Company name can't be null or empty!");
        } else if (companyService.getByName(company.getName()) != null) {
            throw new IllegalArgumentException("Company with given name already exists!");
        } else if (company.getOwner() == null) {
            throw new IllegalArgumentException("Company owner can't be null!");
        } else if (company.getOwner().getUsername() == null || company.getOwner().getUsername().equals("")) {
            throw new IllegalArgumentException("Company owner username can't be null or empty!");
        } else if (!company.getOwner().getUsername().equals(loggedUserUsername)) {
            throw new IllegalArgumentException("Company owner from request is not logged in!");
        }

        return true;
    }

    @Override
    public void registerCompany(Company company) {
        User owner = findByUsername(company.getOwner().getEmail());
        company.setOwner(owner);
        company.setApproved(false);
        companyService.registerCompany(company);
    }

    @Override
    public void acceptCompanyRegistrationRequest(String name) {
        companyService.acceptCompanyRegistrationRequest(name);
    }

    @Override
    public User findByEmail(String email) {
        return userRepository.findByEmail(email);
    }

    @Override
    public User findByUsername(String username) {
        return userRepository.findByUsername(username);
    }

    @Override
    public Boolean checkCompanyName(String name) {
        return companyService.getByName(name) != null;
    }

    @Override
    public boolean checkCompanyForChange(Company company, String loggedUserUsername) {
        if (company.getName() == null || company.getName().equals("")) {
            throw new IllegalArgumentException("Company name can't be null or empty!");
        } else if (companyService.getByName(company.getName()) == null) {
            throw new IllegalArgumentException("Company with given name doesn't exists!");
        } else if (company.getDescription() == null || company.getDescription().equals("")) {
            throw new IllegalArgumentException("Company description can't be null or empty!");
        } else if (company.getOwner() == null) {
            throw new IllegalArgumentException("Company owner can't be null!");
        } else if (company.getOwner().getUsername() == null || company.getOwner().getUsername().equals("")) {
            throw new IllegalArgumentException("Company owner username can't be null or empty!");
        } else if (!company.getOwner().getUsername().equals(loggedUserUsername)) {
            throw new IllegalArgumentException("Company owner from request is not logged in!");
        }

        return true;
    }

    @Override
    public void changeCompanyDescription(Company company) {
        companyService.changeCompanyDescription(company);
    }
}
