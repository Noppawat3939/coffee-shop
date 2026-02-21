import { Route } from "~/routes/transaction/orders";

export default function TransactionOrdersPage() {
  const { data } = Route.useLoaderData();

  console.log(data);
  return <div>TransactionOrdersPage</div>;
}
