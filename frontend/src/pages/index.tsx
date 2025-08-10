import { Box, Button, Flex, Image, Typography } from "@mantine/core";
import { Link } from "@tanstack/react-router";
import { MoveRight } from "lucide-react";

export default function IndexPage() {
  return (
    <Flex align="center" justify="center" h="100dvh" mx="auto">
      <Flex
        gap={24}
        direction="column"
        align="center"
        p={24}
        style={{ border: "1px solid #101010" }}
      >
        <Box pos="relative">
          <Typography
            style={{ fontSize: 72 }}
            c="blue"
            pos="absolute"
            fw={500}
            bottom={-10}
            right={8}
          >
            Brew
          </Typography>
          <Image
            alt="banner"
            loading="lazy"
            src="https://plus.unsplash.com/premium_photo-1677661620509-b05f7a068d31?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
            style={{ objectFit: "contain" }}
            w={300}
          />
        </Box>
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
