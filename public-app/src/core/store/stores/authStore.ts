import { produce } from "immer";
import { AuthenticationRequest, AuthenticationResponse, User } from "../../../shared/model";
import { AppState, GetAppState, SetAppState, apiUrl } from "../store";
import axios, { AxiosRequestConfig } from "axios";
import { ChangePasswordRequest, RegisterRequest, RegisterResponse, UpdateUserInfoRequest, UpdateUserInfoResponse, VerifyResponse } from "../../../shared/model/authentication";
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
  updateUserInfo: (updateUserInfoRequest: UpdateUserInfoRequest) => Promise<boolean>
  changePassword: (changePasswordRequest: ChangePasswordRequest) => Promise<boolean>
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
    let success = false;
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
            draft.auth.user.Role = (res.data.User.Role + "") === "Guest" ? 0 : 1 
            draft.auth.loading = false;
            return draft
          })
        );
        success = true;
        const appState = get();
        await appState.notification.connect();
        toast.success(`Logged in successfully`, { position: "top-center" });
      }
    } catch (e: any) {
      if (e.response && e.response.data && e.response.data.message) {
        console.log(e);
        toast.error(e.response.data.message, { position: "top-center"});
      }
    } finally {
      set(
        produce((draft: AppState) => {
          draft.auth.loading = false;
          return draft;
        })
      )
      return success;
    }
  },
  register: async(registerRequest: RegisterRequest) => {
    set(
      produce((draft: AppState) => {
        draft.auth.loading = true;
        return draft;
      })
    )
    let success = false;
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
        success = true;
        const appState = get();
        await appState.notification.connect();
      }
    } catch(e: any) {
      if (e.response && e.response.data && e.response.data.message) {
        console.log(e);
        toast.error(e.response.data.message, { position: "top-center"});
      }
    } finally {
      set(
        produce((draft: AppState) => {
          draft.auth.loading = false;
          return draft;
        })
      )
      return success;
    }
  },
  updateUserInfo: async (updateUserInfoRequest: UpdateUserInfoRequest) => {
    set(
      produce((draft: AppState) => {
        draft.auth.loading = true;
        return draft;
      })
    )
    let success = false;
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
        toast.success("Your information has been updated successfully.", { position: "top-center" });
        success = true;
      }
    } catch (e: any) {
      if (e.response && e.response.data && e.response.data.message) {
        console.log(e);
        toast.error(e.response.data.message, { position: "top-center"});
      }
    } finally {
      set(
        produce((draft: AppState) => {
          draft.auth.loading = false;
          return draft;
        })
      )
      return success;
    }
  },
  changePassword: async (changePasswordRequest: ChangePasswordRequest) => {
    set(
      produce((draft: AppState) => {
        draft.auth.loading = true;
        return draft;
      })
    )
    let success = false;
    try {
      const res = await axios.patch(`${apiUrl}/api/users/password`, changePasswordRequest, {
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${window.localStorage.getItem("jwt")}`
        }
      });
      if (res.status === 200) {
        set(
          produce((draft: AppState) => {
            draft.auth.loading = false;
            return draft;
          })
        )
        toast.success("Password updated successfully.", { position: "top-center" });
        success = true;
      }
    } catch (e: any) {
      if (e.response && e.response.data && e.response.data.message) {
        console.log(e);
        toast.error(e.response.data.message, { position: "top-center"});
      }
    } finally {
      set(
        produce((draft: AppState) => {
          draft.auth.loading = false;
          return draft;
        })
      )
      return success;
    }
  },
  logout: async () => {
    const state = get();
    await state && state.notification && state.notification.connection && state.notification.connection.stop();
    set((draft: AppState) => {
      draft.notification.connection = null;
      draft.notification.data = [];
      draft.auth.user = undefined;
      return draft;
    });
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