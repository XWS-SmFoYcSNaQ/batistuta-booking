import { Accommodation } from "../../../shared/model";
import { SetAppState, GetAppState, AppState, apiUrl } from "..";
import { produce } from "immer";
import axios from "axios";

export interface AccommodationStoreType {
  data: Accommodation[]
  accommodation: Accommodation | null
  loading: boolean
  fetchAccommodations: () => Promise<void>
  fetchDetails: (id: string, type: 'periods' | 'discounts') => Promise<void>
  createAccommodation: (data: Accommodation) => Promise<void>
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
  fetchDetails: async (id: string, type: 'periods' | 'discounts') => {
    get().accommodation.setLoading(true)
    get().accommodation.clearData()
    try {
      const res = await axios.get(`${apiUrl}/accommodation/details/${type}/${id}`)
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
      await axios.post(`${apiUrl}/accommodation`, data)
    } catch (e: any) {
      if(e.message){
        throw new Error(e.message)
      }
      throw new Error("Error while creating accommodation.")
    }
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