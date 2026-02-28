import { Flex, Input, Stack } from "@mantine/core";
import { SearchIcon } from "lucide-react";
import { SearchWithResetButton } from "~/components";
import { CustomTable } from "~/components/table";
import { numberFormat } from "~/helper";
import { useMembers } from "~/hooks/member";

export default function MembersPage() {
  const {
    loading,
    setSearchTerm,
    handleReset,
    handleSearch,
    searchTerm,
    data,
  } = useMembers();

  return (
    <Stack>
      <Flex justify="space-between">
        <Input
          value={searchTerm}
          placeholder="Search Full name or Phone number"
          leftSection={<SearchIcon size={14} />}
          w={300}
          onChange={setSearchTerm}
          onKeyDown={(e) => e.key === "Enter" && handleSearch()}
        />

        <SearchWithResetButton
          searchProps={{
            onClick: handleSearch,
            loading,
          }}
          resetProps={{
            onClick: handleReset,
            disabled: loading,
          }}
        />
      </Flex>
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
