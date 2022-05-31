package com.Dislinkt.Dislinkt.Model;

import com.sun.istack.NotNull;
import org.hibernate.annotations.GenericGenerator;

import javax.persistence.*;
import java.util.ArrayList;

@Entity
@Table(name = "companies")
public class Company {
    @Id
    @GeneratedValue(generator = "uuid")
    @GenericGenerator(name = "uuid", strategy = "uuid2")
    private String id;
    @NotNull
    private String name;
    private String address;
    private String phoneNumber;
    private String description;
    @NotNull
    @ManyToOne
    private User owner;
    private boolean approved;
    private ArrayList<String> comments;
    private ArrayList<Integer> juniorSalary;
    private ArrayList<Integer> mediorSalary;
    private ArrayList<String> hrInterviews;
    private ArrayList<String> tehnicalInterviews;

    public Company() {}

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public String getPhoneNumber() {
        return phoneNumber;
    }

    public void setPhoneNumber(String phoneNumber) {
        this.phoneNumber = phoneNumber;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public User getOwner() {
        return owner;
    }

    public void setOwner(User owner) {
        this.owner = owner;
    }

    public boolean isApproved() {
        return approved;
    }

    public void setApproved(boolean approved) {
        this.approved = approved;
    }

    public ArrayList<String> getComments() {
        return comments;
    }

    public void setComments(ArrayList<String> comments) {
        this.comments = comments;
    }

    public ArrayList<Integer> getJuniorSalary() {
        return juniorSalary;
    }

    public void setJuniorSalary(ArrayList<Integer> juniorSalary) {
        this.juniorSalary = juniorSalary;
    }

    public ArrayList<Integer> getMediorSalary() {
        return mediorSalary;
    }

    public void setMediorSalary(ArrayList<Integer> mediorSalary) {
        this.mediorSalary = mediorSalary;
    }

    public ArrayList<String> getHrInterviews() {
        return hrInterviews;
    }

    public void setHrInterviews(ArrayList<String> hrInterviews) {
        this.hrInterviews = hrInterviews;
    }

    public ArrayList<String> getTehnicalInterviews() {
        return tehnicalInterviews;
    }

    public void setTehnicalInterviews(ArrayList<String> tehnicalInterviews) {
        this.tehnicalInterviews = tehnicalInterviews;
    }
}
