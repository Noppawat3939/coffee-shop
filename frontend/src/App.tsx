import { Button, Flex, Typography } from "@mantine/core";
import { Link } from "@tanstack/react-router";

export default function App() {
  return (
    <Flex direction="column" justify="center" align="center">
      <Typography c="green" style={{ fontSize: 80, fontWeight: "bold" }}>
        Coffee Shop
      </Typography>
      <Link to="/menus">
        <Button variant="outline">Menu</Button>
      </Link>
    </Flex>
  );
}
