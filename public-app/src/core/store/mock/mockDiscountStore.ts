import { produce } from "immer"
import { SetAppState, GetAppState, AppState } from "../store"
import { DiscountStoreType } from "../stores/discountStore"

export const mockDiscountStore = (
  set: SetAppState,
  get: GetAppState
): Partial<DiscountStoreType> => ({
  data: [],
  fetchDiscounts: async () => {
    set(
      produce((draft: AppState) => {
        draft.discount.data = []
        return draft
      })
    )
  },
  createDiscount: async (data) => {},
})