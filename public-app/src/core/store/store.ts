import { mergeDeepRight } from "ramda"
import { StoreApi, UseBoundStore, create } from "zustand";
import { immer } from "zustand/middleware/immer"
import { devtools, persist } from "zustand/middleware"
import { AccommodationStoreType, accommodationStore } from "./stores/accommodationStore";
import { PeriodStoreType, periodStore } from "./stores/periodStore";
import { DiscountStoreType, discountStore } from "./stores/discountStore";
import { AuthStoreType, authStore } from "./stores/authStore";
import { UserStoreType, userStore } from "./stores/userStore";
import { RatingStoreType, ratingStore } from "./stores/ratingStore";
import { NotificationStoreType, notificationStore } from "./stores/notificationStore";

export interface AppState {
  accommodation: AccommodationStoreType
  period: PeriodStoreType
  discount: DiscountStoreType
  auth: AuthStoreType
  user: UserStoreType
  rating: RatingStoreType
  notification: NotificationStoreType
}

export const apiUrl = process.env.REACT_APP_API_URL;
export type SetAppState = StoreApi<AppState>["setState"]
export type GetAppState = StoreApi<AppState>["getState"]

const storeGenerator = (set: SetAppState, get: GetAppState): AppState => ({
  accommodation: accommodationStore(set, get),
  period: periodStore(set, get),
  discount: discountStore(set, get),
  auth: authStore(set, get),
  user: userStore(set, get),
  rating: ratingStore(set, get),
  notification: notificationStore(set, get)
})

const storeMerge = (persistedState: any, currentState: any) => {
  return mergeDeepRight(currentState, persistedState)
}

export const appStore: UseBoundStore<StoreApi<AppState>> = create(
  devtools(immer(persist(storeGenerator, {
    name: "app-store",
    merge: storeMerge,
    partialize: (state: AppState) => ({ auth: state.auth, notification: state.notification })
  })))
)

