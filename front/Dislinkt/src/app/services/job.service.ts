import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Job } from '../model/job';

@Injectable({
  providedIn: 'root'
})
export class JobService {
  private getRecomendedJobsUrl: string;
  private addNewJobUrl: string;

  constructor(private http: HttpClient) { 
    this.getRecomendedJobsUrl = 'http://localhost:8000/user/jobDislinktSearch';
    this.addNewJobUrl = 'http://localhost:8000/user/jobDislinkt';
  }

  public getRecomendedJobs(): Observable<Array<Job>> {
    let headers = new HttpHeaders();
    headers.append('Content-Type', 'application/json');
    return this.http.get<Array<Job>>(this.getRecomendedJobsUrl, {headers: headers, withCredentials: true});
  }
  public addNewJob(job: Job): Observable<string> {
    return this.http.post<string>(this.addNewJobUrl, job, {withCredentials: true});
  }
}
