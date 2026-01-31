import { Button, Card, Stack, Typography } from "@mantine/core";
import { useEffect, useRef } from "react";
import { dateFormat } from "~/helper";
import { useCountdown } from "~/hooks";
import type { IEnquiryTransactionResponse } from "~/interfaces/payment.interface";

type PromptpayQrcodeProps = Partial<IEnquiryTransactionResponse> & {
  paymentExpired: boolean;
  onReCreateQR: () => void;
  onExpired?: () => void;
};

export default function PromptpayQrcode({
  payment_code,
  expired_at,
  paymentExpired,
  onReCreateQR,
  onExpired,
}: PromptpayQrcodeProps) {
  const canvasRef = useRef<HTMLCanvasElement | null>(null);

  const [timeRef, timeRefFormatted] = useCountdown(expired_at, onExpired);

  useEffect(() => {
    if (!payment_code || !canvasRef.current) return;

    (async () => {
      const qrcode = (await import("qrcode")).default;
      const text = paymentExpired ? `${payment_code}-expired` : payment_code;
      await qrcode.toCanvas(canvasRef.current, text, {
        width: 150,
        margin: 1,
        errorCorrectionLevel: "M",
      });
    })();
  }, [payment_code]);

  return (
    <Stack align="center">
      <Card radius="md" withBorder>
        <canvas ref={canvasRef} />
        <Stack gap={0} align="center">
          <Typography fw="bold" fz="sm" {...(paymentExpired && { c: "gray" })}>
            {paymentExpired ? "expired" : "expired at:"}
          </Typography>
          <Typography fz="sm" {...(paymentExpired && { c: "gray" })}>
            {dateFormat(expired_at, "DD MMM YYYY, HH:mm")}
          </Typography>
          {timeRef > 0 && (
            <Typography
              fz="xs"
              c="red"
            >{`(${timeRefFormatted} minutes)`}</Typography>
          )}
        </Stack>
      </Card>
      {paymentExpired && (
        <Button onClick={onReCreateQR} top={10} size="sm" variant="outline">
          Re generate QR
        </Button>
      )}
    </Stack>
  );
}
