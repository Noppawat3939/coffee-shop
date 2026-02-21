import {
  createRootRoute,
  Outlet,
  useRouterState,
} from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { Fragment } from "react/jsx-runtime";
import { MaxWidthLayout, PortalLayout } from "~/components/layouts";
import { PORTAL_WITH_PATHNAME } from "~/helper/constant";

export const Route = createRootRoute({
  component: RootLayout,
});

function RootLayout() {
  const {
    location: { pathname },
  } = useRouterState();

  const pathSliced = pathname.slice(1);

  const withPortal = Object.values(PORTAL_WITH_PATHNAME).includes(pathSliced);

  const Layout = withPortal ? PortalLayout : MaxWidthLayout;

  return (
    <Fragment>
      <Layout pathname={pathSliced}>
        <Outlet />
      </Layout>
      <TanStackRouterDevtools />
    </Fragment>
  );
}
