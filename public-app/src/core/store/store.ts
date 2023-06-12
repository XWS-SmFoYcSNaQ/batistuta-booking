import { mergeDeepRight } from "ramda"
import { StoreApi, UseBoundStore, create } from "zustand";
import { immer } from "zustand/middleware/immer"
import { devtools, persist } from "zustand/middleware"
import { mockAccommodationStore } from "./mock/mockAccommodationStore";
import { AccommodationStoreType, accommodationStore } from "./stores/accommodationStore";
import { PeriodStoreType, periodStore } from "./stores/periodStore";
import { mockPeriodStore } from "./mock/mockPeriodStore";
import { DiscountStoreType, discountStore } from "./stores/discountStore";
import { mockDiscountStore } from "./mock/mockDiscountStore";
import { AuthStoreType, authStore } from "./stores/authStore";

export interface AppState {
  accommodation: AccommodationStoreType
  period: PeriodStoreType
  discount: DiscountStoreType
  auth: AuthStoreType
}

export const isStoreMocked = process.env.REACT_APP_USE_MOCK_STORE === "true"
export const apiUrl = process.env.REACT_APP_API_URL;
export type SetAppState = StoreApi<AppState>["setState"]
export type GetAppState = StoreApi<AppState>["getState"]

const storeGenerator = (set: SetAppState, get: GetAppState): AppState => ({
  accommodation: !isStoreMocked
    ? accommodationStore(set, get)
    : { ...accommodationStore(set, get), ...mockAccommodationStore(set, get) },
  period: !isStoreMocked
    ? periodStore(set, get)
    : { ...periodStore(set, get), ...mockPeriodStore(set, get) },
  discount: !isStoreMocked
    ? discountStore(set, get)
    : { ...discountStore(set, get), ...mockDiscountStore(set, get) },
  auth: authStore(set, get)
})

const storeMerge = (persistedState: any, currentState: any) => {
  return mergeDeepRight(currentState, persistedState)
}

export const appStore: UseBoundStore<StoreApi<AppState>> = create(
  devtools(immer(persist(storeGenerator, {
    name: "app-store",
    merge: storeMerge,
    partialize: (state: AppState) => ({auth:state.auth})
  })))
)

