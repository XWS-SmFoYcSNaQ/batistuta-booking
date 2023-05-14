import { User } from "./user";

export interface AuthenticationResponse {
  Success?: boolean;
  Token?: string;
  Errors?: Error[];
  Message?: string;
  User: User;
}

export interface AuthenticationRequest {
  Username?: string;
  Password?: string;
}

interface Error {
  PropertyName?: string;
  ErrorMessage?: string;
}