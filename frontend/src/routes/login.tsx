import { createFileRoute, redirect } from "@tanstack/react-router";
import Cookies from "js-cookie";
import { ACCESS_TOKEN_COOKIE_KEY } from "~/helper/constant";
import LoginPage from "~/pages/login";

export const Route = createFileRoute("/login")({
  beforeLoad: () => {
    const session = Cookies.get(ACCESS_TOKEN_COOKIE_KEY);

    if (session) return redirect({ to: "/home" });

    redirect({ to: "/login" });
  },
  component: () => <LoginPage />,
});
