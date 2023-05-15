import App from "./App";
import { AccommodationRoot, accommodationRoutes } from "./features/accommodation";
import Login from "./features/auth/login/Login";
import Register from "./features/auth/register/Register";
import UserProfile from "./features/profile/UserProfile";
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
        path: "login",
        element: <Login/>
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
