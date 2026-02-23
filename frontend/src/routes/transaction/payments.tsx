import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";
import { dateFormat } from "~/helper";
import { payment } from "~/services";

export const Route = createFileRoute("/transaction/payments")({
  component: lazyRouteComponent(() => import("~/pages/transaction/payments")),
  loader: async () => {
    const res = await payment.getTransactions({ page: 1, limit: 50 });

    const data = res.data?.map((item) => ({
      ...item,
      amount: item.amount,
      expired_at: dateFormat(item.expired_at, "DD MMM YYYY, HH:mm"),
      created_at: dateFormat(item.created_at, "DD MMM YYYY, HH:mm"),
    }));

    return { data };
  },
});
