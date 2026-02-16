import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";
import { payment } from "~/services";

export const Route = createFileRoute("/transaction/payments")({
  component: lazyRouteComponent(() => import("~/pages/transaction/payments")),
  loader: async () => await payment.getTransactions({ page: 1, limit: 100 }),
});
