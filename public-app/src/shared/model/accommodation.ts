import { Discount, Period } from ".";
import { AccommodationRating } from "./rating";

export interface Accommodation {
  id?: string;
  name?: string;
  benefits?: string;
  minGuests?: number;
  maxGuests?: number;
  basePrice?: number;
  location?: string;
  automaticReservation?: number;
  periods?: Period[];
  discounts?: Discount[];
  ratings?: AccommodationRating[];
}