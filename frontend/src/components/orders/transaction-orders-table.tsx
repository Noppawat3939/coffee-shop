import { Badge, Table, Text, Tooltip, type BadgeProps } from "@mantine/core";
import { useCallback } from "react";
import { priceFormat } from "~/helper";
import type { IOrderJoin } from "~/interfaces/order.interface";

type TransactionOrdersTableProps = {
  data: IOrderJoin[];
};
export default function TransactionOrdersTable({
  data: orders,
}: TransactionOrdersTableProps) {
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

  const rows = orders.map((order, i) => (
    <Table.Tr key={order.id}>
      <Table.Td>{i + 1}</Table.Td>
      <Table.Td>
        <Tooltip position="top-start" label={order.order_number}>
          <Text size="sm" w={400} truncate="end">
            {order.order_number}
          </Text>
        </Tooltip>
      </Table.Td>

      <Table.Td>{priceFormat(order.total)}</Table.Td>

      <Table.Td>{order.employee?.username}</Table.Td>
      <Table.Td>
        {order.member ? order.member?.full_name : order.customer}
      </Table.Td>

      <Table.Td>{order.created_at}</Table.Td>

      <Table.Td>
        <Badge
          color={mappingOrderStatusColor(order.status)}
          w="100%"
          size="sm"
          variant="light"
        >
          {order.status}
        </Badge>
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <Table.ScrollContainer minWidth={200} type="native">
      <Table verticalSpacing="xs" stickyHeader stickyHeaderOffset={0}>
        <Table.Thead>
          <Table.Tr>
            <Table.Th w={60}>No.</Table.Th>
            <Table.Th>Order ID</Table.Th>
            <Table.Th>Amount</Table.Th>
            <Table.Th w={200}>Staff name</Table.Th>
            <Table.Th w={200}>Customer</Table.Th>
            <Table.Th w={200}>Create date time</Table.Th>
            <Table.Th w={120}>Status</Table.Th>
          </Table.Tr>
        </Table.Thead>
        <Table.Tbody>{rows}</Table.Tbody>
      </Table>
    </Table.ScrollContainer>
  );
}
