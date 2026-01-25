import type { AxiosError } from "axios";
import { useCallback, useState } from "react";

type AsyncFunc<Args extends unknown[], Return> = (
  ...args: Args
) => Promise<Return>;

export type OveridedAxiosError = Readonly<
  AxiosError<Partial<{ message: string; code: number }>>
>;

type Options<Arg> = {
  onSuccess?: <Return>(data: Return, arg?: Arg) => void;
  onError?: (err?: OveridedAxiosError) => void;
  onFinish?: () => void;
};

export default function useAxios<Args extends unknown[], Return>(
  fn: AsyncFunc<Args, Return>,
  options?: Options<Args>
) {
  const [data, setData] = useState<Return | null>(null);
  const [loading, setLoading] = useState(false);

  const execute = useCallback(
    async (...args: Args) => {
      setLoading(true);
      try {
        const result = await fn(...args);

        setData(result);

        options?.onSuccess?.(result, args);

        return result;
      } catch (err) {
        options?.onError?.(err as OveridedAxiosError);
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
