import { RouteObject } from "react-router-dom";
import { Create } from "./create";
import { Availability } from "./availability";
import { AccommodationList } from "./list";
import { Discounts } from "./discounts";

export const accommodationRoutes: RouteObject[] = [
  {
    path: "my",
    element: <AccommodationList/>,
    children: [
      {
        path: "create",
        element: <Create />,
      }
    ]
  },
  {
    path: "all",
    element: <AccommodationList host={false}/>
  },
  {
    path: "availability/:id",
    element: <Availability />,
  },
  {
    path: "discounts/:id",
    element: <Discounts />,
  },
];
