package com.Dislinkt.Dislinkt.Model;


import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateTimeDeserializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateTimeSerializer;

import java.time.LocalDateTime;

public class Job {
    private String id;
    private String requirements;
    private String position;
    private String userId;
    private String description;
    @JsonSerialize(using = LocalDateTimeSerializer.class)
    @JsonDeserialize(using = LocalDateTimeDeserializer.class)
    private LocalDateTime creationDay;

    public Job() {}

    public Job(String id, String requirements, String position, String userId, String description, LocalDateTime creationDay) {
        this.id = id;
        this.requirements = requirements;
        this.position = position;
        this.userId = userId;
        this.description = description;
        this.creationDay = creationDay;
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
}
