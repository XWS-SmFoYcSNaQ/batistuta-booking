import App from "./App";
import {
  AccommodationRoot,
  accommodationRoutes,
} from "./features/accommodation";
import { RoomsRoot, roomRoutes } from "./features/room-reservation";
import {
  MyReservations,
  reservationsRoutes,
} from "./features/reservations/my-reservations";
import { ReservationConfirmation } from "./features/reservations/reservations-to-confirm/reservationConfirmation";
import Login from "./features/auth/login/Login";
import Register from "./features/auth/register/Register";
import UserProfile from "./features/profile/UserProfile";
import { userRoutes } from "./features/user";
import { UserRoot } from "./features/user/UserRoot";
import { ErrorPage } from "./shared/ErrorPage";
import { createBrowserRouter } from "react-router-dom";
import { RatingList } from "./features/rating"

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "accommodation",
        element: <AccommodationRoot />,
        children: accommodationRoutes,
      },
      {
        path: "user",
        element: <UserRoot />,
        children: userRoutes,
      },
      {
        path: "login",
        element: <Login />,
      },
      {
        path: "register",
        element: <Register />,
      },
      {
        path: "profile",
        element: <UserProfile />,
      },
      {
        path: "rooms",
        element: <RoomsRoot />,
        children: roomRoutes,
      },
      {
        path: "reservations",
        element: <MyReservations />,
        children: reservationsRoutes,
      },
      {
        path: "reservations-to-confirm",
        element: <ReservationConfirmation />,
      },
      {
        path: "ratings/:id",
        element: <RatingList />
      }
    ],
  },
]);
