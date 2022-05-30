package com.Dislinkt.Dislinkt.Controller;

import com.Dislinkt.Dislinkt.Model.Company;
import com.Dislinkt.Dislinkt.Model.Job;
import com.Dislinkt.Dislinkt.Model.User;
import com.Dislinkt.Dislinkt.Service.UserService;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ArrayNode;
import com.fasterxml.jackson.databind.util.JSONPObject;
import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.net.URI;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;


@RestController
@CrossOrigin(origins = "http://localhost:4201")
public class UserController {

    @Autowired
    private UserService userService;

    @PostMapping("/user")
    public ResponseEntity<Void> register(@RequestBody User user) {
        if(userService.register(user)){
            return new ResponseEntity<Void>(HttpStatus.OK);
        }
        else{
            return new ResponseEntity<Void>(HttpStatus.NOT_ACCEPTABLE);
        }
    }

    @PostMapping("/user/login")
    public ResponseEntity<Void> login(@RequestBody User user, HttpServletRequest request, HttpServletResponse response) {
        if (userService.login(user)){
            User.Role userRole = userService.getRole(user);
            String username = user.getUsername();

            if (request.getSession().getAttribute("role") == null){
                request.getSession(true).setAttribute("role", userRole);
            }
            else {
                request.getSession(false).setAttribute("role", userRole);
            }

            if (request.getSession().getAttribute("username") == null){
                request.getSession(true).setAttribute("username", username);
            }
            else {
                request.getSession(false).setAttribute("username", username);
            }

            if(userService.getRole(user).equals(User.Role.agent_owner)) {
                String url = "http://localhost:8000/user/loginOwner";

                HttpHeaders headers = new HttpHeaders();
                RestTemplate restTemplate = new RestTemplate();
                headers.setContentType(MediaType.APPLICATION_JSON);
                JSONObject json = new JSONObject();
                json.put("id", userService.getId(user));
                HttpEntity requestt = new HttpEntity(headers);
                ResponseEntity<String> responsee = restTemplate.postForEntity(url,json, String.class, 1);
                if (responsee.getStatusCode() == HttpStatus.OK) {
                    request.getSession().setAttribute("ownerSessionId", responsee.getBody());
                    Cookie cookie = new Cookie("ownerSessionId", responsee.getBody());
                    cookie.setPath("/job");
                    response.addCookie(cookie);
                }
            }
            return new ResponseEntity<Void>(HttpStatus.OK);
        }
        else {
            return new ResponseEntity<Void>(HttpStatus.NOT_FOUND);
        }
    }

    @GetMapping("/user/users")
    public ResponseEntity<List<User>> getAll() {
            return new ResponseEntity<List<User>>(userService.getAll(), HttpStatus.OK);
    }

    @PostMapping("/job")
    public ResponseEntity<String> AddNewJob(@RequestBody Job job, HttpServletRequest request) {
        if(request.getSession().getAttribute("role").equals(User.Role.agent_owner)){
            String cookieValue = "";
            Cookie[] cookies = request.getCookies();
            for (Cookie c : cookies){
                if (c.getName().equals("ownerSessionId")){
                    cookieValue = c.getValue();
                }
            }
            String url = "http://localhost:8000/job";
            HttpHeaders headers = new HttpHeaders();
            headers.add("Cookie", "ownerSessionId="+cookieValue+"; Path=/job");
            RestTemplate restTemplate = new RestTemplate();
            JSONObject json = new JSONObject();
            json.put("id", job.getId());
            json.put("userId", job.getUserId());
            json.put("position", job.getPosition());
            json.put("description", job.getDescription());
            json.put("requirements", job.getRequirements());
            HttpEntity<String> requestt = new HttpEntity<String>(json.toString(), headers);
            ResponseEntity<String> responsee = restTemplate.postForEntity(url, requestt, String.class);
            return new ResponseEntity<String>(responsee.getBody(),HttpStatus.OK);
        }
        return new ResponseEntity<String>("",HttpStatus.NOT_ACCEPTABLE);
    }
    @GetMapping("/job")
    public ResponseEntity<Job> FindJobById(@RequestParam("id") String content, HttpServletRequest request) throws JsonProcessingException {
        if(request.getSession().getAttribute("role") != null){
            String url = "http://localhost:8000/job" + "/" + content;
            HttpHeaders headers = new HttpHeaders();
            RestTemplate restTemplate = new RestTemplate();
            HttpEntity requestt = new HttpEntity(headers);
            String jsonStr = restTemplate.exchange(url, HttpMethod.GET, requestt, String.class, 1).getBody();
            jsonStr = jsonStr.split("\\{")[2];
            jsonStr = "{" + jsonStr;
            ObjectMapper mapper = new ObjectMapper();
            JsonNode rootNode = mapper.readTree(jsonStr);
            if (rootNode instanceof JsonNode) {
                // Read the json as a single object:
                Job job = mapper.readValue(rootNode.toString(), Job.class);
                return new ResponseEntity<Job>(job,HttpStatus.OK);
            }
        }
        return new ResponseEntity<Job>(HttpStatus.NOT_ACCEPTABLE);
    }
    @GetMapping("/job/searchByUser")
    public ResponseEntity<Job[]> FindJobByUser(@RequestParam("userId") String content, HttpServletRequest request) throws JsonProcessingException {
        if(request.getSession().getAttribute("role") != null){
            String url = "http://localhost:8000/job/searchByUser" + "/" + content;
            HttpHeaders headers = new HttpHeaders();
            RestTemplate restTemplate = new RestTemplate();
            HttpEntity requestt = new HttpEntity(headers);
            String jsonStr = restTemplate.exchange(url, HttpMethod.GET, requestt, String.class, 1).getBody();
            jsonStr = jsonStr.split("\\[")[1];
            jsonStr = "[" + jsonStr;
            ObjectMapper mapper = new ObjectMapper();
            JsonNode rootNode = mapper.readTree(jsonStr);
            if (rootNode instanceof ArrayNode) {
                // Read the json as a list:
                Job[] jobs = mapper.readValue(rootNode.toString(), Job[].class);
                return new ResponseEntity<Job[]>(jobs,HttpStatus.OK);
            } else if (rootNode instanceof JsonNode) {
                // Read the json as a single object:
                Job job = mapper.readValue(rootNode.toString(), Job.class);
                List<Job> jobs = new ArrayList<Job>();
                jobs.add(job);
                return new ResponseEntity<Job[]>(jobs.toArray(new Job[0]),HttpStatus.OK);
            }
        }
        return new ResponseEntity<Job[]>(HttpStatus.NOT_ACCEPTABLE);
    }
    @GetMapping("/job/jobs")
    public ResponseEntity<Job[]> FindJobs(HttpServletRequest request) throws JsonProcessingException {
        if(request.getSession().getAttribute("role") != null){
            String url = "http://localhost:8000/job/jobs";
            HttpHeaders headers = new HttpHeaders();
            RestTemplate restTemplate = new RestTemplate();
            HttpEntity requestt = new HttpEntity(headers);
            String jsonStr = restTemplate.exchange(url, HttpMethod.GET, requestt, String.class, 1).getBody();
            jsonStr = jsonStr.split("\\[")[1];
            jsonStr = "[" + jsonStr;
            ObjectMapper mapper = new ObjectMapper();
            JsonNode rootNode = mapper.readTree(jsonStr);
            if (rootNode instanceof ArrayNode) {
                // Read the json as a list:
                Job[] jobs = mapper.readValue(rootNode.toString(), Job[].class);
                return new ResponseEntity<Job[]>(jobs,HttpStatus.OK);
            } else if (rootNode instanceof JsonNode) {
                // Read the json as a single object:
                Job job = mapper.readValue(rootNode.toString(), Job.class);
                List<Job> jobs = new ArrayList<Job>();
                jobs.add(job);
                return new ResponseEntity<Job[]>(jobs.toArray(new Job[0]),HttpStatus.OK);
            }
        }
        return new ResponseEntity<Job[]>(HttpStatus.NOT_ACCEPTABLE);
    }
    @GetMapping("/job/searchByDescription")
    public ResponseEntity<Job[]> FindJobByDesctiption(@RequestParam("description") String content, HttpServletRequest request) throws JsonProcessingException {
        if(request.getSession().getAttribute("role") != null){

            String url = "http://localhost:8000/job/searchByDescription"+ "/" + content;
            HttpHeaders headers = new HttpHeaders();
            RestTemplate restTemplate = new RestTemplate();
            HttpEntity requestt = new HttpEntity(headers);
            String jsonStr = restTemplate.exchange(url, HttpMethod.GET, requestt, String.class, 1).getBody();
            jsonStr = jsonStr.split("\\[")[1];
            jsonStr = "[" + jsonStr;
            ObjectMapper mapper = new ObjectMapper();
            JsonNode rootNode = mapper.readTree(jsonStr);
            if (rootNode instanceof ArrayNode) {
                // Read the json as a list:
                Job[] jobs = mapper.readValue(rootNode.toString(), Job[].class);
                return new ResponseEntity<Job[]>(jobs,HttpStatus.OK);
            } else if (rootNode instanceof JsonNode) {
                // Read the json as a single object:
                Job job = mapper.readValue(rootNode.toString(), Job.class);
                List<Job> jobs = new ArrayList<Job>();
                jobs.add(job);
                return new ResponseEntity<Job[]>(jobs.toArray(new Job[0]),HttpStatus.OK);
            }
        }
        return new ResponseEntity<Job[]>(HttpStatus.NOT_ACCEPTABLE);
    }
    @GetMapping("/job/searchByPosition")
    public ResponseEntity<Job[]> FindJobByPosition(@RequestParam("position") String content, HttpServletRequest request) throws JsonProcessingException {
        if(request.getSession().getAttribute("role") != null){
            String url = "http://localhost:8000/job/searchByPosition"+ "/" + content;
            HttpHeaders headers = new HttpHeaders();
            RestTemplate restTemplate = new RestTemplate();
            HttpEntity requestt = new HttpEntity(headers);
            String jsonStr = restTemplate.exchange(url, HttpMethod.GET, requestt, String.class, 1).getBody();
            jsonStr = jsonStr.split("\\[")[1];
            jsonStr = "[" + jsonStr;
            ObjectMapper mapper = new ObjectMapper();
            JsonNode rootNode = mapper.readTree(jsonStr);
            if (rootNode instanceof ArrayNode) {
                // Read the json as a list:
                Job[] jobs = mapper.readValue(rootNode.toString(), Job[].class);
                return new ResponseEntity<Job[]>(jobs,HttpStatus.OK);
            } else if (rootNode instanceof JsonNode) {
                // Read the json as a single object:
                Job job = mapper.readValue(rootNode.toString(), Job.class);
                List<Job> jobs = new ArrayList<Job>();
                jobs.add(job);
                return new ResponseEntity<Job[]>(jobs.toArray(new Job[0]),HttpStatus.OK);
            }
        }
        return new ResponseEntity<Job[]>(HttpStatus.NOT_ACCEPTABLE);
    }
    @GetMapping("/job/searchByRequirements")
    public ResponseEntity<Job[]> FindJobByRequirements(@RequestParam("requirements") String content, HttpServletRequest request) throws JsonProcessingException {
        if(request.getSession().getAttribute("role") != null){
            String url = "http://localhost:8000/job/searchByRequirements"+ "/" + content;
            HttpHeaders headers = new HttpHeaders();
            RestTemplate restTemplate = new RestTemplate();
            HttpEntity requestt = new HttpEntity(headers);
            String jsonStr = restTemplate.exchange(url, HttpMethod.GET, requestt, String.class, 1).getBody();
            jsonStr = jsonStr.split("\\[")[1];
            jsonStr = "[" + jsonStr;
            ObjectMapper mapper = new ObjectMapper();
            JsonNode rootNode = mapper.readTree(jsonStr);
            if (rootNode instanceof ArrayNode) {
                // Read the json as a list:
                Job[] jobs = mapper.readValue(rootNode.toString(), Job[].class);
                return new ResponseEntity<Job[]>(jobs,HttpStatus.OK);
            } else if (rootNode instanceof JsonNode) {
                // Read the json as a single object:
                Job job = mapper.readValue(rootNode.toString(), Job.class);
                List<Job> jobs = new ArrayList<Job>();
                jobs.add(job);
                return new ResponseEntity<Job[]>(jobs.toArray(new Job[0]),HttpStatus.OK);
            }
        }
        return new ResponseEntity<Job[]>(HttpStatus.NOT_ACCEPTABLE);
    }

    @PostMapping("/user/registerCompany")
    public ResponseEntity<Boolean> registerCompany(@RequestBody Company company, HttpServletRequest request) {
        if (request.getSession().getAttribute("role") != null && request.getSession().getAttribute("username") != null) {
            String loggedUserUsername = request.getSession().getAttribute("username").toString();

            if (userService.checkCompanyForRegistration(company, loggedUserUsername)) {
                userService.registerCompany(company);

                return new ResponseEntity<>(true, HttpStatus.OK);
            }

            return new ResponseEntity<>(false, HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(false, HttpStatus.NOT_ACCEPTABLE);
    }

    @PutMapping("/user/acceptCompanyRegistrationRequest")
    public ResponseEntity<Boolean> acceptCompanyRegistrationRequest(@RequestBody Company company, HttpServletRequest request) {
        if (request.getSession().getAttribute("role") != null && request.getSession().getAttribute("username") != null) {
            if (request.getSession().getAttribute("role").equals(User.Role.admin)) {
                if (company.getName() != null && userService.checkCompanyName(company.getName())) {
                    userService.acceptCompanyRegistrationRequest(company.getName());

                    return new ResponseEntity<>(true, HttpStatus.OK);
                }
            }

            return new ResponseEntity<>(false, HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(false, HttpStatus.NOT_ACCEPTABLE);
    }

    @PutMapping("/user/changeCompanyDescription")
    public ResponseEntity<Boolean> changeCompanyDescription(@RequestBody Company company, HttpServletRequest request) {
        if (request.getSession().getAttribute("role") != null && request.getSession().getAttribute("username") != null) {
            if (request.getSession().getAttribute("role").equals(User.Role.agent_owner)) {
                String loggedUserUsername = request.getSession().getAttribute("username").toString();

                if (userService.checkCompanyForChange(company, loggedUserUsername)) {
                    userService.changeCompanyDescription(company);

                    return new ResponseEntity<>(true, HttpStatus.OK);
                }
            }

            return new ResponseEntity<>(false, HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<>(false, HttpStatus.NOT_ACCEPTABLE);
    }
}
