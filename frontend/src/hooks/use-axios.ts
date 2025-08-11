import { useCallback, useState } from "react";

type AsyncFunc<Args extends unknown[], Return> = (
  ...args: Args
) => Promise<Return>;

export default function useAxios<Args extends unknown[], Return>(
  fn: AsyncFunc<Args, Return>
) {
  const [data, setData] = useState<Return | null>(null);
  const [loading, setLoading] = useState(false);

  const execute = useCallback(
    async (...args: Args) => {
      setLoading(true);
      try {
        const result = await fn(...args);

        setData(result);
        return result;
      } catch (err) {
        return null;
      } finally {
        setLoading(false);
      }
    },
    [fn]
  );

  return { data, loading, execute };
}
