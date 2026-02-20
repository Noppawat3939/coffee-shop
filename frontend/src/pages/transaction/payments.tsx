import { Table } from "@mantine/core";
import { dateFormat } from "~/helper";
import { Route } from "~/routes/transaction/payments";

const DATE_FORMAT = "MMM DD YYYY, HH:mm";

export default function TransactionPaymentsPage() {
  const { data: payments } = Route.useLoaderData();

  const rows = payments.map((payment, i) => (
    <Table.Tr key={payment.id}>
      <Table.Td>{i + 1}</Table.Td>
      <Table.Td>{payment.transaction_number}</Table.Td>
      <Table.Td>{payment.amount}</Table.Td>
      <Table.Td>{dateFormat(payment.created_at, DATE_FORMAT)}</Table.Td>
      <Table.Td>{dateFormat(payment.expired_at, DATE_FORMAT)}</Table.Td>
      <Table.Td>{payment.status}</Table.Td>
    </Table.Tr>
  ));

  return (
    <Table stickyHeader stickyHeaderOffset={0}>
      <Table.Thead>
        <Table.Tr>
          <Table.Th>No.</Table.Th>
          <Table.Th>Transaction ID</Table.Th>
          <Table.Th>Amount</Table.Th>
          <Table.Th>Create date-time</Table.Th>
          <Table.Th>Expired date-ime</Table.Th>
          <Table.Th>Status</Table.Th>
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>{rows}</Table.Tbody>
      {/* <Table.Caption>Scroll page to see sticky thead</Table.Caption> */}
    </Table>
  );
}
