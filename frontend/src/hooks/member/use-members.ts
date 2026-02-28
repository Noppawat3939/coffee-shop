import { useInputState } from "@mantine/hooks";
import { Route } from "~/routes/account/members";
import { useAxios } from "..";
import { member } from "~/services";
import { useRouter } from "@tanstack/react-router";
import { useTransition } from "react";

export default function useMembers() {
  const { invalidate } = useRouter();
  // intial loader
  const { data: intialData } = Route.useLoaderData();
  const [pending, startTransition] = useTransition();
  const [searchTerm, setSearchTerm] = useInputState("");

  const { execute, data, loading, reset } = useAxios(member.getMembers);

  const handleSearch = () => {
    if (!searchTerm) return;

    const isNumber = /^[0-9]+$/.test(searchTerm);

    execute({
      ...(isNumber ? { phone_number: searchTerm } : { full_name: searchTerm }),
    });
  };

  const handleReset = () => {
    startTransition(() => {
      invalidate();
    });
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
