package com.Dislinkt.Dislinkt.Service;


import com.Dislinkt.Dislinkt.Model.User;

import java.util.List;

public interface UserService {
    public Boolean register(User user);
    public Boolean login(User user);
    public List<User> getAll();
    public String getId(User user);
    public User.Role getRole(User user);
}
