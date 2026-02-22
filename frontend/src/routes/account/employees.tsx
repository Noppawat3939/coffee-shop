import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";
import { employee } from "~/services";

export const Route = createFileRoute("/account/employees")({
  component: lazyRouteComponent(() => import("~/pages/account/employees")),
  loader: async () => {
    const res = await employee.getEmployees({ page: 1, limit: 50 });
    return res;
  },
});
