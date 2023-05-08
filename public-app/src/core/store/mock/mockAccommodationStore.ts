import { produce } from "immer";
import { SetAppState, GetAppState, AppState } from "..";
import { AccommodationStoreType } from "../stores/accommodationStore";

export const mockAccommodationStore = (
  set: SetAppState,
  get: GetAppState
): Partial<AccommodationStoreType> => ({
  data: [],
  fetchAccommodations: async () => {
    set(
      produce((draft: AppState) => {
        draft.accommodation.data = [{ name: "123" }, { name: "456" }]
        return draft
      })
    )
  }
})