export interface User {
  Role?: UserRole;
  Username?: string;
  FirstName?: string;
  LastName?: string;
  Email?: string;
  LivingPlace?: string;
}

export enum UserRole {
  Guest = 0,
  Host = 1
}