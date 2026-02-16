import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";

export const Route = createFileRoute("/transaction/orders")({
  component: lazyRouteComponent(() => import("~/pages/transaction/orders")),
});
