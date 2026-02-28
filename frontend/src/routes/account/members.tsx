import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";
import { dateFormat } from "~/helper";
import { member } from "~/services";

export const Route = createFileRoute("/account/members")({
  component: lazyRouteComponent(() => import("~/pages/account/members")),
  loader: async () => {
    const res = await member.getMembers({ page: 1, limit: 50 });

    const data = res?.data?.map((item) => ({
      ...item,
      created_at: dateFormat(item.created_at, "DD/MM/YYYY"),
    }));

    return { data };
  },
});
