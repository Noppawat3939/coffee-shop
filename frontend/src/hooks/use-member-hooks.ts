import { useInputState } from "@mantine/hooks";
import { useRouter } from "@tanstack/react-router";
import { useTransition } from "react";
import { Route } from "~/routes/account/members";
import { useAxios } from ".";
import { member } from "~/services";
import { isNumberChar } from "~/helper";

// Query members
export default function useMembers() {
  const { invalidate } = useRouter();
  const { data: intialData } = Route.useLoaderData();

  const [pending, startTransition] = useTransition();
  const [searchTerm, setSearchTerm] = useInputState("");

  // service
  const { execute, data, loading, reset } = useAxios(member.getMembers);

  const handleSearch = () => {
    if (!searchTerm) return;

    execute({
      ...(isNumberChar(searchTerm)
        ? { phone_number: searchTerm }
        : { full_name: searchTerm }),
    });
  };

  const handleReset = () => {
    startTransition(invalidate);

    reset();
    setSearchTerm("");
  };

  const members = data?.data ?? intialData;

  return {
    data: members,
    handleReset,
    handleSearch,
    loading: pending || loading,
    searchTerm,
    setSearchTerm,
  };
}
