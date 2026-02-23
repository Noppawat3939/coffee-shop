import {
  Badge,
  Button,
  Table,
  Text,
  Tooltip,
  type BadgeProps,
} from "@mantine/core";
import { useCallback } from "react";
import { priceFormat } from "~/helper";
import type { IPaymentTransactionsWithOrdersResponse } from "~/interfaces/payment.interface";

type TransactionPaymentsTableProps = {
  data: IPaymentTransactionsWithOrdersResponse[];
  onRow: (row: IPaymentTransactionsWithOrdersResponse, index?: number) => void;
};

export default function TransactionPaymentsTable({
  data: payments,
  onRow,
}: TransactionPaymentsTableProps) {
  const mappingPaymentStatusColor = useCallback(
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

  const rows = payments.map((payment, i) => (
    <Table.Tr key={payment.id}>
      <Table.Td>{i + 1}</Table.Td>
      <Table.Td>
        <Tooltip position="top-start" label={payment.transaction_number}>
          <Text size="sm" w={200} truncate="end">
            {payment.transaction_number}
          </Text>
        </Tooltip>
      </Table.Td>
      <Table.Td>
        <Tooltip position="top-start" label={payment.order_number_ref}>
          <Text size="sm" truncate="end" w={200}>
            {payment.order_number_ref}
          </Text>
        </Tooltip>
      </Table.Td>
      <Table.Td>{priceFormat(payment.amount)}</Table.Td>
      <Table.Td>{payment.created_at}</Table.Td>
      <Table.Td>{payment.expired_at}</Table.Td>
      <Table.Td>
        <Badge
          w="100%"
          size="sm"
          variant="light"
          color={mappingPaymentStatusColor(payment.status)}
        >
          {payment.status}
        </Badge>
      </Table.Td>
      <Table.Td ta="center">
        <Button
          variant="outline"
          radius={20}
          size="xs"
          onClick={() => onRow(payment, i)}
        >
          View
        </Button>
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <Table.ScrollContainer minWidth={200} type="native">
      <Table verticalSpacing="xs" stickyHeader stickyHeaderOffset={0}>
        <Table.Thead>
          <Table.Tr>
            <Table.Th w={60}>No.</Table.Th>
            <Table.Th w={200}>Transaction ID</Table.Th>
            <Table.Th w={200}>Order Ref</Table.Th>
            <Table.Th>Amount</Table.Th>
            <Table.Th>Create date time</Table.Th>
            <Table.Th>Expired date time</Table.Th>
            <Table.Th w={100}>Status</Table.Th>
            <Table.Th ta="center" w={120}>
              Actions
            </Table.Th>
          </Table.Tr>
        </Table.Thead>
        <Table.Tbody>{rows}</Table.Tbody>
      </Table>
    </Table.ScrollContainer>
  );
}
