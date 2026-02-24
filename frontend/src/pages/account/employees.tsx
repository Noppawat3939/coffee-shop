import {
  Badge,
  Button,
  Flex,
  Group,
  Stack,
  type ButtonProps,
} from "@mantine/core";
import { Edit, Lock } from "lucide-react";
import { CustomTable } from "~/components/table";
import { Route } from "~/routes/account/employees";

const actionBtnProps = {
  variant: "transparent",
  radius: 20,
  size: "xs",
} satisfies ButtonProps;

export default function EmployeesPage() {
  const { data } = Route.useLoaderData();

  return (
    <Stack>
      <Flex justify="end">
        <Button>Create</Button>
      </Flex>

      <CustomTable
        columns={[
          { key: "id", isIndex: true, header: "No.", data, thProps: { w: 60 } },
          { key: "username", header: "Username", data },
          { key: "name", header: "Name", data },
          { key: "role", header: "Role", data },
          {
            key: "created_at",
            header: "Create date time",
            data,
            thProps: { w: 200 },
          },
          {
            key: "active",
            header: "Status",
            thProps: { w: 120 },
            data,
            render: (row) => (
              <Badge variant="light" color="teal">
                {row ? "active" : "de-active"}
              </Badge>
            ),
          },
        ]}
        actionsHeader="Actions"
        actions={() => (
          <Group justify="center" gap="xs">
            <Button {...actionBtnProps}>
              <Edit size="14" />
            </Button>
            <Button size="xs" variant="transparent" radius={20}>
              <Lock size="14" />
            </Button>
          </Group>
        )}
      />
    </Stack>
  );
}
