import { createFileRoute, lazyRouteComponent } from "@tanstack/react-router";
import { employee } from "~/services";
import type { TVerifyUserResponse } from "~/services/auth";

export const Route = createFileRoute("/profile")({
  component: lazyRouteComponent(() => import("~/pages/profile")),

  loader: async ({ context }) => {
    const contextData = context as TVerifyUserResponse;

    if (!contextData?.data?.id) return;
    // get id from context and call to service
    const response = await employee.getEmployeeByID(contextData.data.id);

    return { data: response.data };
  },
});
