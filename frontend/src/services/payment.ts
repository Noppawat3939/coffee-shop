import type { OrderStatus } from "~/interfaces/order.interface";
import { service as svc } from ".";
import type { Response } from "./service-instance";
import { randomUniqueID } from "~/helper";
import type { AxiosRequestConfig } from "axios";

const prefix = "Payment";

const buildHeaders = (): AxiosRequestConfig => ({
  headers: { ["X-Idempotency-Key"]: randomUniqueID() },
});

const createTransaction = async (body: { order_number: string }) => {
  const { data } = await svc.post<Response>(
    `${prefix}/txn/order`,
    body,
    buildHeaders()
  );
  return data;
};

const enquireTransaction = async (body: { transaction_number: string }) => {
  const { data } = await svc.post<Response>(
    `${prefix}/txn/enquiry`,
    body,
    buildHeaders()
  );
  return data;
};

const updateTransaction = async (body: {
  orderNumber: string;
  status: OrderStatus;
}) => {
  const { data } = await svc.post<Response>(
    `${prefix}/txn/${body.orderNumber}/${body.status}`,
    buildHeaders()
  );
  return data;
};

export default {
  createTransaction,
  enquireTransaction,
  updateTransaction,
};
