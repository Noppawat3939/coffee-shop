import dayjs from "dayjs";
import { useCallback, useEffect, useRef, useState } from "react";
import duration from "dayjs/plugin/duration";

dayjs.extend(duration);

export default function useCountdown(exp?: string, onExpire?: () => void) {
  const expiredTs = useRef(dayjs(exp).valueOf());
  const timerRef = useRef<ReturnType<typeof setInterval> | null>(null);
  const expiredRef = useRef<boolean>(false);

  const calc = useCallback(() => {
    const diffMs = expiredTs.current - Date.now();
    return Math.max(Math.ceil(diffMs / 1000), 0);
  }, [expiredTs.current]);

  const [secondsLeft, setSecondsLeft] = useState(calc());

  useEffect(() => {
    expiredRef.current = false;
    expiredTs.current = dayjs(exp).valueOf();
    setSecondsLeft(calc());

    timerRef.current = setInterval(() => {
      setSecondsLeft((prev) => {
        if (prev <= 1) {
          if (timerRef.current) clearInterval(timerRef.current);
          expiredRef.current = true;
          onExpire?.();
          return 0;
        }
        return calc();
      });
    }, 1000);

    return () => {
      if (timerRef.current) {
        clearInterval(timerRef.current);
        timerRef.current = null;
      }
    };
  }, [exp, onExpire]);

  const formatted = dayjs
    .duration(secondsLeft, "seconds")
    .format(secondsLeft >= 3600 ? "HH:mm:ss" : "mm:ss");

  return [secondsLeft, formatted] as const;
}
