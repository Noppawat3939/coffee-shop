import { Button, Flex, Image } from "@mantine/core";
import { Link } from "@tanstack/react-router";
import { MoveRight } from "lucide-react";

export default function IndexPage() {
  return (
    <Flex align="center" justify="center" h="100dvh" mx="auto">
      <Flex gap={24} direction="column" align="center" p={24}>
        <Image
          alt="banner"
          loading="lazy"
          src="https://images.unsplash.com/photo-1552010266-6458fda4d692?q=80&w=1025&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
          w={420}
        />
        <Link to="/menus">
          <Button variant="outline">
            <Flex gap={5} align="center">
              Menu
              <MoveRight style={{ width: 14 }} />
            </Flex>
          </Button>
        </Link>
      </Flex>
    </Flex>
  );
}
