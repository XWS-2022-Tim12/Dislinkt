package com.Dislinkt.Dislinkt.Repository;

import com.Dislinkt.Dislinkt.Model.Company;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface CompanyRepository extends JpaRepository<Company, String> {
    @Query("Select c from Company c where c.name = ?1")
    Company getByName(String name);
}
