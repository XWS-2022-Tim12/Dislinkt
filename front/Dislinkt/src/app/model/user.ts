import { Timestamp } from "rxjs";


export class User {
    id: string;
    firstname: string;
    email: string;
    mobileNumber: string;
    gender: string;
    birthDay: Date;
    username: string;
    biography: string;
    experience: string;
    education: string;
    skills: string;
    interests: string;
    password: string;
    followingUsers: Array<string>
    followedByUsers: Array<string>
    followingRequests: Array<string>
    public: boolean
    blockedUsers: Array<string>
}
