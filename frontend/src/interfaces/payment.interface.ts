import type { OrderStatus } from "./order.interface";

export interface ICreateTransaction {
  order_number: string;
}

export interface IEnquiryTransaction {
  transaction_number: string;
}

export interface IEnquiryTransactionResponse {
  transaction_number: string;
  order_number_ref: string;
  amount: number;
  status: OrderStatus;
  payment_code: string;
  expired_at: string;
  created_at: string;
  order: {
    id: number;
    order_number: string;
    total: number;
    status: OrderStatus;
  };
}

export interface IUpdateTransaction {
  orderNumber: string;
  status: OrderStatus;
}

export interface ICreateTransactionResponse {
  transaction_number: string;
  amount: number;
  status: OrderStatus;
  payment_code: string;
  expired_at: string;
  created_at: string;
}
