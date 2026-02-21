import {
  Button,
  Card,
  Flex,
  Stack,
  Typography,
  type TypographyProps,
} from "@mantine/core";
import { useNavigate } from "@tanstack/react-router";
import { memo } from "react";
import { dateFormat, priceFormat } from "~/helper";
import { Route } from "~/routes/bill/$order_number";

type CustomTextProps = Readonly<
  Omit<TypographyProps, "children"> & { text?: string; isNumber?: boolean }
>;

export default function BillByOrderNumber() {
  const { data } = Route.useLoaderData();
  const navigate = useNavigate();

  return (
    <div>
      <Card mih={520} w={400} top={24} withBorder>
        <Typography fw="bolder" mb={12} fz="md">
          {"Order details"}
        </Typography>
        <Card withBorder p={12} mb={12}>
          <CustomText fz="xs" text={`order no: ${data.order_number}`} />
          <CustomText fz="xs" text={`payment: ${data.status}`} />
          <CustomText fz="xs" text={`employee: ${data.employee.name}`} />
          <CustomText fz="xs" text={`customer: ${data.customer}`} />
          <CustomText
            fz="xs"
            text={`created date: ${dateFormat(data.created_at, "DD MMM YYYY, HH:mm")}`}
          />
        </Card>
        <Card withBorder p={12}>
          <Stack gap={2} mt={12} mb={24}>
            {(data?.order_menu_variations || [])?.length > 0 &&
              data.order_menu_variations?.map((mv) => (
                <Flex justify="space-between" key={mv.id}>
                  <CustomText flex={0.55} text={mv.menu_variation.menu?.name} />
                  <CustomText
                    isNumber
                    flex={0.15}
                    text={mv.menu_variation.price.toString()}
                  />
                  <CustomText isNumber flex={0.15} text={`x ${mv.amount}`} />
                  <CustomText
                    isNumber
                    flex={0.15}
                    text={(mv.menu_variation.price * mv.amount).toString()}
                  />
                </Flex>
              ))}
          </Stack>
          <Flex justify="space-between">
            <CustomText fw="bold" flex={0.9} text="total" />
            <CustomText fw="bold" flex={0.1} text={priceFormat(data.total)} />
          </Flex>
        </Card>

        <Flex justify="center">
          <Button onClick={() => navigate({ to: "/menus" })} mt={200} w={150}>
            Back
          </Button>
        </Flex>
      </Card>
    </div>
  );
}

const CustomText = memo(function ({
  text,
  isNumber,
  ...rest
}: CustomTextProps) {
  return (
    <Typography fz="sm" {...(isNumber && { ta: "right" })} {...rest}>
      {text ?? ""}
    </Typography>
  );
});

CustomText.displayName = "CustomText";
