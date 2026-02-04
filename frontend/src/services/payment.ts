import { service as svc } from ".";
import type { Response } from "./service-instance";
import { randomUniqueID } from "~/helper";
import type { AxiosRequestConfig } from "axios";
import type {
  ICreateTransaction,
  ICreateTransactionResponse,
  IEnquiryTransactionResponse,
  IUpdateTransaction,
} from "~/interfaces/payment.interface";

const prefix = "Payments";

const buildHeaders = (): AxiosRequestConfig => ({
  headers: { ["X-Idempotency-Key"]: randomUniqueID() },
});

const createTransaction = async (body: ICreateTransaction) => {
  const { data } = await svc.post<Response<ICreateTransactionResponse>>(
    `${prefix}/txns/order`,
    body,
    buildHeaders()
  );
  return data;
};

const enquireTransaction = async (transaction_number: string) => {
  const { data } = await svc.post<Response<IEnquiryTransactionResponse>>(
    `${prefix}/txn/enquiry`,
    { transaction_number },
    buildHeaders()
  );
  return data;
};

const updateTransaction = async (body: IUpdateTransaction) => {
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
