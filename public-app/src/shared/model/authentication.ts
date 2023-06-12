import { User } from "./user";

export interface AuthenticationResponse {
  Success?: boolean;
  Token?: string;
  ErrorMessage?: string;
  User: User;
}

export interface AuthenticationRequest {
  Username?: string;
  Password?: string;
}

export interface RegisterRequest {
  Role?: string;
  Username?: string;
  Password?: string;
  FirstName?: string;
  LastName?: string;
  Email?: string;
  LivingPlace?: string;
}

export interface RegisterResponse {
  Success?: boolean;
  Token?: string;
  Errors?: Error[];
  Message?: string;
  User?: User;
}

interface Error {
  PropertyName?: string;
  ErrorMessage?: string;
}

export interface UpdateUserInfoRequest {
  FirstName?: string;
  LastName?: string;
  LivingPlace?: string;
}

export interface UpdateUserInfoResponse {
  Success?: boolean;
  ErrorMessage?: string;
  User?: User;
}

export interface ChangePasswordRequest {
  CurrentPassword?: string;
  NewPassword?: string;
}