import { RouteObject } from "react-router-dom";
import { Create } from "./create";

export const accommodationRoutes: RouteObject[] = [
  {
    path: "create",
    element: <Create/>
  },
];
