import { Notification, type NotificationProps } from "@mantine/core";
import { useToggle } from "@mantine/hooks";
import { Fragment, useCallback, useMemo, useState, useTransition } from "react";

type OmmitedNotificationProps = Omit<
  NotificationProps,
  "pos" | "withCloseButton" | "pos" | "top" | "right" | "left" | "bottom"
>;

export default function useNotification(duration = 2500) {
  const [, startTranstion] = useTransition();

  const [show, toggle] = useToggle([false, true]);
  const [notiProps, setNotiProps] = useState<OmmitedNotificationProps>();

  const openChange = useCallback(() => {
    toggle();

    setTimeout(toggle, duration);
  }, [duration]);

  const props = {
    open: (state?: OmmitedNotificationProps) => {
      setNotiProps(state);

      startTranstion(openChange);
    },
  };

  const component = useMemo(
    () =>
      show ? (
        <Notification
          {...notiProps}
          withCloseButton={false}
          pos="fixed"
          top={20}
          right={20}
        />
      ) : (
        <Fragment />
      ),
    [notiProps, show]
  );
  return [props, component] as const;
}
