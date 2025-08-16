import { useSearch } from "@tanstack/react-router";
import { MainLayout } from "~/components";

export default function CheckoutPage() {
  const search = useSearch({ strict: false }) as Record<string, number>;

  return (
    <MainLayout title="Checkout">
      checkout
      <br />
      {JSON.stringify(search, undefined, 2)}
    </MainLayout>
  );
}
