import { Button, Container, Stack, Title } from "@mantine/core";
import { useNavigate } from "@tanstack/react-router";
import { EmployeeRole } from "~/interfaces/employee.interface";
import { Route } from "~/routes/__root";

export default function Page() {
  const navigate = useNavigate();
  const context = Route.useRouteContext();

  return (
    <Container size="xs" style={{ height: "100vh" }}>
      <Stack justify="center" align="center" h="100%" gap="xl">
        <Title order={2}>Select Menu</Title>

        <Button
          size="lg"
          fullWidth
          radius="md"
          onClick={() => navigate({ to: "/menus" })}
        >
          Menu
        </Button>

        {context?.data.role !== EmployeeRole.staff && (
          <Button
            size="lg"
            fullWidth
            radius="md"
            variant="outline"
            onClick={() => navigate({ to: "/transaction/orders" })}
          >
            Admin
          </Button>
        )}

        <Button
          size="lg"
          fullWidth
          radius="md"
          variant="subtle"
          onClick={() => navigate({ to: "/profile" })}
        >
          Profile
        </Button>
      </Stack>
    </Container>
  );
}
