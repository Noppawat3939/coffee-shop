import { Stack } from "@mantine/core";
import type { PropsWithChildren } from "react";

type MaxWidthLayoutProps = Readonly<PropsWithChildren & { maxWidth?: number }>;

export default function MaxWidthLayout({
  children,
  maxWidth = 520,
}: MaxWidthLayoutProps) {
  return (
    <section itemID="max-width-layout" accessKey="">
      <Stack styles={{ root: { margin: "0 auto" } }} maw={maxWidth}>
        {children}
      </Stack>
    </section>
  );
}
