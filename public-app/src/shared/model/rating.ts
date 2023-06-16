export interface Rating {
  Id: string;
  UserId: string;
  Value: number;
  TargetId: string;
  TargetType: number;
  LastModified: string;
  UserFirstName?: string;
  UserLastName?: string;
  UserEmail?: string;
}

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