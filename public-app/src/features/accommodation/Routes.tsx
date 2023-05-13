import { RouteObject } from "react-router-dom";
import { Create } from "./create";
import { Details } from "./details";
import { AccommodationList } from "./list";

export const accommodationRoutes: RouteObject[] = [
  {
    path: "",
    element: <AccommodationList/>,
    children: [
      {
        path: "create",
        element: <Create />,
      }
    ]
  },
  {
    path: ":id",
    element: <Details />,
  },
];
