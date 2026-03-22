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

    const token = Cookies.get(ACCESS_TOKEN_COOKIE_KEY);

    if (!token) {
      throw redirect({ to: "/login" });
    }

    try {
      if (!cacheData) {
        const res = await auth.verifyToken();
        cacheData = res.data;
      }

      return { data: cacheData };
    } catch (err) {
      console.error("Failed verify token:", err);

      const isUnauthorized =
        err instanceof AxiosError && err.status === HttpStatusCode.Unauthorized;

      if (!isUnauthorized) {
        await auth.employeeLogout();
        Cookies.remove(ACCESS_TOKEN_COOKIE_KEY);
        throw redirect({ to: "/login" });
      }

      try {
        // refresh token
        const res = await auth.revokeToken();
        const newToken = res.data.access_token;

        // expired in 3 hours
        const expires = new Date();
        expires.setHours(expires.getHours() + 3);

        Cookies.set(ACCESS_TOKEN_COOKIE_KEY, newToken, {
          expires,
          secure: true,
        });

        // retry verify after set new token
        const verify = await auth.verifyToken();
        cacheData = verify.data;

        return { data: cacheData };
      } catch {
        Cookies.remove(ACCESS_TOKEN_COOKIE_KEY);
        cacheData = null;

        throw redirect({ to: "/login" });
      }
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
