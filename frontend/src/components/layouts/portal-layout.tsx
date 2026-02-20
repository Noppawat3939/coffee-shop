import { Container } from "@mantine/core";
import type { PropsWithChildren } from "react";

type PortalLayoutProps = Readonly<PropsWithChildren>;

export default function PortalLayout({ children }: PortalLayoutProps) {
  return (
    <Container itemID="portal-layout" maw={2400} miw={1024}>
      {children}
    </Container>
  );
}
