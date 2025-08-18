import { Code, Flex, Stack, Typography } from "@mantine/core";
import { useSearch } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { MainLayout } from "~/components";
import { priceFormat, sum } from "~/helper";
import { useAxios } from "~/hooks";
import type { IVariation } from "~/interfaces/menu.interface";
import { menu } from "~/services";

type Order = IVariation & { amount: number };

export default function CheckoutPage() {
  const search = useSearch({ strict: false }) as Record<string, number>;

  const { execute } = useAxios(menu.getVariations);

  const [orders, setOrders] = useState<Order[]>([]);
  const [totalPrice, setTotalPrice] = useState(0);

  const handleMappingOrders = async () => {
    const variationIdAmountMap = new Map();

    for (const [key, value] of Object.entries(search)) {
      const idString = key.replace("variation_id_", "");

      variationIdAmountMap.set(+idString, value);
    }

    const res = await execute({
      id: [...variationIdAmountMap.keys()].join(","),
    });

    const mappedOrders =
      res?.data &&
      res.data
        .map((od) => ({
          ...od,
          amount: variationIdAmountMap.get(od.id),
        }))
        .sort((a, b) => (a.menu?.name || "").localeCompare(b.menu?.name || ""));

    const priceOrders =
      res?.data &&
      res.data.map((od) => od.price * variationIdAmountMap.get(od.id));

    const sumPrice = priceOrders ? sum(priceOrders) : 0;

    setOrders(mappedOrders || []);
    setTotalPrice(sumPrice);
  };

  useEffect(() => {
    if (search) {
      handleMappingOrders();
    }
  }, [search]);

  return (
    <MainLayout title="Checkout">
      <Stack gap={12} mt={10}>
        {orders.map((order, oIdx) => (
          <Flex key={`order-${oIdx}`} direction="column">
            <Flex justify="space-between" align="center">
              <Flex align="center" columnGap={8}>
                <Typography tt="capitalize">{`${order.menu?.name} ${order.type}`}</Typography>
                <Typography fz="sm">{`x ${order.amount}`}</Typography>
              </Flex>
              <Typography>{order.price}</Typography>
            </Flex>
          </Flex>
        ))}
        <Flex justify="space-between">
          <Typography>{"Total"}</Typography>
          <Typography>{priceFormat(totalPrice)}</Typography>
        </Flex>
      </Stack>

      <Code mt={20}>QR Code generate is here</Code>
    </MainLayout>
  );
}
