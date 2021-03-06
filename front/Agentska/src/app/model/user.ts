export class User {
    id: string;
    firstName: string;
    email: string;
    mobileNumber: string;
    gender: string;
    birthDay: Date;
    username: string;
    password: string;
    role: Role;
}

export enum Role {
    admin="admin",
    agent_user="agent_user",
    agent_owner="agent_owner"
}
