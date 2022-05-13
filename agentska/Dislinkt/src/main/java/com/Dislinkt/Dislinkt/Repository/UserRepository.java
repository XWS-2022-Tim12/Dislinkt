package com.Dislinkt.Dislinkt.Repository;

import com.Dislinkt.Dislinkt.Model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    @Query("Select u from User u where u.email = ?1")
    public User findByEmail(String email);
    @Query("Select u from User u where u.username = ?1")
    public User findByUsername(String username);
    @Query("Select u from User u where u.username = ?1 and u.password = ?2")
    public User findByUsernameAndPassword(String email,String password);
}
