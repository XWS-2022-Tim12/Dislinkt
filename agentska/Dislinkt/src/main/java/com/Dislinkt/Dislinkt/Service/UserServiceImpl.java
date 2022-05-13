package com.Dislinkt.Dislinkt.Service;


import com.Dislinkt.Dislinkt.Model.User;
import com.Dislinkt.Dislinkt.Repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;


@Service
public class UserServiceImpl implements UserService {
    @Autowired
    private UserRepository userRepository;

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
}
