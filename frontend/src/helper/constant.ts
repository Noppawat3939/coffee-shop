export const ACCESS_TOKEN_COOKIE_KEY = "session";

export const PORTAL_WITH_PATHNAME = {
  payments: "transaction/payments",
  orders: "transaction/orders",
};

export const PORTAL_HEADER_WITH_PATHNAME = {
  [PORTAL_WITH_PATHNAME.payments]: "Transaction Payments",
  [PORTAL_WITH_PATHNAME.orders]: "Transaction Orders",
};
