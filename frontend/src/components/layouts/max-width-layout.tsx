import { Container, Stack } from "@mantine/core";
import type { PropsWithChildren } from "react";

type MaxWidthLayoutProps = Readonly<
  PropsWithChildren & { maxWidth?: number; pathname?: string }
>;

export default function MaxWidthLayout({
  children,
  maxWidth = 520,
}: MaxWidthLayoutProps) {
  return (
    <Container itemID="max-width-layout">
      <Stack styles={{ root: { margin: "0 auto" } }} maw={maxWidth}>
        {children}
      </Stack>
    </Container>
  );
}
