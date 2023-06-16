export interface Reservation {
    id?: string;
    accommodationId?: string;
    startDate?: string;
    endDate?: string;
    userId?: string;
    numberOfGuests?: number;
    numberOfCanceledReservations?: number
  }