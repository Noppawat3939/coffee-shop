import { service as svc } from ".";
import type { Response } from "./service-instance";
import type { ICreateOrders } from "~/interfaces/order.interface";

const prefix = "Orders";

const createOrder = async (body: ICreateOrders) => {
  const { data } = await svc.post<Response>(`${prefix}`, body);
  return data;
};

export default { createOrder };
