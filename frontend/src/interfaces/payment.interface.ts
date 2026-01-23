import type { OrderStatus } from "./order.interface";

export interface ICreateTransaction {
  order_number: string;
}

export interface IEnquiryTransaction {
  transaction_number: string;
}

export interface IUpdateTransaction {
  orderNumber: string;
  status: OrderStatus;
}
