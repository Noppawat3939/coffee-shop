import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/account/members")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/account/members"!</div>;
}
