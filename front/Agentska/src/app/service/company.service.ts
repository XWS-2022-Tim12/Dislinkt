import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Company } from '../model/company';

@Injectable({
  providedIn: 'root'
})
export class CompanyService {
  private getAllCompaniesUrl: string;
  private EditCompanyUrl: string;
  private getCompanyByNameUrl: string;

  constructor(private http: HttpClient) { 
    this.getAllCompaniesUrl = 'http://localhost:8080/company/companies';
    this.EditCompanyUrl = 'http://localhost:8080/company/edit';
    this.getCompanyByNameUrl = 'http://localhost:8080/company/companyByName'
  }

  public getAllCompanies(): Observable<Array<Company>> {
    let headers = new HttpHeaders();
    headers.append('Content-Type', 'application/json');

    return this.http.get<Array<Company>>(this.getAllCompaniesUrl, {headers: headers});
  }

  public editCompany(company: Company): Observable<Boolean> {
    return this.http.put<boolean>(this.EditCompanyUrl, company);
  }

  public getCompanyByName(): Observable<Company> {
    let headers = new HttpHeaders();
    headers.append('Content-Type', 'application/json');
    let params = new HttpParams().set("name",sessionStorage.getItem('name'));

    return this.http.get<Company>(this.getCompanyByNameUrl, {headers: headers, params: params});
  }
}
