import App from "./App";
import { AccommodationRoot, accommodationRoutes } from "./features/accommodation";
import { RoomsRoot, roomRoutes } from "./features/room-reservation";
import Login from "./features/auth/login/Login";
import { ErrorPage } from "./shared/ErrorPage";
import { createBrowserRouter } from "react-router-dom";

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
        path: "login",
        element: <Login/>
      }
    ]
  },
]);
