import axios from "axios"
import { Period } from "../../../shared/model"
import { AppState, GetAppState, SetAppState, apiUrl } from "../store"
import { produce } from "immer"
import { roomRoutes } from "../../../features/room-reservation"

export interface PeriodStoreType {
  data: Period[]
  loading: boolean
  fetchPeriods: (accommodationId: string) => Promise<void>
  createPeriod: (data: Period) => Promise<void>
  clearData: () => void
  setLoading: (val: boolean) => void
}

export const periodStore = (
  set: SetAppState,
  get: GetAppState
): PeriodStoreType => ({
  data: [],
  loading: false,
  fetchPeriods: async (accommodationId: string) => {
    get().period.setLoading(true)
    get().period.clearData()
    try {
      const res = await axios.get(`${apiUrl}/accommodation/period/${accommodationId}`)
      set(
        produce((draft: AppState) => {
          draft.period.data = res.data.data.periods
          return draft
        })
      )
    } catch (e) {
      console.log(e)
    }
    get().period.setLoading(false)
  },
  createPeriod: async (data: Period) => {
    try {
      await axios.post(`${apiUrl}/accommodation/period`, data)
    } catch (e: any) {
      if(e.response && e.response.data && e.response.data.message){
        throw new Error(e.response.data.message)
      }
      throw new Error("Error while creating period.")
    }
  },
  clearData: () => {
    set(
      produce((draft: AppState) => {
        draft.period.data = []
        return draft
      })
    )
  },
  setLoading: (val: boolean) => {
    set(
      produce((draft: AppState) => {
        draft.period.loading = val;
        return draft
      })
    )
  }
})