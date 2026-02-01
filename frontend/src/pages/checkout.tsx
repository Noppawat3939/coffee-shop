import { Button, Card, Divider, Flex, Stack, Typography } from "@mantine/core";
import { useNavigate, useSearch } from "@tanstack/react-router";
import { useCallback, useEffect, useState } from "react";
import { MainLayout, PromptpayQrcode } from "~/components";
import { isExpired } from "~/helper";
import { useAxios, useNotification } from "~/hooks";
import type { IOrder } from "~/interfaces/order.interface";
import type {
  ICreateTransactionResponse,
  IEnquiryTransactionResponse,
} from "~/interfaces/payment.interface";
import { order, payment } from "~/services";
import type { Response } from "~/services/service-instance";

type SearchParams = Partial<{
  order_number: string;
  transaction_number: string;
}>;

export default function CheckoutPage() {
  const search = useSearch({ strict: false }) satisfies SearchParams;

  const navigate = useNavigate();

  const [paymentExpired, setPaymentExpired] = useState(false);

  const { execute: getOrderByOrderNumber, data } = useAxios(
    order.getOrderByOrderNumber
  );

  const {
    execute: enquireTxn,
    data: txn,
    loading: fetchingTxn,
  } = useAxios(payment.enquireTransaction, {
    onSuccess: (res) => {
      const { data } = res as Response<IEnquiryTransactionResponse>;

      setPaymentExpired(isExpired(data?.expired_at));
    },
  });

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

        enquireTxn({ transaction_number: search.transaction_number });
      },
    }
  );

  const orderData = data?.data as IOrder;

  const [md, ctx] = useNotification();

  useEffect(() => {
    if (search?.order_number && search?.transaction_number) {
      getOrderByOrderNumber(search.order_number);
      enquireTxn({ transaction_number: search.transaction_number });
    }
  }, [search?.transaction_number]);

  const onExpired = useCallback(() => setPaymentExpired(true), []);

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
        {...txn?.data}
        onReCreateQR={onReCreateQR}
        paymentExpired={paymentExpired}
        onExpired={onExpired}
      />

      <Flex justify="center" mt={100} columnGap={12}>
        <Button
          loading={creating || fetchingTxn}
          disabled={paymentExpired}
          w={120}
          {...(!paymentExpired && { bg: "teal" })}
          onClick={() =>
            // md.open({ title: "updated order to paid", color: "teal" })
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

      {ctx}
    </MainLayout>
  );
}
