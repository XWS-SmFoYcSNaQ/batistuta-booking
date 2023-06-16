import { Rating } from "./rating";

export interface User {
  Id?: string;
  Role?: UserRole;
  Username?: string;
  FirstName?: string;
  LastName?: string;
  Email?: string;
  LivingPlace?: string;
  Ratings?: Rating[];
}

export enum UserRole {
  Guest = 0,
  Host = 1
}