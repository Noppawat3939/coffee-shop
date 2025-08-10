import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";

export const Route = createFileRoute("/menus")({
  component: lazyRouteComponent(() => import("~/pages/menus")),
});
