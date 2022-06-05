import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Post } from '../model/post';

@Injectable({
  providedIn: 'root'
})
export class PostService {
  private createPostUrl: string;

  constructor(private http: HttpClient) { 
    this.createPostUrl = 'http://localhost:8000/user/post/newPost';
  }

  public createPost(post: Post): Observable<string> {
    return this.http.post<string>(this.createPostUrl, post,{withCredentials: true});
  }
}
