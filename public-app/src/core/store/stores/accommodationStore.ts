import { Accommodation } from "../../../shared/model";
import { SetAppState, GetAppState, AppState, apiUrl } from "..";
import { produce } from "immer";
import axios from "axios";

export interface AccommodationStoreType {
  data: Accommodation[]
  accommodation: Accommodation | null
  loading: boolean
  fetchAccommodations: () => Promise<void>
  fetchMyAccommodations: () => Promise<void>
  fetchDetails: (id: string) => Promise<void>
  createAccommodation: (data: Accommodation) => Promise<void>
  rateAccommodation: ({ id, value }: { id: string, value: number }) => Promise<void>
  clearData: () => void
  setLoading: (val: boolean) => void
}

export const accommodationStore = (
  set: SetAppState,
  get: GetAppState
): AccommodationStoreType => ({
  data: [],
  accommodation: null,
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
  fetchMyAccommodations: async () => {
    get().accommodation.setLoading(true)
    get().accommodation.clearData()
    try {
      const res = await axios.get(`${apiUrl}/accommodation/me`, {
        headers: {
          Authorization: `Bearer ${window.localStorage.getItem("jwt")}`
        }
      })
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
  fetchDetails: async (id: string) => {
    get().accommodation.setLoading(true)
    get().accommodation.clearData()
    try {
      const res = await axios.get(`${apiUrl}/accommodation/details/${id}`)
      set(
        produce((draft: AppState) => {
          draft.accommodation.accommodation = res.data
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
      await axios.post(`${apiUrl}/accommodation`, data, {
        headers: {
          Authorization: `Bearer ${window.localStorage.getItem("jwt")}`
        }
      })
    } catch (e: any) {
      if (e.response && e.response.data && e.response.data.message) {
        throw new Error(e.response.data.message)
      }
      throw new Error("Error while creating accommodation.")
    }
  },
  rateAccommodation: async ({ id, value }: { id: string, value: number }) => {
    console.log(id, value)
  },
  clearData: () => {
    set(
      produce((draft: AppState) => {
        draft.accommodation.data = []
        draft.accommodation.accommodation = null
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