import { createFileRoute } from "@tanstack/react-router";
import BillByOrderNumber from "~/pages/bill/order_number";
import { order } from "~/services";

export const Route = createFileRoute("/bill/$order_number")({
  component: BillByOrderNumber,
  loader: ({ params: { order_number } }) =>
    order.getOrderByOrderNumber(order_number),
});
