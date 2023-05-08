import { Accommodation } from "../../../shared/model";
import { SetAppState, GetAppState } from "..";

export interface AccommodationStoreType {
  data: Accommodation[]
  fetchAccommodations: () => Promise<void>
}

export const accommodationStore = (
  set: SetAppState,
  get: GetAppState
): AccommodationStoreType => ({
  data: [],
  fetchAccommodations: async () => {}
})