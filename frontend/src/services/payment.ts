import { service as svc } from ".";
import type { Response } from "./service-instance";
import { randomUniqueID } from "~/helper";

const prefix = "Payment";

const generatePromptpayQR = async (body: { amount: number }) => {
  const { data } = await svc.post<Response<{ qr: string }>>(
    `${prefix}/generate-promptpay-qr`,
    body
  );
  return data;
};

const createTransaction = async (body: { order_number: string }) => {
  const { data } = await svc.post<Response>(`${prefix}/txn/order`, body, {
    headers: {
      ["X-Idempotency-Key"]: randomUniqueID(),
    },
  });
  return data;
};

const enquireTransaction = async (body: { transaction_number: string }) => {
  const { data } = await svc.post<Response>(`${prefix}/txn/enquiry`, body, {
    headers: { ["X-Idempotency-Key"]: randomUniqueID() },
  });
  return data;
};

export default { generatePromptpayQR, createTransaction, enquireTransaction };
