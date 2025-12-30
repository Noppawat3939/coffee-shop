import { useLayoutEffect } from "react";

export default function useOneEffect(fn?: Function) {
  useLayoutEffect(() => {
    return () => {
      fn?.();
      console.warn("end");
    };
  }, []);
}
