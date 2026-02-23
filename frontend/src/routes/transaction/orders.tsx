import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";
import { dateFormat } from "~/helper";
import { order } from "~/services";

export const Route = createFileRoute("/transaction/orders")({
  component: lazyRouteComponent(() => import("~/pages/transaction/orders")),
  loader: async () => {
    const res = await order.getOrders({ page: 1, limit: 50 });

    const data = res.data?.map((item) => ({
      ...item,
      created_at: dateFormat(item.created_at, "DD MMM YYYY, HH:mm"),
    }));

    return { data };
  },
});
