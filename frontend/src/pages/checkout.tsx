import { Button, Card, Divider, Flex, Stack, Typography } from "@mantine/core";
import { useSearch } from "@tanstack/react-router";
import { useEffect } from "react";
import { MainLayout } from "~/components";
import { useAxios, useNotification } from "~/hooks";
import type { IOrder } from "~/interfaces/order.interface";
import { order } from "~/services";

type SearchParams = Partial<{ order_number: string }>;

export default function CheckoutPage() {
  const search = useSearch({ strict: false }) satisfies SearchParams;

  const { execute: getOrderByOrderNumber, data } = useAxios(
    order.getOrderByOrderNumber
  );

  const orderData = data?.data as IOrder;

  const [md, ctx] = useNotification();

  useEffect(() => {
    if (search?.order_number) {
      getOrderByOrderNumber(search.order_number);
    }
  }, [search?.order_number]);

  return (
    <MainLayout title="Checkout">
      <Stack gap={12} mt={10}>
        <Card withBorder>
          <Typography fw={500} c="gray" fz="sm">
            {"Customer infomation"}
          </Typography>
          <Typography>{`Customer: ${orderData?.customer}`}</Typography>
        </Card>
        {orderData?.order_menu_variations?.map((item, itemIdx) => (
          <Flex key={`item-${itemIdx}`} direction="column">
            <Flex justify="space-between" align="center">
              <Flex align="center" columnGap={8}>
                <Typography tt="capitalize">
                  {item.menu_variation.menu?.name}
                </Typography>
                <Typography fz="sm">{`x ${item.amount}`}</Typography>
              </Flex>
              <Typography>{item.price}</Typography>
            </Flex>
          </Flex>
        ))}
        <Divider />
        <Flex justify="space-between">
          <Typography>{"Total"}</Typography>
          <Typography>{orderData?.total}</Typography>
        </Flex>
      </Stack>

      <Flex justify="center" mt={100} columnGap={12}>
        <Button
          w={120}
          bg="teal"
          onClick={() =>
            md.open({ title: "updated order to paid", color: "teal" })
          }
        >
          {"Paid"}
        </Button>
        <Button w={120} variant="outline">
          {"Cancel"}
        </Button>
      </Flex>

      {ctx}
    </MainLayout>
  );
}
