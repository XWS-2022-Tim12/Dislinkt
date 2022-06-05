import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Post } from '../model/post';

@Injectable({
  providedIn: 'root'
})
export class PostService {
  private createPostUrl: string;
  private getUserPostsUrl: string;
  private likePostUrl: string;
  private dislikePostUrl: string;
  private commentPostUrl: string;

  constructor(private http: HttpClient) { 
    this.createPostUrl = 'http://localhost:8000/user/post/newPost';
    this.getUserPostsUrl = 'http://localhost:8000/user/post/findUserPosts/';
    this.likePostUrl = 'http://localhost:8000/user/post/likePost';
    this.dislikePostUrl = 'http://localhost:8000/user/post/dislikePost';
    this.commentPostUrl = 'http://localhost:8000/user/post/commentPost';
  }

  public createPost(post: Post): Observable<string> {
    return this.http.post<string>(this.createPostUrl, post,{withCredentials: true});
  }

  public getUserPosts(username: string): Observable<Array<Post>> {
    return this.http.get<Array<Post>>(this.getUserPostsUrl + username,{withCredentials: true});
  }

  public likePost(post: Post): Observable<string> {
    return this.http.put<string>(this.likePostUrl, post, {withCredentials: true, responseType: 'text' as 'json'});
  }

  public dislikePost(post: Post): Observable<string> {
    return this.http.put<string>(this.dislikePostUrl, post, {withCredentials: true, responseType: 'text' as 'json'});
  }

  public commentPost(post: Post): Observable<string> {
    return this.http.put<string>(this.commentPostUrl, post, {withCredentials: true, responseType: 'text' as 'json'});
  }
}
