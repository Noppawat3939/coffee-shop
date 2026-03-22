import { useRouter } from "@tanstack/react-router";
import { Route } from "~/routes/profile";

export default function useProfile() {
  const { invalidate } = useRouter();
  const initial = Route.useLoaderData();
  const user = initial?.data;

  return { data: user, refetch: invalidate };
}
