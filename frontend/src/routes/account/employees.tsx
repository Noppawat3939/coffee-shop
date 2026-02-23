import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";
import { dateFormat } from "~/helper";
import { employee } from "~/services";

export const Route = createFileRoute("/account/employees")({
  component: lazyRouteComponent(() => import("~/pages/account/employees")),
  loader: async () => {
    const res = await employee.getEmployees({ page: 1, limit: 50 });

    const data = res?.data?.map((item) => ({
      ...item,
      created_at: dateFormat(item?.created_at, "DD MMM YYYY"),
    }));

    return { data };
  },
});
