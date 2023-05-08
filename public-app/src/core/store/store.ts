import { mergeDeepRight } from "ramda"
import { StoreApi, UseBoundStore, create } from "zustand";
import { immer } from "zustand/middleware/immer"
import { devtools, persist } from "zustand/middleware"
import { mockAccommodationStore } from "./mock/mockAccommodationStore";
import { AccommodationStoreType, accommodationStore } from "./stores/accommodationStore";

export interface AppState {
  accommodation: AccommodationStoreType
}

export const isStoreMocked = process.env.REACT_APP_USE_MOCK_STORE === "true"
export const apiUrl = process.env.REACT_APP_API_URL;
export type SetAppState = StoreApi<AppState>["setState"]
export type GetAppState = StoreApi<AppState>["getState"]

const storeGenerator = (set: SetAppState, get: GetAppState): AppState => ({
  accommodation: !isStoreMocked
    ? accommodationStore(set, get)
    : { ...accommodationStore(set, get), ...mockAccommodationStore(set, get) },
})

const storeMerge = (persistedState: any, currentState: any) => {
  return mergeDeepRight(currentState, persistedState)
}

export const appStore: UseBoundStore<StoreApi<AppState>> = create(
  devtools(immer(persist(storeGenerator, {
    name: "app-store",
    merge: storeMerge,
    partialize: (state: AppState) => ({})
  })))
)

