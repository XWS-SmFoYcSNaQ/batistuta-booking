import { Accommodation } from "../../../shared/model";
import { SetAppState, GetAppState, AppState, apiUrl } from "..";
import { produce } from "immer";
import axios from "axios";

export interface AccommodationStoreType {
  data: Accommodation[]
  loading: boolean
  fetchAccommodations: () => Promise<void>
  createAccommodation: (data: Accommodation) => Promise<void>
  clearData: () => void
  setLoading: (val: boolean) => void
}

export const accommodationStore = (
  set: SetAppState,
  get: GetAppState
): AccommodationStoreType => ({
  data: [],
  loading: false,
  fetchAccommodations: async () => {
    get().accommodation.setLoading(true)
    get().accommodation.clearData()
    try {
      const res = await axios.get(`${apiUrl}/accommodation`)
      set(
        produce((draft: AppState) => {
          draft.accommodation.data = res.data.data
          return draft
        })
      )
    } catch (e) {
      console.log(e)
    }
    get().accommodation.setLoading(false)
  },
  createAccommodation: async (data: Accommodation) => {
    try {
      await axios.post(`${apiUrl}/accommodation`, data)
    } catch (e) {
      console.log(e)
      throw new Error("Error while creating accommodation.")
    }
  },
  clearData: () => {
    set(
      produce((draft: AppState) => {
        draft.accommodation.data = []
        return draft
      })
    )
  },
  setLoading: (val: boolean) => {
    set(
      produce((draft: AppState) => {
        draft.accommodation.loading = val;
        return draft
      })
    )
  }
})