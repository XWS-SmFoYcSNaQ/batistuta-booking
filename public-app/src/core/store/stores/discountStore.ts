import { produce } from "immer"
import { AppState, GetAppState, SetAppState, apiUrl } from ".."
import { Discount } from "../../../shared/model"
import axios from "axios"

export interface DiscountStoreType {
  data: Discount[]
  loading: boolean
  fetchDiscounts: (accommodationId: string) => Promise<void>
  createDiscount: (data: Discount) => Promise<void>
  clearData: () => void
  setLoading: (val: boolean) => void
}

export const discountStore = (
  set: SetAppState,
  get: GetAppState
): DiscountStoreType => ({
  data: [],
  loading: false,
  fetchDiscounts: async (accommodationId: string) => {
    get().discount.setLoading(true)
    get().discount.clearData()
    try {
      const res = await axios.get(`${apiUrl}/accommodation/discount/${accommodationId}`)
      set(
        produce((draft: AppState) => {
          draft.discount.data = res.data.data
          return draft
        })
      )
    } catch (e) {
      console.log(e)
    }
    get().discount.setLoading(false)
  },
  createDiscount: async (data: Discount) => {
    try {
      await axios.post(`${apiUrl}/accommodation/discount`, data, {
        headers: {
          Authorization: `Bearer ${window.localStorage.getItem("jwt")}`
        }
      })
    } catch (e: any) {
      if(e.response && e.response.data && e.response.data.message){
        throw new Error(e.response.data.message)
      }
      throw new Error("Error while creating discount.")
    }
  },
  clearData: () => {
    set(
      produce((draft: AppState) => {
        draft.discount.data = []
        return draft
      })
    )
  },
  setLoading: (val: boolean) => {
    set(
      produce((draft: AppState) => {
        draft.discount.loading = val;
        return draft
      })
    )
  }
})