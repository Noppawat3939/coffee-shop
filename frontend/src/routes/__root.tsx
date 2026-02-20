import {
  createRootRoute,
  Outlet,
  useRouterState,
} from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { Fragment } from "react/jsx-runtime";
import { MaxWidthLayout, PortalLayout } from "~/components/layouts";

const PORTAL_WITH_PATHNAME: string[] = ["transaction/payments"];

export const Route = createRootRoute({
  component: RootLayout,
});

function RootLayout() {
  const {
    location: { pathname },
  } = useRouterState();

  const withPortal = PORTAL_WITH_PATHNAME.includes(pathname.slice(1));

  const Layout = withPortal ? PortalLayout : MaxWidthLayout;

  return (
    <Fragment>
      <Layout>
        <Outlet />
      </Layout>
      <TanStackRouterDevtools />
    </Fragment>
  );
}
