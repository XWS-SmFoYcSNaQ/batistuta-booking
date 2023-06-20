export interface Notification {
  id?: string;
  title?: string;
  content?: string;
  type?: NotificationType;
  createdAt?: Date;
  new?: boolean;
}

export enum NotificationType {
  ReservationRequestCreated = 0,
  ReservationCancelled = 1,
  HostRated = 2,
  AccommodationRated = 3,
  HostFeaturedStatusChanged = 4,
  ReservationRequestResponded = 5
}