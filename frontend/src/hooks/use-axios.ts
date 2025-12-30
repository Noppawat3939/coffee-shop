import type { AxiosError } from "axios";
import { useCallback, useState } from "react";

type AsyncFunc<Args extends unknown[], Return> = (
  ...args: Args
) => Promise<Return>;

type Options = {
  onSuccess?: <Return>(data: Return) => void;
  onError?: (err?: AxiosError) => void;
  onFinish?: () => void;
};

export default function useAxios<Args extends unknown[], Return>(
  fn: AsyncFunc<Args, Return>,
  options?: Options
) {
  const [data, setData] = useState<Return | null>(null);
  const [loading, setLoading] = useState(false);

  const execute = useCallback(
    async (...args: Args) => {
      setLoading(true);
      try {
        const result = await fn(...args);

        setData(result);
        options?.onSuccess?.(result);

        return result;
      } catch (err) {
        options?.onError?.(err as AxiosError);
        return null;
      } finally {
        setLoading(false);
        options?.onFinish?.();
      }
    },
    [fn]
  );

  return { data, loading, execute };
}
