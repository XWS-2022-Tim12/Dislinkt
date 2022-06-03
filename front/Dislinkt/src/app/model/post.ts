export class Post {
    id: string;
    text: string;
    image: string;
    link: string;
    likes: number;
    dislikes: number;
    comments: Array<string>;
    username: string;
    imageContent: any;
}