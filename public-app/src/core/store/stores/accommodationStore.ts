import { Accommodation } from "../../../shared/model";
import { SetAppState, GetAppState, AppState, apiUrl } from "..";
import { produce } from "immer";
import axios from "axios";

export interface AccommodationStoreType {
  data: Accommodation[]
  fetchAccommodations: () => Promise<void>
  clearData: () => void
}

export const accommodationStore = (
  set: SetAppState,
  get: GetAppState
): AccommodationStoreType => ({
  data: [],
  fetchAccommodations: async () => {
    get().accommodation.clearData()
    
    try {
      const res = await axios.get(`${apiUrl}/accommodation`)
      set(
        produce((draft: AppState) => {
          draft.accommodation.data = res.data.data
          return draft
        })
      )
    } catch (e: any) {
      console.log(e)
    }
  },
  clearData: () => {
    set(
      produce((draft: AppState) => {
        draft.accommodation.data = []
        return draft
      })
    )
  }
})