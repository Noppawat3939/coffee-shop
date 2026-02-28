import { Stack } from "@mantine/core";
import { CustomTable } from "~/components/table";
import { numberFormat } from "~/helper";
import { Route } from "~/routes/account/members";

export default function MembersPage() {
  const { data } = Route.useLoaderData();

  return (
    <Stack>
      <CustomTable
        columns={[
          { key: "id", isIndex: true, header: "No.", data, thProps: { w: 60 } },
          { key: "full_name", header: "Full name", data, thProps: { w: 150 } },
          {
            key: "phone_number",
            header: "Phone number",
            data,
            thProps: { w: 150 },
          },
          { key: "provider", header: "Provider", data, thProps: { w: 100 } },
          {
            key: "created_at",
            header: "Create date",
            data,
            thProps: { w: 150 },
          },
          {
            key: "member_point",
            header: "Total point",
            data,
            thProps: { w: 100 },
            render: (r) => numberFormat(r.member_point?.total_points) || "-",
          },
        ]}
      />
    </Stack>
  );
}
