import { Button, Flex, Stack, Typography } from "@mantine/core";
import { useSearch } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { MainLayout } from "~/components";
import { useAxios, useNotification } from "~/hooks";
import type { IVariation } from "~/interfaces/menu.interface";
import { menu } from "~/services";

type Order = IVariation & { amount: number };

export default function CheckoutPage() {
  const search = useSearch({ strict: false }) as Record<string, number>;

  const { execute: genVariations } = useAxios(menu.getVariations);
  // const { execute: genQR, loading } = useAxios(payment.generatePromptpayQR);

  const [orders, setOrders] = useState<Order[]>([]);

  const [md, ctx] = useNotification();

  const handleMappingOrders = async () => {
    const variationIdAmountMap = new Map();

    for (const [key, value] of Object.entries(search)) {
      const idString = key.replace("variation_id_", "");

      variationIdAmountMap.set(+idString, value);
    }

    const res = await genVariations({
      id: [...variationIdAmountMap.keys()],
    });

    const mappedOrders =
      res?.data &&
      res.data
        .map((od) => ({
          ...od,
          amount: variationIdAmountMap.get(od.id),
        }))
        .sort((a, b) => (a.menu?.name || "").localeCompare(b.menu?.name || ""));

    // const priceOrders =
    //   res?.data &&
    //   res.data.map((od) => od.price * variationIdAmountMap.get(od.id));

    // const sumPrice = priceOrders ? sum(priceOrders) : 0;
    // const qrRes = await genQR({ amount: sumPrice });

    // if (sumPrice && qrRes?.data.qr) {
    //   setTotalPrice(sumPrice);
    //   setQrPromptpayBase64(`data:image/png;base64,${qrRes?.data.qr}`);
    // }

    setOrders(mappedOrders || []);
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
        </Flex>
      </Stack>

      {/* {qrPromptpayBase64 && (
        <Flex justify="center" my={60}>
          <Card p={4} withBorder radius={12}>
            {loading ? (
              <Loader />
            ) : (
              <Image
                loading="lazy"
                alt="qr-promptpay"
                src={qrPromptpayBase64}
                w={120}
              />
            )}
          </Card>
        </Flex>
      )} */}

      <Flex justify="center" columnGap={12}>
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
