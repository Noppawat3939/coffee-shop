import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";

export const Route = createFileRoute("/home")({
  component: lazyRouteComponent(() => import("~/pages/home")),
});
