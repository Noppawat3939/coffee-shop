import "@mantine/core/styles.css";

import { MantineProvider } from "@mantine/core";
import type { PropsWithChildren } from "react";

type AppProviderProps = Readonly<PropsWithChildren>;

export default function AppProvider({ children }: AppProviderProps) {
  return (
    <MantineProvider theme={{ primaryColor: "dark" }}>
      {children}
    </MantineProvider>
  );
}
