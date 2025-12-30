import { useState, useTransition } from "react";
import { useAxios, useCookie, useOneEffect } from ".";
import { auth } from "~/services";
import { HttpStatusCode } from "axios";
import type { Response } from "~/services/service-instance";
import { useNavigate } from "@tanstack/react-router";
import { ACCESS_TOKEN_COOKIE_KEY as SESSION } from "~/pages/login";

type AuthStatus = "idle" | "authenticated" | "unauthenticated";

export default function useAuth() {
  const navigation = useNavigate();
  const [_, startTransition] = useTransition();

  const { remove } = useCookie(SESSION);

  const [status, setStatus] = useState<AuthStatus>("idle");

  const { execute } = useAxios(auth.verifyToken, {
    onSuccess: (res) => {
      const response = res as Response;

      response.code === HttpStatusCode.Ok && setStatus("authenticated");
    },
    onError: (e) => {
      if (e?.status === HttpStatusCode.Unauthorized) {
        setStatus("unauthenticated");
        remove();
        startTransition(() => navigation({ to: "/login" }));
      }
    },
  });

  useOneEffect(execute);

  const loggedIn = status === "authenticated";

  return loggedIn;
}
