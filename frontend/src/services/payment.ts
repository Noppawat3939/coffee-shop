import { service as svc } from ".";
import type { Response } from "./service-instance";
const prefix = "payment";

const generatePromptpayQR = async (body: { amount: number }) => {
  const { data } = await svc.post<Response<{ qr: string }>>(
    `${prefix}/generate-promptpay-qr`,
    body
  );
  return data;
};

export default { generatePromptpayQR };
