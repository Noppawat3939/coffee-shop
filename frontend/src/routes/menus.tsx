import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";
import { menu } from "~/services";

export const Route = createFileRoute("/menus")({
  component: lazyRouteComponent(() => import("~/pages/menus")),
  loader: menu.getMenus,
});
