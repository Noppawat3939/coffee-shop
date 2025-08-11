import { Typography } from "@mantine/core";
import { useEffect } from "react";
import { useAxios } from "~/hooks";
import { apis } from "~/services";

export default function MenusPage() {
  const { execute, data } = useAxios(apis.getMenus);

  useEffect(() => {
    return () => {
      execute();
    };
  }, []);

  return (
    <div>
      {data?.data &&
        data.data.map((item, idx) => (
          <Typography key={`menu-${idx}`}>{item.name}</Typography>
        ))}
    </div>
  );
}
