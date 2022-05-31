import { User } from "./user";

export class Company {
    id: string;
    name: string;
    address: string;
    phoneNumber: string;
    description: string;
    owner: User;
    approved: boolean;
    comments: Array<String>;
    juniorSalary: Array<Number>;
    mediorSalary: Array<Number>;
    hrInterviews: Array<String>;
    tehnicalInterviews: Array<String>;
}
