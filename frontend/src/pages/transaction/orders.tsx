import { TransactionOrdersTable } from "~/components/orders";
import { Route } from "~/routes/transaction/orders";

export default function TransactionOrdersPage() {
  const { data } = Route.useLoaderData();

  return (
    <div>
      <TransactionOrdersTable data={data} />
    </div>
  );
}
