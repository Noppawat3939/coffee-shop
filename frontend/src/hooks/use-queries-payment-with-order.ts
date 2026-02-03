import { useCallback, useLayoutEffect, useMemo, useState } from "react";
import { Route } from "~/routes/checkout";
import { useAxios } from ".";
import { order, payment } from "~/services";
import type { IEnquiryTransactionResponse } from "~/interfaces/payment.interface";
import { isExpired } from "~/helper";
import type { Response } from "~/services/service-instance";
import { OrderStatus } from "~/interfaces/order.interface";

export default function useQueriesPaymentWithOrder() {
  const search = Route.useSearch();

  const {
    execute: getOrder,
    data: orderData,
    loading: fetchingOrder,
  } = useAxios(order.getOrderByOrderNumber);

  const {
    execute: getPayment,
    data: txnData,
    loading: fetchingPayment,
  } = useAxios(payment.enquireTransaction, {
    onSuccess: (res) => {
      const { data } = res as Response<IEnquiryTransactionResponse>;

      setPaymentExpired(isExpired(data.expired_at));
    },
  });

  const refetchPayment = (transaction_number: string) =>
    getPayment(transaction_number);

  const [paymentExpired, setPaymentExpired] = useState(false);

  useLayoutEffect(() => {
    if (!search || !search.order_number || !search.transaction_number) return;

    return () => {
      (async () => {
        await Promise.all([
          getPayment(search.transaction_number),
          getOrder(search.order_number),
        ]);
      })();
    };
  }, [search.transaction_number]);

  const onPaymentExpired = useCallback(() => setPaymentExpired(true), []);

  const isPaymentPAID = useMemo(
    () => txnData?.data?.status === OrderStatus.Paid,
    [txnData?.data?.status]
  );

  return {
    isPaymentPAID,
    loading: fetchingOrder || fetchingPayment,
    onPaymentExpired,
    orderData: orderData?.data,
    paymentExpired,
    refetchPayment,
    search,
    txnData: txnData?.data,
  };
}
