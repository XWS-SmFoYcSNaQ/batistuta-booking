import { produce } from "immer";
import { AuthenticationRequest, AuthenticationResponse, User } from "../../../shared/model";
import { AppState, GetAppState, SetAppState, apiUrl } from "../store";
import axios, { AxiosRequestConfig } from "axios";
import { RegisterRequest, RegisterResponse, UpdateUserInfoRequest, UpdateUserInfoResponse } from "../../../shared/model/authentication";
import { toast } from "react-toastify";

export interface AuthStoreType {
  user: User | undefined
  loading: boolean
  login: (username: string, password: string) => Promise<boolean>
  register: (registerRequest: RegisterRequest) => Promise<boolean>
  setLoading: (val: boolean) => void
  logout: () => void
  updateUserInfo: (updateUserInfoRequest: UpdateUserInfoRequest) => Promise<boolean>
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
  setLoading: (val: boolean) => {
    set(
      produce((draft: AppState) => {
        draft.auth.loading = val;
        return draft
      })
    )
  },
  login: async(username: string, password: string) => {
    set(
      produce((draft: AppState) => {
        draft.auth.loading = true;
        return draft;
      })
    )
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
      set(
        produce((draft: AppState) => {
          draft.auth.loading = false;
          return draft;
        })
      )
      if (e.response && e.response.data && e.response.data.ErrorMessage) {
        toast.error(e.response.data.ErrorMessage);
        throw new Error(e.response.ErrorMessage);
      }
      throw new Error("Login error");
    }
  },
  register: async(registerRequest: RegisterRequest) => {
    set(
      produce((draft: AppState) => {
        draft.auth.loading = true;
        return draft;
      })
    )
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
      set(
        produce((draft: AppState) => {
          draft.auth.loading = false;
          return draft;
        })
      )
      if (e.response && e.response.data && e.response.data.errorMessage) {
        toast.error(e.response.data.Message);
        throw new Error(e.response.data.Message);
      }
      toast.error("Registration Error", { position: "top-center"});
      return false;
    }
  },
  updateUserInfo: async (updateUserInfoRequest: UpdateUserInfoRequest) => {
    set(
      produce((draft: AppState) => {
        draft.auth.loading = true;
        return draft;
      })
    )
    try {
      const res = await axios.patch<UpdateUserInfoResponse>(`${apiUrl}/api/users/${get().auth.user?.Username}`, updateUserInfoRequest, {
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${window.localStorage.getItem("jwt")}`
        }
      });
      if (res.data && res.data.Success) {
        set(
          produce((draft: AppState) => {
            draft.auth.user = res.data.User;
            draft.auth.loading = false;
            return draft;
          })
        );
        toast.success("Your information has been updated successfully.");
        return true;
      }
      toast.error(res.data.ErrorMessage);
      return false;
    }
    catch (e: any) {
      set(
        produce((draft: AppState) => {
          draft.auth.loading = false;
          return draft;
        })
      )
      if (e.response && e.response.data && e.response.data.errorMessage) {
        toast.error(e.response.data.Message);
        throw new Error(e.response.data.Message);
      }
      toast.error("Error updating user information", {position: "top-center"});
      return false;
    }
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