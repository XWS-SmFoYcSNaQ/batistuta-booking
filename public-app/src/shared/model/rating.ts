export interface AccommodationRating {
  id: string;
  userId: string;
  value: number;
  accommodationId?: string;
}

export interface HostRating {
  id: string;
  userId: string;
  value: number;
  hostId?: string;
}