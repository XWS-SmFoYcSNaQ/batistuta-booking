import { produce } from "immer";
import { AuthenticationRequest, AuthenticationResponse, User } from "../../../shared/model";
import { AppState, GetAppState, SetAppState, apiUrl } from "../store";
import axios, { AxiosRequestConfig } from "axios";

export interface AuthStoreType {
  user: User | undefined
  loading: boolean
  login: (username: string, password: string) => Promise<boolean>
  setLoading: (val: boolean) => void
  logout: () => void
}

export const authStore = (
  set: SetAppState,
  get: GetAppState
): AuthStoreType => ({
  user: undefined,
  loading: false,
  login: async(username: string, password: string) => {
    get().auth.setLoading(true)
    try {
      const authenticationRequest : AuthenticationRequest = {
        Username: username,
        Password: password
      };
      const config : AxiosRequestConfig = {
        headers: {
          'Content-Type': 'application/json'
        }
      }
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
      if (e.response && e.response.data && e.response.data.errorMessage) {
        throw new Error(e.response.errorMessage);
      }
      throw new Error("Login error");
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
  }
})