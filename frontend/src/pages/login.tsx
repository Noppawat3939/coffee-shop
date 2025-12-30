import { Button, Image, PasswordInput, Stack, TextInput } from "@mantine/core";
import { useInputState } from "@mantine/hooks";
import { useNavigate } from "@tanstack/react-router";
import { useTransition } from "react";
import { MainLayout } from "~/components";
import { useAxios, useCookie } from "~/hooks";
import type { IEmployeeLoggedIn } from "~/interfaces/auth.interface";
import { auth } from "~/services";
import type { Response } from "~/services/service-instance";

export const ACCESS_TOKEN_COOKIE_KEY = "cofshop_session";

export default function LoginPage() {
  const navigation = useNavigate();

  const { set } = useCookie(ACCESS_TOKEN_COOKIE_KEY);

  const [pending, startTransition] = useTransition();

  const [username, setUsername] = useInputState("");
  const [password, setPassword] = useInputState("");

  const { execute: login, loading } = useAxios(auth.employeeLogin, {
    onSuccess: (res) => {
      const response = res as Response<IEmployeeLoggedIn>;

      set(response.data.access_token, {
        expires: 1, // expired in 1 day
        secure: true,
      });

      startTransition(() => navigation({ to: "/menus" }));
    },
  });

  return (
    <MainLayout title="Login to shop">
      <form
        autoComplete="off"
        autoCapitalize="off"
        onSubmit={(e) => {
          e.preventDefault();

          login({ username, password });
        }}
      >
        <Stack gap={20}>
          <TextInput
            required
            label="Username"
            placeholder="Your username"
            value={username}
            onChange={setUsername}
          />
          <PasswordInput
            required
            label="Password"
            placeholder="Your password"
            value={password}
            onChange={setPassword}
          />
          <Button loading={loading || pending} type="submit">
            {"Login"}
          </Button>
        </Stack>
      </form>

      <Image
        mt={36}
        alt="banner"
        loading="lazy"
        src="https://images.unsplash.com/photo-1552010266-6458fda4d692?q=80&w=1025&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
      />
    </MainLayout>
  );
}
