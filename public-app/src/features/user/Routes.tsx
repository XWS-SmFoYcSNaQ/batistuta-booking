import { RouteObject } from "react-router";
import { HostList } from "./list";

export const userRoutes: RouteObject[] = [
  {
    path: "hosts",
    element: <HostList />,
  },
];
