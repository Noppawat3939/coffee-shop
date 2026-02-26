import { Stack, type BadgeProps, Badge } from "@mantine/core";
import { useCallback } from "react";
import { CustomTable } from "~/components/table";
import { priceFormat } from "~/helper";
import { Route } from "~/routes/transaction/orders";

export default function TransactionOrdersPage() {
  const { data } = Route.useLoaderData();

  const mappingOrderStatusColor = useCallback(
    (status: string): BadgeProps["color"] => {
      const dict: Record<string, BadgeProps["color"]> = {
        paid: "cyan",
        canceled: "red",
        to_pay: "yellow",
      };
      return dict[status] ?? "dark";
    },
    []
  );

  return (
    <Stack>
      <CustomTable
        columns={[
          { key: "id", isIndex: true, header: "No.", data, thProps: { w: 60 } },
          { key: "order_number", header: "Order ID", data },
          {
            key: "total",
            header: "Amount",
            data,
            render: ({ total }) => priceFormat(total),
          },
          { key: "employee", header: "Staff name", data, thProps: { w: 200 } },
          {
            key: "member",
            header: "Customer",
            data,
            thProps: { w: 200 },
            render: (r) => (r.member ? r.member.full_name : r.customer),
          },
          {
            key: "created_at",
            header: "Create date time",
            data,
            thProps: { w: 200 },
          },
          {
            key: "status",
            header: "Status",
            data,
            thProps: { w: 120 },
            render: ({ status }) => (
              <Badge
                w="100%"
                size="sm"
                variant="light"
                color={mappingOrderStatusColor(status)}
              >
                {status}
              </Badge>
            ),
          },
        ]}
      />
    </Stack>
  );
}
