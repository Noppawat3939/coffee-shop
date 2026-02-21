export const ACCESS_TOKEN_COOKIE_KEY = "session";

export const PORTAL_WITH_PATHNAME = {
  payments: "transaction/payments",
  orders: "transaction/orders",
  employees: "account/employess",
  members: "account/members",
};

export const PORTAL_HEADER_WITH_PATHNAME = {
  [PORTAL_WITH_PATHNAME.payments]: "Transaction Payments",
  [PORTAL_WITH_PATHNAME.orders]: "Transaction Orders",
  [PORTAL_WITH_PATHNAME.employees]: "Employees Mangement",
  [PORTAL_WITH_PATHNAME.members]: "Members Management",
};
