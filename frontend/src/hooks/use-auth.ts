import { useState, useTransition } from "react";
import { useAxios, useCookie, useOneEffect } from ".";
import { auth } from "~/services";
import { HttpStatusCode } from "axios";
import { useNavigate } from "@tanstack/react-router";
import { ACCESS_TOKEN_COOKIE_KEY } from "~/helper/constant";
import type { TVerifyUserResponse } from "~/services/auth";

type AuthStatus = "idle" | "authenticated" | "unauthenticated";

type UseAuth = { username: string; loggedIn: boolean };

export default function useAuth(): UseAuth {
  const navigation = useNavigate();
  const [_, startTransition] = useTransition();

  const { remove } = useCookie(ACCESS_TOKEN_COOKIE_KEY);

  const [status, setStatus] = useState<AuthStatus>("idle");
  const [username, setUsername] = useState("");

  const { execute } = useAxios(auth.verifyToken, {
    onSuccess: (res) => {
      const response = res as TVerifyUserResponse;

      response.code === HttpStatusCode.Ok && setStatus("authenticated");
      setUsername(response.data.username);
    },
    onError: (e) => {
      if (
        e?.status === HttpStatusCode.Unauthorized ||
        e?.status === HttpStatusCode.BadRequest
      ) {
        setStatus("unauthenticated");
        remove();
        startTransition(() => navigation({ to: "/login" }));
      }
    },
  });

  useOneEffect(execute);

  const loggedIn = status === "authenticated";

  return { username, loggedIn };
}
