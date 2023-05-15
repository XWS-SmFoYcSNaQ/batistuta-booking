import { User } from "./user";

export interface AuthenticationResponse {
  Success?: boolean;
  Token?: string;
  Errors?: Error[];
  Message?: string;
  User: User;
}

export interface VerifyResponse {
  Verified?: boolean;
  ErrorMessage?: string;
  UserId?: string;
  UserRole?: string;
}

export interface AuthenticationRequest {
  Username?: string;
  Password?: string;
}

interface Error {
  PropertyName?: string;
  ErrorMessage?: string;
}