import { Component, OnInit } from '@angular/core';
import { Post } from 'src/app/model/post';
import { PostService } from 'src/app/services/post.service';

@Component({
  selector: 'app-new-post',
  templateUrl: './new-post.component.html',
  styleUrls: ['./new-post.component.css']
})
export class NewPostComponent implements OnInit {
  text: string;
  file: File;
  bytes: any;
  emptyFields: boolean = false;

  constructor(private postService: PostService) { }

  ngOnInit(): void {
  }

  getFile(event: any) {
    this.file = event.target.files[0];
    const reader = new FileReader();
    reader.onload = (e: any) => {
      this.bytes = e.target.result.split('base64,')[1];
    };
    reader.readAsDataURL(this.file);
  }

  createPost() {
    if (!this.text && !this.file) {
      this.emptyFields = true;
    }
    else {
      this.emptyFields = false;

      let post = new Post();
      post.text = this.text;
      post.imageContent = this.bytes;

      this.postService.createPost(post).subscribe(ret => {
        
      })
    }
  }
}
