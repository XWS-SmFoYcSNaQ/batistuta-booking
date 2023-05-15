import App from "./App";
import { AccommodationRoot, accommodationRoutes } from "./features/accommodation";
import { RoomsRoot, roomRoutes } from "./features/room-reservation";
import Login from "./features/auth/login/Login";
import Register from "./features/auth/register/Register";
import UserProfile from "./features/profile/UserProfile";
import { ErrorPage } from "./shared/ErrorPage";
import { createBrowserRouter } from "react-router-dom";
import { AllAccommodations } from "./features/accommodation/AllAccommodations";
import { MyReservations, reservationsRoutes } from "./features/my-reservations";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "accommodation",
        element: <AccommodationRoot />,
        children: accommodationRoutes
      },
      {
        path: "rooms",
        element: <RoomsRoot />,
        children: roomRoutes
      },
      {
        path: "reservations",
        element: <MyReservations/>,
        children: reservationsRoutes
      },
      {
        path: "login",
        element: <Login/>
      },
      {
        path: "all-accommodations",
        element: <AllAccommodations/>
      },
      {
        path: "register",
        element: <Register/>
      },
      {
        path: "profile",
        element: <UserProfile/>
      }
    ]
  },
]);
