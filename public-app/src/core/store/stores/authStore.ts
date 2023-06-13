import { produce } from "immer";
import { AuthenticationRequest, AuthenticationResponse, User } from "../../../shared/model";
import { AppState, GetAppState, SetAppState, apiUrl } from "../store";
import axios, { AxiosRequestConfig } from "axios";
import { RegisterRequest, RegisterResponse, VerifyResponse } from "../../../shared/model/authentication";
import { toast } from "react-toastify";

export interface AuthStoreType {
  user: User | undefined
  loading: boolean
  userId: string | undefined
  login: (username: string, password: string) => Promise<boolean>
  register: (registerRequest: RegisterRequest) => Promise<boolean>
  setLoading: (val: boolean) => void
  logout: () => void
  verify: () => Promise<boolean>
}

const config : AxiosRequestConfig = {
  headers: {
    'Content-Type': 'application/json'
  }
}

export const authStore = (
  set: SetAppState,
  get: GetAppState
): AuthStoreType => ({
  user: undefined,
  loading: false,
  userId: undefined,
  login: async(username: string, password: string) => {
    get().auth.setLoading(true)
    try {
      const authenticationRequest : AuthenticationRequest = {
        Username: username,
        Password: password
      };
      const res = await axios.post<AuthenticationResponse>(`${apiUrl}/api/auth/login`, authenticationRequest,  config);
      if (res.data && res.data.Success) {
        window.localStorage.setItem("jwt", res.data.Token!);
        set(
          produce((draft: AppState) => {
            draft.auth.user = res.data.User;
            draft.auth.loading = false;
            return draft
          })
        );
        return true;
      }
      return false;
    } catch (e: any) {
      get().auth.loading = false;
      if (e.response && e.response.data && e.response.data.ErrorMessage) {
        toast.error(e.response.data.ErrorMessage);
        throw new Error(e.response.ErrorMessage);
      }
      throw new Error("Login error");
    }
  },
  register: async(registerRequest: RegisterRequest) => {
    get().auth.setLoading(true);
    try {
      const res = await axios.post<RegisterResponse>(`${apiUrl}/api/auth/register`, registerRequest, config);
      if (res.data && res.data.Success) {
        window.localStorage.setItem("jwt", res.data.Token!);
        set(
          produce((draft: AppState) => {
            draft.auth.user = res.data.User;
            draft.auth.loading = false;
            return draft
          })
        );
        return true;
      }
      toast.error(res.data.Message);
      return false;
    } catch(e: any) {
      get().auth.loading = false;
      if (e.response && e.response.data && e.response.data.errorMessage) {
        toast.error(e.response.data.Message);
        throw new Error(e.response.data.Message);
      }
      throw new Error("Registration error");
    }
  },
  setLoading: (val: boolean) => {
    set(
      produce((draft: AppState) => {
        draft.auth.loading = val;
        return draft
      })
    )
  },
  logout: () => {
    set (
      produce((draft: AppState) => {
        draft.auth.user = undefined
        return;
      })
    )
  },
  verify: async() => {
    try {
      const verifyConfig : AxiosRequestConfig = {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer '+window.localStorage.getItem("jwt")
        }
      }
      const resp = await axios.post<VerifyResponse>(`${apiUrl}/api/auth/verify`,{},  verifyConfig);
      if (resp.data && resp.data.Verified) {
        set(
          produce((draft: AppState) => {
            draft.auth.userId = resp.data.UserId;
            return draft
          })
        );
        return true;
      }
      return false;
    } catch (e: any) {
      get().auth.loading = false;
      if (e.response && e.response.data && e.response.data.errorMessage) {
        throw new Error(e.response.errorMessage);
      }
      throw new Error("Verify error");
    }
  }
})