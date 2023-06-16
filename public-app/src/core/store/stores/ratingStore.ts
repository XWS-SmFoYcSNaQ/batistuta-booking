import axios from "axios"
import { AppState, GetAppState, SetAppState, apiUrl } from "../store"
import { Rating } from "../../../shared/model"
import { produce } from "immer"

export interface RatingStoreType {
  data: Rating[]
  loading: boolean
  fetchRatingsByTargetId: (ratingId: string) => Promise<void>
  removeRating: (ratingId: string) => Promise<void>
  clearData: () => void
  setLoading: (val: boolean) => void
}

export const ratingStore = (
  set: SetAppState,
  get: GetAppState
): RatingStoreType => ({
  data: [],
  loading: false,
  fetchRatingsByTargetId: async (ratingId: string) => {
    get().rating.setLoading(true)
    get().rating.clearData()
    try {
      const res = await axios.get(`${apiUrl}/rating/target/${ratingId}`)
      set(
        produce((draft: AppState) => {
          draft.rating.data = res.data.Data
          return draft
        })
      )
    } catch (e) {
      console.log(e)
    }
    get().rating.setLoading(false)
  },
  removeRating: async (ratingId: string) => {
    try {
      await axios.delete(`${apiUrl}/rating/${ratingId}`, {
        headers: {
          Authorization: `Bearer ${window.localStorage.getItem("jwt")}`
        }
      })
    } catch (e: any) {
      if (e.response && e.response.data && e.response.data.message) {
        throw new Error(e.response.data.message)
      }
      throw new Error("Error while deleting rating.")
    }
  },
  clearData: () => {
    set(
      produce((draft: AppState) => {
        draft.rating.data = []
        return draft
      })
    )
  },
  setLoading: (val: boolean) => {
    set(
      produce((draft: AppState) => {
        draft.rating.loading = val;
        return draft
      })
    )
  }
})