import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";

type SearchParams = Readonly<
  Record<"order_number" | "transaction_number", string>
>;

export const Route = createFileRoute("/checkout")({
  component: lazyRouteComponent(() => import("~/pages/checkout")),
  validateSearch: (search: SearchParams) => search,
});
