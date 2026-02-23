import { Badge, Button, Flex, Table } from "@mantine/core";
import { Edit, Lock } from "lucide-react";
import type { IEmployee } from "~/interfaces/employee.interface";

type EmployeesTableProps = {
  data: IEmployee[];
  onRow: (action: "edit" | "active", data: IEmployee, index?: number) => void;
};

export default function EmployeesTable({
  data: employees,
  onRow,
}: EmployeesTableProps) {
  const rows = employees.map((employee, i) => (
    <Table.Tr key={employee.id}>
      <Table.Td>{i + 1}</Table.Td>
      <Table.Td>{employee.username}</Table.Td>
      <Table.Td>{employee.name}</Table.Td>
      <Table.Td>
        <Badge color="blue" variant="light">
          {employee.role}
        </Badge>
      </Table.Td>
      <Table.Td>{employee.created_at}</Table.Td>
      <Table.Td>
        <Badge variant="light" {...(employee?.active && { color: "teal" })}>
          {employee.active ? "active" : "de-active"}
        </Badge>
      </Table.Td>
      <Table.Td>
        <Flex justify="center" gap={5}>
          <Button
            itemID="edit-btn"
            onClick={() => onRow("edit", employee, i)}
            size="xs"
            variant="transparent"
            radius={20}
          >
            <Edit size="14" />
          </Button>
          <Button
            itemID="active-btn"
            onClick={() => onRow("active", employee, i)}
            size="xs"
            variant="transparent"
            radius={20}
          >
            <Lock size="14" />
          </Button>
        </Flex>
      </Table.Td>
    </Table.Tr>
  ));
  return (
    <Table.ScrollContainer minWidth={200} type="native">
      <Table verticalSpacing="xs" stickyHeader stickyHeaderOffset={0}>
        <Table.Thead>
          <Table.Tr>
            <Table.Th w={60}>No.</Table.Th>
            <Table.Th>Username</Table.Th>
            <Table.Th>Name</Table.Th>
            <Table.Th>Role</Table.Th>
            <Table.Th w={200}>Create date time</Table.Th>
            <Table.Th w={120}>Status</Table.Th>
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
