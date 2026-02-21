import {
  createRootRoute,
  Outlet,
  redirect,
  useRouterState,
} from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { AxiosError, HttpStatusCode } from "axios";
import { Fragment } from "react/jsx-runtime";
import { MaxWidthLayout, PortalLayout } from "~/components/layouts";
import {
  ACCESS_TOKEN_COOKIE_KEY,
  PORTAL_WITH_PATHNAME,
} from "~/helper/constant";
import { auth } from "~/services";
import Cookies from "js-cookie";
import type { TVerifyUserResponse } from "~/services/auth";

let cacheData: TVerifyUserResponse["data"] | null = null;

export const Route = createRootRoute({
  component: RootLayout,
  beforeLoad: async () => {
    try {
      if (cacheData === null) {
        const res = await auth.verifyToken();
        cacheData = res.data;
      }

      return { data: cacheData };
    } catch (err) {
      console.error(err);
      if (
        err instanceof AxiosError &&
        [HttpStatusCode.Unauthorized, HttpStatusCode.BadRequest].includes(
          err?.status as number
        )
      ) {
        // auto logout
        try {
          await auth.employeeLogout();
        } catch (err) {
          console.error(err);
        } finally {
          Cookies.remove(ACCESS_TOKEN_COOKIE_KEY);
          redirect({ to: "/login" });
        }
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
