import { Button, Flex, Stack } from "@mantine/core";
import { EmployeesTable } from "~/components/employees";
import { Route } from "~/routes/account/employees";

export default function EmployeesPage() {
  const { data } = Route.useLoaderData();

  return (
    <Stack>
      <Flex justify="end">
        <Button>Create</Button>
      </Flex>
      <EmployeesTable
        data={data}
        onRow={(act, row) => {
          console.log("handle", { act, ...row });
        }}
      />
    </Stack>
  );
}
