import { Button, Card, Divider, Flex, Stack, Typography } from "@mantine/core";
import { useNavigate } from "@tanstack/react-router";
import { CircleCheck } from "lucide-react";
import { memo, useCallback } from "react";
import { MainLayout, PromptpayQrcode } from "~/components";
import { useAxios, useQueriesPaymentWithOrder } from "~/hooks";
import { OrderStatus } from "~/interfaces/order.interface";
import type { ICreateTransactionResponse } from "~/interfaces/payment.interface";
import { payment } from "~/services";
import type { Response } from "~/services/service-instance";

export default function CheckoutPage() {
  const navigate = useNavigate();

  const {
    isPaymentPAID,
    loading,
    onPaymentExpired,
    orderData,
    paymentExpired,
    refetchPayment,
    search,
    txnData,
  } = useQueriesPaymentWithOrder();

  const goToBill = useCallback(
    () =>
      navigate({
        to: "/bill/$order_number",
        params: { order_number: search.order_number },
      }),
    [search.order_number]
  );

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

  const { execute: updatePayment, loading: updating } = useAxios(
    payment.updateTransaction,
    {
      onSuccess: goToBill,
    }
  );

  const onReCreateQR = useCallback(
    () => createTransaction({ order_number: search.order_number }),
    [search?.order_number]
  );

  return (
    <MainLayout title="Checkout">
      <Stack gap={12} mt={10} mb={12}>
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

      {isPaymentPAID ? (
        <OrderPaidSection onView={goToBill} />
      ) : (
        <PromptpayQrcode
          {...txnData}
          onReCreateQR={onReCreateQR}
          paymentExpired={paymentExpired}
          onExpired={onPaymentExpired}
        />
      )}

      <Flex
        {...(isPaymentPAID && { style: { display: "none" } })}
        justify="center"
        mt={100}
        columnGap={12}
      >
        <Button
          loading={updating || creating || loading}
          disabled={paymentExpired}
          w={120}
          onClick={() =>
            updatePayment({
              orderNumber: search.order_number,
              status: OrderStatus.Paid,
            })
          }
          {...(!paymentExpired && { bg: "teal" })}
        >
          {"Paid"}
        </Button>
        <Button
          disabled={creating}
          loading={updating}
          w={120}
          variant="outline"
          onClick={() =>
            updatePayment({
              orderNumber: search.order_number,
              status: OrderStatus.Canceled,
            })
          }
        >
          {"Cancel"}
        </Button>
      </Flex>
    </MainLayout>
  );
}

const OrderPaidSection = memo(function (props: { onView: () => void }) {
  return (
    <Stack align="center">
      <Card w={200} withBorder>
        <Flex align="center" gap={5} mb={20}>
          <CircleCheck color="green" />
          <Typography display="flex">{" Order paid"}</Typography>
        </Flex>
        <Button onClick={props.onView}>{"View"}</Button>
      </Card>
    </Stack>
  );
});
