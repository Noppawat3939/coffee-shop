import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";
import { order } from "~/services";

export const Route = createFileRoute("/transaction/orders")({
  component: lazyRouteComponent(() => import("~/pages/transaction/orders")),
  loader: async () => {
    const res = await order.getOrders({ page: 1, limit: 50 });

    return res;
  },
});
