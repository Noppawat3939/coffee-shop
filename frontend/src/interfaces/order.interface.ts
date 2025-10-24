export interface ICreateOrders {
  customer?: string;
  variations: CreateOrderVariation[];
}

interface CreateOrderVariation {
  menu_variation_id: number;
  amount: number;
}
