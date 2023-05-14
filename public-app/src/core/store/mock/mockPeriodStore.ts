import { produce } from "immer"
import { SetAppState, GetAppState, AppState } from "../store"
import { PeriodStoreType } from "../stores/periodStore"

export const mockPeriodStore = (
  set: SetAppState,
  get: GetAppState
): Partial<PeriodStoreType> => ({
  data: [],
  fetchPeriods: async () => {
    set(
      produce((draft: AppState) => {
        draft.period.data = []
        return draft
      })
    )
  },
  createPeriod: async (data) => {},
})