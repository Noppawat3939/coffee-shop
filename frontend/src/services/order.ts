import type { IEmployee } from "~/interfaces/employee.interface";
import { service as svc } from ".";
import type { Response } from "./service-instance";
import type {
  ICreateOrders,
  IOrder,
  IOrderJoin,
} from "~/interfaces/order.interface";

type GetOrderByOrderNumberResponse = IOrder & { employee: IEmployee };

const prefix = "Orders";

const createOrder = async (body: ICreateOrders) => {
  const { data } = await svc.post<Response<IOrder>>(`${prefix}`, body);
  return data;
};

export const getOrderByOrderNumber = async (order_number: string) => {
  const { data } = await svc.get<Response<GetOrderByOrderNumberResponse>>(
    `${prefix}/order-number/${order_number}`
  );
  return data;
};

export const getOrders = async <T extends object>(params: T) => {
  const { data } = await svc.get<Response<IOrderJoin[]>>(`${prefix}`, {
    params,
  });
  return data;
};

export default { createOrder, getOrderByOrderNumber, getOrders };
