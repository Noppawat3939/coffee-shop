import { useCallback, useState } from "react";
import Cookies from "js-cookie";

type CookieOptions = {
  expires?: number | Date; // days or Date
  path?: string;
  secure?: boolean;
  sameSite?: "strict" | "lax" | "none";
};

export default function useCookie(key: string) {
  const [value, setValue] = useState<string | undefined>(() =>
    Cookies.get(key)
  );

  const setCookie = useCallback(
    (value: string, options: CookieOptions = {}) => {
      Cookies.set(key, value, {
        path: "/",
        ...options,
      });
      setValue(value);
    },
    [key]
  );

  const getCookie = useCallback(() => {
    const val = Cookies.get(key);
    setValue(val);
    return val;
  }, [key]);

  const removeCookie = useCallback(() => {
    Cookies.remove(key, { path: "/" });
    setValue(undefined);
  }, [key]);

  return {
    value,
    set: setCookie,
    get: getCookie,
    remove: removeCookie,
  };
}
