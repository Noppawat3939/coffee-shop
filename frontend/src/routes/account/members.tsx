import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";
import { dateFormat } from "~/helper";
import { member } from "~/services";

const MAX_LIMIT = 50;

export const Route = createFileRoute("/account/members")({
  component: lazyRouteComponent(() => import("~/pages/account/members")),
  loader: async () => {
    const res = await member.getMembers({ page: 1, limit: MAX_LIMIT });

    const data = res?.data?.map((item) => ({
      ...item,
      created_at: dateFormat(item.created_at, "DD/MM/YYYY"),
    }));

    return { data };
  },
});
