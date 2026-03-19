import {
  createRootRoute,
  Outlet,
  redirect,
  useRouterState,
} from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { Fragment } from "react/jsx-runtime";
import { MaxWidthLayout, PortalLayout } from "~/components/layouts";
import {
  ACCESS_TOKEN_COOKIE_KEY,
  PORTAL_WITH_PATHNAME,
} from "~/helper/constant";
import { auth } from "~/services";
import Cookies from "js-cookie";
import type { TVerifyUserResponse } from "~/services/auth";
import { AxiosError, HttpStatusCode } from "axios";

let cacheData: TVerifyUserResponse["data"] | null = null;

export const Route = createRootRoute({
  component: RootLayout,
  beforeLoad: async ({ location }) => {
    if (location.pathname === "/login") return;

    try {
      if (cacheData === null) {
        const res = await auth.verifyToken();
        cacheData = res.data;
      }

      return { data: cacheData };
    } catch (err) {
      console.error(err);
      const hasToken = Cookies.get(ACCESS_TOKEN_COOKIE_KEY);

      if (hasToken) {
        const isUnauthorized =
          err instanceof AxiosError &&
          err.status === HttpStatusCode.Unauthorized;

        if (isUnauthorized) {
          const revoked = await auth.revokeToken();
          const newToken = revoked.data.access_token;

          Cookies.set(ACCESS_TOKEN_COOKIE_KEY, newToken, {
            expires: 1,
            secure: true,
          });
          return;
        }

        await auth.employeeLogout();
        Cookies.remove(ACCESS_TOKEN_COOKIE_KEY);
      }

      throw redirect({ to: "/login" });
    }
  },
});

function RootLayout() {
  const {
    location: { pathname },
  } = useRouterState();

  const context = Route.useRouteContext();

  const pathSliced = pathname.slice(1);

  const withPortal = Object.values(PORTAL_WITH_PATHNAME).includes(pathSliced);

  const Layout = withPortal ? PortalLayout : MaxWidthLayout;

  return (
    <Fragment>
      <Layout pathname={pathSliced} {...context}>
        <Outlet />
      </Layout>
      <TanStackRouterDevtools />
    </Fragment>
  );
}
