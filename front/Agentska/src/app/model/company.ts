import { User } from "./user";

export class Company {
    id: string;
    name: string;
    address: string;
    phoneNumber: string;
    description: string;
    owner: User;
    approved: boolean;
}
