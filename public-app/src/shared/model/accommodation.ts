import { Discount, Period } from ".";

export interface Accommodation {
  id?: string;
  name?: string;
  benefits?: string;
  minGuests?: number;
  maxGuests?: number;
  basePrice?: number;
  periods?: Period[];
  discounts?: Discount[];
}