import { Button, Card, Divider, Flex, Stack, Typography } from "@mantine/core";
import { useNavigate } from "@tanstack/react-router";
import { useCallback } from "react";
import { MainLayout, PromptpayQrcode } from "~/components";
import { useAxios, useQueriesPaymentWithOrder } from "~/hooks";
import type { ICreateTransactionResponse } from "~/interfaces/payment.interface";
import { payment } from "~/services";
import type { Response } from "~/services/service-instance";

export default function CheckoutPage() {
  const navigate = useNavigate();

  const {
    search,
    orderData,
    txnData,
    loading,
    paymentExpired,
    refetchPayment,
    onPaymentExpired,
  } = useQueriesPaymentWithOrder();

  const { execute: createTransaction, loading: creating } = useAxios(
    payment.createTransaction,
    {
      onSuccess: (res) => {
        const { data } = res as Response<ICreateTransactionResponse>;

        if (!search.transaction_number) return;

        navigate({
          to: "/checkout",
          search: {
            order_number: search.order_number,
            transaction_number: data.transaction_number,
          },
          replace: true,
          reloadDocument: false,
        });

        refetchPayment(search.transaction_number);
      },
    }
  );

  const onReCreateQR = useCallback(() => {
    if (search.order_number) {
      createTransaction({ order_number: search.order_number });
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

      <PromptpayQrcode
        {...txnData}
        onReCreateQR={onReCreateQR}
        paymentExpired={paymentExpired}
        onExpired={onPaymentExpired}
      />

      <Flex justify="center" mt={100} columnGap={12}>
        <Button
          loading={creating || loading}
          disabled={paymentExpired}
          w={120}
          {...(!paymentExpired && { bg: "teal" })}
          onClick={() =>
            navigate({
              to: "/bill/$order_number",
              params: { order_number: search.order_number ?? "" },
            })
          }
        >
          {"Paid"}
        </Button>
        <Button w={120} variant="outline">
          {"Cancel"}
        </Button>
      </Flex>
    </MainLayout>
  );
}
