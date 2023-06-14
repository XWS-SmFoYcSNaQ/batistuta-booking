import { RouteObject } from "react-router-dom";
import { RoomList } from "./list/RoomList";
import { SingleRoom } from "./single/SingleRoom";

export const roomRoutes: RouteObject[] = [
  {
    path: "",
    element: <RoomList />,
    children: [
    ]
  },
  {
    path: ":id",
    element: <SingleRoom />,
  }
];