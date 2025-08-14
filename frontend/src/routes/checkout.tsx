import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";

export const Route = createFileRoute("/checkout")({
  component: lazyRouteComponent(() => import("~/pages/checkout")),
});
