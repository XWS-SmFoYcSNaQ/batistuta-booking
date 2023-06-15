import { produce } from "immer";
import { User } from "../../../shared/model";
import { SetAppState, GetAppState, AppState, apiUrl } from "../store";
import axios from "axios";

export interface UserStoreType {
  data: User[]
  loading: boolean
  fetchHosts: () => Promise<void>
  rateHost: ({ id, value }: { id: string, value: number }) => Promise<void>
  clearData: () => void
  setLoading: (val: boolean) => void
}

export const userStore = (
  set: SetAppState,
  get: GetAppState
): UserStoreType => ({
  data: [],
  loading: false,
  fetchHosts: async () => {
    get().user.setLoading(true)
    get().user.clearData()
    try {
      const res = await axios.get(`${apiUrl}/api/hosts`)
      set(
        produce((draft: AppState) => {
          draft.user.data = res.data.Hosts
          return draft
        })
      )
    } catch (e) {
      console.log(e)
    }
    get().user.setLoading(false)
  },
  rateHost: async ({ id, value }: { id: string, value: number }) => {
    const data = {
      HostId: id,
      Value: value
    }
    try {
      await axios.post(`${apiUrl}/rating/host`, data, {
        headers: {
          Authorization: `Bearer ${window.localStorage.getItem("jwt")}`
        }
      })
    } catch (e: any) {
      if (e.response && e.response.data && e.response.data.message) {
        throw new Error(e.response.data.message)
      }
      throw new Error("Error while rating host.")
    }
  },
  clearData: () => {
    set(
      produce((draft: AppState) => {
        draft.user.data = []
        return draft
      })
    )
  },
  setLoading: (val: boolean) => {
    set(
      produce((draft: AppState) => {
        draft.user.loading = val;
        return draft
      })
    )
  }
})