import { Discount, Period } from ".";

export interface Accommodation {
  id?: string;
  name?: string;
  benefits?: string;
  minGuests?: number;
  maxGuests?: number;
  basePrice?: number;
  location?: string;
  periods?: Period[];
  discounts?: Discount[];
  automaticReservation?: number;
}