import type { IOrder, OrderStatus } from "./order.interface";

export interface IPaymentTransactionLog {
  // base payment
  id: number;
  order_id: number;
  amount: number;
  status: OrderStatus;
  payment_code: string;
  qr_signature: string;
  order_number_ref: string;
  transaction_number: string;
  expired_at: string;
  created_at: string;
}

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

export interface IPaymentTransactionsWithOrdersResponse
  extends IPaymentTransactionLog {
  orders?: IOrder;
}
