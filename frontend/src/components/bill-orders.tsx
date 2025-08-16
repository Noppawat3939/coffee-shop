import { Box, Divider, Flex, Stack, Typography } from "@mantine/core";
import { useMemo } from "react";
import { priceFormat, sum } from "~/helper";
import type { IMenu, IVariation } from "~/interfaces/menu.interface";

type BillOrdersProps = {
  orders: IMenu[];
  getAmountByOrder: (menuId: number, varationId: number) => number;
};

export default function BillOrders({
  orders,
  getAmountByOrder,
}: BillOrdersProps) {
  const total = useMemo(() => {
    const allPrice = orders.flatMap(
      (od) =>
        od?.variations &&
        od.variations.map((v) => getAmountByOrder(od.id, v.id) * v.price)
    ) as number[];

    return sum(allPrice);
  }, [orders.length, getAmountByOrder]);

  return (
    <Stack>
      {orders.map((order, oIdx) => (
        <Box key={`order-${oIdx}`}>
          <Typography fz="lg">{order.name}</Typography>

          {order.variations?.map((variation, vIdx) => {
            const varWithAmount = variation as IVariation & {
              amount: number;
            };

            const amount = getAmountByOrder(order.id, variation.id);

            return (
              <Flex justify="space-between" key={`variation-${vIdx}`}>
                <Flex px={6} align="center" columnGap={10}>
                  <Typography>{varWithAmount.type}</Typography>
                  <Typography
                    aria-label="amount"
                    fz="sm"
                  >{`x ${amount}`}</Typography>
                </Flex>
                <Typography>{variation.price * amount}</Typography>
              </Flex>
            );
          })}
        </Box>
      ))}
      <Divider />
      <Flex justify="space-between">
        <Typography>{"Total"}</Typography>
        <Typography>{priceFormat(total)}</Typography>
      </Flex>
    </Stack>
  );
}
