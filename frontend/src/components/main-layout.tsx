import { Box, Flex, Typography } from "@mantine/core";
import type { PropsWithChildren, ReactNode } from "react";

type MainLayoutProps = Readonly<
  PropsWithChildren & Partial<{ title: string; extra: ReactNode }>
>;

export default function MainLayout({
  children,
  title,
  extra,
}: MainLayoutProps) {
  return (
    <Box aria-description="main-layout" h={"100dvh"} px={16} py={24}>
      <Flex direction="column">
        <Flex justify="space-between" align="center">
          {title && <Typography fz="h3">{title}</Typography>}
          {extra && extra}
        </Flex>
        {children}
      </Flex>
    </Box>
  );
}
