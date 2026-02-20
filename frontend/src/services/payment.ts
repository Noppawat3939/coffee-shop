import { service as svc } from ".";
import type { Response } from "./service-instance";
import { randomUniqueID } from "~/helper";
import type { AxiosRequestConfig } from "axios";
import type {
  ICreateTransaction,
  ICreateTransactionResponse,
  IEnquiryTransactionResponse,
  IUpdateTransaction,
  IPaymentTransactionsWithOrdersResponse,
} from "~/interfaces/payment.interface";

const prefix = "Payments/txns";

const buildHeaders = (): AxiosRequestConfig => ({
  headers: { ["X-Idempotency-Key"]: randomUniqueID() },
});

const createTransaction = async (body: ICreateTransaction) => {
  const { data } = await svc.post<Response<ICreateTransactionResponse>>(
    `${prefix}/order`,
    body,
    buildHeaders()
  );
  return data;
};

const enquireTransaction = async (transaction_number: string) => {
  const { data } = await svc.post<Response<IEnquiryTransactionResponse>>(
    `${prefix}/enquiry`,
    { transaction_number },
    buildHeaders()
  );
  return data;
};

const updateTransaction = async (body: IUpdateTransaction) => {
  const { data } = await svc.post<Response>(
    `${prefix}/${body.orderNumber}/${body.status}`,
    buildHeaders()
  );
  return data;
};

const getTransactions = async <T extends object>(params?: T) => {
  const { data } = await svc.get<
    Response<IPaymentTransactionsWithOrdersResponse[]>
  >(prefix, {
    params,
    ...buildHeaders(),
  });
  return data;
};

export default {
  createTransaction,
  enquireTransaction,
  updateTransaction,
  getTransactions,
};
