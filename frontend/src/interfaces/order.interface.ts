import type { IEmployee } from "./employee.interface";
import type { IMember } from "./member.interface";
import type { IVariation } from "./menu.interface";

export interface ICreateOrders {
  customer?: string;
  variations: ICreateOrderVariation[];
  member_id?: number;
}

interface ICreateOrderVariation {
  menu_variation_id: number;
  amount: number;
}

export interface IOrder {
  // base order
  id: number;
  order_number: string;
  status: string;
  customer: string;
  total: number;
  created_at: string;
  updated_at: string;
  status_logs?: IOrderStatusLog[];
  order_menu_variations?: IOrderMenuVariation[];
}

export interface IOrderMenuVariation {
  id: number;
  order_id: number;
  menu_variation_id: number;
  amount: number;
  price: number;
  menu_variation: IVariation;
}

export interface IOrderStatusLog {
  id: number;
  order_id: number;
  status: string;
  created_at: string;
}

export enum OrderStatus {
  ToPaid = "to_paid",
  Paid = "paid",
  Canceled = "canceled",
}

export interface IOrderJoin extends IOrder {
  member?: IMember;
  employee?: IEmployee;
}
