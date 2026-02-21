import { Card, Drawer, Flex, Stack, Text } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { useState } from "react";
import { TransactionPaymentsTable } from "~/components/payments";
import { dateFormat, isBeforeToDay, priceFormat } from "~/helper";
import { Route } from "~/routes/transaction/payments";

export default function TransactionPaymentsPage() {
  const { data: payments } = Route.useLoaderData();

  const [opened, { open, close }] = useDisclosure(false);

  const [paymentData, setPaymentData] = useState<
    (typeof payments)[number] | null
  >(null);

  return (
    <Stack>
      <TransactionPaymentsTable
        data={payments}
        onRow={(data) => {
          setPaymentData(data);
          open();
        }}
      />

      <Drawer
        title="Transaction Payment Details"
        styles={{ content: { minWidth: 960 } }}
        position="right"
        onClose={close}
        opened={opened}
        shadow="md"
      >
        <Stack gap={0}>
          <Card withBorder>
            {[
              {
                label: "Transaction ID",
                value: paymentData?.transaction_number,
              },
              { label: "Order Ref", value: paymentData?.order_number_ref },
              { label: "Amount", value: priceFormat(paymentData?.amount ?? 0) },
              {
                label: "Payment status",
                value: paymentData?.status.toUpperCase(),
              },
              {
                label: "Staff username",
                value: paymentData?.order?.employee?.username,
              },
              {
                label: "Staff name",
                value: paymentData?.order?.employee?.name,
              },
              { label: "Customer name", value: paymentData?.order?.customer },
              {
                label: "Member name",
                value: paymentData?.order?.member?.full_name || "-",
              },
              {
                label: "Member created date",
                value: paymentData?.order?.member?.created_at
                  ? dateFormat(
                      paymentData?.order?.member?.created_at,
                      "MMM DD YYYY"
                    )
                  : "-",
              },
              {
                label: "Payment code",
                value: paymentData?.payment_code,
              },
              {
                label: "Expiration date time",
                value: `${paymentData?.expired_at} ${isBeforeToDay(paymentData?.expired_at) ? "(expired)" : ""}`,
              },
              { label: "Create date time", value: paymentData?.created_at },
            ].map((item, i, items) => (
              <Flex
                key={`detail-${i}`}
                gap={0}
                mb={i === items.length - 1 ? 0 : 10}
              >
                <Text size="sm" miw={150} flex={0.2}>{`${item.label}:`}</Text>
                <Text size="sm" flex={0.8} display="flex">
                  {item.value}
                </Text>
              </Flex>
            ))}
          </Card>
        </Stack>
      </Drawer>
    </Stack>
  );
}
