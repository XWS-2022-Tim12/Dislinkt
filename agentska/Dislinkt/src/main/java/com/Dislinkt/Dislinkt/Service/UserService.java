package com.Dislinkt.Dislinkt.Service;


import com.Dislinkt.Dislinkt.Model.Company;
import com.Dislinkt.Dislinkt.Model.User;

import java.util.List;

public interface UserService {
    Boolean register(User user);
    Boolean login(User user);
    List<User> getAll();
    String getId(User user);
    User.Role getRole(User user);
    Boolean checkCompanyForRegistration(Company company, String loggedUserUsername);
    void registerCompany(Company company);
    void acceptCompanyRegistrationRequest(String name);
    User findByEmail(String email);
    User findByUsername(String username);
    Boolean checkCompanyName(String name);
    boolean checkCompanyForChange(Company company, String loggedUserUsername);
    void changeCompanyDescription(Company company);
}
