import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Job } from '../model/job';

@Injectable({
  providedIn: 'root'
})
export class JobService {
  private getAllJobsUrl: string;
  private getJobByIdUrl: string;
  private editJobUrl: string;

  constructor(private http: HttpClient) { 
    this.getAllJobsUrl = 'http://localhost:8080/job/jobs';
    this.getJobByIdUrl = 'http://localhost:8080/job';
    this.editJobUrl = 'http://localhost:8080/job/editJob';
  }

  public getAllJobs(): Observable<Array<Job>> {
    let headers = new HttpHeaders();
    headers.append('Content-Type', 'application/json');
    
    return this.http.get<Array<Job>>(this.getAllJobsUrl, {headers: headers, withCredentials: true});
  }

  public getJobById(id: string): Observable<Job> {
    let headers = new HttpHeaders();
    headers.append('Content-Type', 'application/json');
    let params = new HttpParams().set("id",id);

    return this.http.get<Job>(this.getJobByIdUrl, {headers: headers, params: params, withCredentials: true});
  }

  public editJob(job: Job): Observable<Boolean> {
    return this.http.put<boolean>(this.editJobUrl, job, {withCredentials: true});
  }
}
