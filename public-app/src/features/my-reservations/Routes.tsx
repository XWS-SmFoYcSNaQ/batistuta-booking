import { RouteObject } from "react-router-dom";
import { MyReservations } from "./MyReservations";
import { AcceptedReservations } from "./AcceptedReservations";
import { PendingReservations } from "./PendingReservations";

export const reservationsRoutes: RouteObject[] = [
  {
    path: "",
    element: <MyReservations />,
    children: [
      {
        path: "accepted",
        element: <AcceptedReservations />,
      },
      {
        path: "pending",
        element: <PendingReservations />,
      }
    ]
  }
];