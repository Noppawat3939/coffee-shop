import { StrictMode } from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createRouter } from "@tanstack/react-router";

import "./index.css";

import { routeTree } from "./routeTree.gen";
import { AppProvider } from "~/provider";
import { LoadingOverlay } from "@mantine/core";

// Create a new router instance
const router = createRouter({
  routeTree,
  defaultPendingComponent: () => (
    <LoadingOverlay
      visible
      c={"blue"}
      zIndex={1000}
      overlayProps={{ blur: 2 }}
    />
  ),
});

// Register the router instance for type safety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

// Render the app
const rootElement = document.getElementById("root")!;
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <StrictMode>
      <AppProvider>
        <RouterProvider router={router} />
      </AppProvider>
    </StrictMode>
  );
}
