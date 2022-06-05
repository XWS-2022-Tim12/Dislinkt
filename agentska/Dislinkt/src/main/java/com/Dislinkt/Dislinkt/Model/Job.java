package com.Dislinkt.Dislinkt.Model;


import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateTimeDeserializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateTimeSerializer;

import java.time.LocalDateTime;
import java.util.ArrayList;

public class Job {
    private String id;
    private String requirements;
    private String position;
    private String userId;
    private String description;
    @JsonSerialize(using = LocalDateTimeSerializer.class)
    @JsonDeserialize(using = LocalDateTimeDeserializer.class)
    private LocalDateTime creationDay;
    private ArrayList<String> comments;
    private ArrayList<Integer> juniorSalary;
    private ArrayList<Integer> mediorSalary;
    private ArrayList<String> hrInterviews;
    private ArrayList<String> tehnicalInterviews;

    public Job() {}

    public Job(String id, String requirements, String position, String userId, String description, LocalDateTime creationDay, ArrayList<String> comments, ArrayList<Integer> juniorSalary, ArrayList<Integer> mediorSalary, ArrayList<String> hrInterviews, ArrayList<String> tehnicalInterviews) {
        this.id = id;
        this.requirements = requirements;
        this.position = position;
        this.userId = userId;
        this.description = description;
        this.creationDay = creationDay;
        this.comments = comments;
        this.juniorSalary = juniorSalary;
        this.mediorSalary = mediorSalary;
        this.hrInterviews = hrInterviews;
        this.tehnicalInterviews = tehnicalInterviews;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getRequirements() {
        return requirements;
    }

    public void setRequirements(String requirements) {
        this.requirements = requirements;
    }

    public String getPosition() {
        return position;
    }

    public void setPosition(String position) {
        this.position = position;
    }

    public String getUserId() {
        return userId;
    }

    public void setUserId(String userId) {
        this.userId = userId;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public LocalDateTime getCreationDay() {
        return creationDay;
    }

    public void setCreationDay(LocalDateTime creationDay) {
        this.creationDay = creationDay;
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
