import {
  Badge,
  Button,
  Card,
  Drawer,
  Flex,
  Stack,
  Text,
  type BadgeProps,
  Tooltip,
} from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { useCallback, useState } from "react";
import { CustomTable } from "~/components/table";
import { dateFormat, isBeforeToDay, priceFormat } from "~/helper";
import { Route } from "~/routes/transaction/payments";

export default function TransactionPaymentsPage() {
  const { data: payments } = Route.useLoaderData();

  const [opened, { open, close }] = useDisclosure(false);

  const [paymentData, setPaymentData] = useState<
    (typeof payments)[number] | null
  >(null);

  const mappingPaymentStatusColor = useCallback(
    (status: string): BadgeProps["color"] => {
      const dict: Record<string, BadgeProps["color"]> = {
        paid: "cyan",
        canceled: "red",
        to_pay: "yellow",
      };

      return dict[status] ?? "dark";
    },
    []
  );

  return (
    <Stack>
      <CustomTable
        columns={[
          {
            key: "id",
            data: payments,
            isIndex: true,
            header: "No.",
            thProps: { w: 60 },
          },
          {
            key: "transaction_number",
            header: "Transaction ID",
            data: payments,
            thProps: { w: 200 },
            render: ({ transaction_number }) => (
              <Tooltip position="top-start" label={transaction_number}>
                <Text size="sm" w={200} truncate="end">
                  {transaction_number}
                </Text>
              </Tooltip>
            ),
          },
          {
            key: "order_number_ref",
            header: "Order Ref",
            data: payments,
            thProps: { w: 200 },
            render: ({ order_number_ref }) => (
              <Tooltip position="top-start" label={order_number_ref}>
                <Text size="sm" w={200} truncate="end">
                  {order_number_ref}
                </Text>
              </Tooltip>
            ),
          },
          {
            key: "amount",
            header: "Amount",
            data: payments,
          },
          {
            key: "created_at",
            header: "Create date time",
            data: payments,
          },
          {
            key: "expired_at",
            header: "Expired date time",
            data: payments,
          },
          {
            key: "status",
            header: "Status",
            data: payments,
            thProps: { w: 100 },
            render: ({ status }) => (
              <Badge
                w="100%"
                size="sm"
                variant="light"
                color={mappingPaymentStatusColor(status)}
              >
                {status}
              </Badge>
            ),
          },
        ]}
        actions={(data) => (
          <Button
            variant="outline"
            radius={20}
            size="xs"
            onClick={() => {
              setPaymentData(data);
              open();
            }}
          >
            View
          </Button>
        )}
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
