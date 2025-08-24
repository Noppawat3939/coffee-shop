import { createFileRoute } from "@tanstack/react-router";
import MembershipPage from "~/pages/membership";

export const Route = createFileRoute("/membership")({
  component: MembershipPage,
});
