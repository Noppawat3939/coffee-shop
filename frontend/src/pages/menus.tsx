import {
  Accordion,
  ActionIcon,
  Button,
  Flex,
  Image,
  Modal,
  Typography,
} from "@mantine/core";
import { useDisclosure, useMap } from "@mantine/hooks";
import { useNavigate } from "@tanstack/react-router";
import { Coffee } from "lucide-react";
import { Fragment, useEffect, useMemo, useRef } from "react";
import {
  BillOrders,
  IncreaseDecreaseInput,
  MainLayout,
  MemberChecker,
} from "~/components";
import { priceFormat, sum } from "~/helper";
import { useAxios } from "~/hooks";
import type { IMember } from "~/interfaces/member.interface";
import type { IMenu, IVariation } from "~/interfaces/menu.interface";
import type { ICreateOrders, IOrder } from "~/interfaces/order.interface";
import type { ICreateTransactionResponse } from "~/interfaces/payment.interface";
import { menu, order, payment } from "~/services";
import type { Response } from "~/services/service-instance";

type OrderKey = `${IMenu["id"]}_${IVariation["id"]}` | string;

export default function MenusPage() {
  const navigation = useNavigate();

  const orderNumberRef = useRef<string | null>(null);

  const { execute, data } = useAxios(menu.getMenus);

  const { execute: createTxnLog } = useAxios(payment.createTransaction, {
    onSuccess: (res, params) => {
      const order_number = params?.[0]?.order_number ?? "";
      const response = res as Response<ICreateTransactionResponse>;

      navigation({
        to: `/checkout?order_number=${order_number}&transaction_number=${response.data.transaction_number}`,
        reloadDocument: true,
      });
    },
  });

  const { execute: createOrder, loading } = useAxios(order.createOrder, {
    onSuccess: (res) => {
      const response = res as Response<IOrder>;
      const { order_number } = response.data;

      if (order_number) {
        orderNumberRef.current = order_number;
        createTxnLog({ order_number });
      }
    },
  });

  const orderMap: Map<OrderKey, number> = useMap(); // expected: {menuId_menuVariationId ==> amount}

  const [
    openedBillOrders,
    { open: onOpenBillOrders, close: onCloseBillOrders },
  ] = useDisclosure(false);
  const [
    openedMemberChecker,
    { open: onOpenMemberChecker, close: onCloseMemberChecker },
  ] = useDisclosure(false);

  useEffect(() => {
    return () => {
      execute();
    };
  }, []);

  const sumOrders = sum([...orderMap.values()] as number[]);

  const getSelectedOrders = useMemo(() => {
    const menus = data?.data || [];

    return menus
      .map((menu) => {
        const variations = menu.variations
          ?.map((v) => {
            const key = `${menu.id}_${v.id}`;
            const amount = orderMap.get(key) as number;

            if (amount > 0) {
              return { ...v };
            }
            return null;
          })
          .filter(Boolean) as typeof menu.variations;

        if (!variations?.length) return null;

        return { ...menu, variations };
      })
      .filter(Boolean) as IMenu[];
  }, [data?.data, orderMap.size]);

  const goToCheckout = (member?: IMember) => {
    const variations: ICreateOrders["variations"] = [];

    for (const [key, amount] of orderMap.entries()) {
      const [, variationId] = key.split("_");

      const menu_variation_id = +variationId;

      variations.push({ menu_variation_id, amount });
    }

    const params: ICreateOrders = {
      variations,
      ...(member?.full_name &&
        member?.id && {
          customer: member.full_name,
          member_id: member.id,
        }),
    };

    createOrder(params);
  };

  return (
    <MainLayout
      title={"Menu"}
      extra={
        <ActionIcon
          variant={"filled"}
          onClick={onOpenBillOrders}
          disabled={!sumOrders}
        >
          <Coffee width={14} />
        </ActionIcon>
      }
    >
      <Accordion mt={12} variant="filled" transitionDuration={300}>
        {data?.data &&
          data.data.map((item) => (
            <Accordion.Item key={item.id} value={item.id.toString()}>
              <Accordion.Control fz="h2" tt="capitalize">
                {item.name}
              </Accordion.Control>
              <Accordion.Panel>
                {item.variations &&
                  item.variations.map((variation, idx) => (
                    <Fragment key={`variation-${idx}`}>
                      <Flex py={4} px={8} columnGap={24} mb={5}>
                        {variation.image && (
                          <Image
                            src={variation.image}
                            alt="menu"
                            h={100}
                            w={100}
                            radius={12}
                            style={{ objectFit: "cover" }}
                            loading="lazy"
                          />
                        )}
                        <Flex direction="column" justify="center" rowGap={2}>
                          <Typography fz="h4" tt="capitalize">
                            {variation.type}
                          </Typography>
                          <Typography c="gray" opacity={0.7} fz="xs">
                            {item.description}
                          </Typography>
                          <Typography fz="h3" fw={500} tt="capitalize">
                            {priceFormat(variation.price)}
                          </Typography>
                        </Flex>
                        <IncreaseDecreaseInput
                          intialValue={0}
                          onChange={(value) => {
                            const name = `${variation.menu_id}_${variation.id}`;

                            orderMap.set(name, value);
                          }}
                        />
                      </Flex>
                    </Fragment>
                  ))}
              </Accordion.Panel>
            </Accordion.Item>
          ))}
      </Accordion>

      <Modal
        opened={openedBillOrders}
        onClose={onCloseBillOrders}
        overlayProps={{ blur: 2 }}
        title={
          <Typography fz="h4" fw={500}>
            {"Orders"}
          </Typography>
        }
      >
        <Modal.Body pt={0} px={6}>
          <BillOrders
            orders={getSelectedOrders}
            getAmountByOrder={(menuId, varId) =>
              orderMap.get(`${menuId}_${varId}`) || 0
            }
          />
        </Modal.Body>
        <Flex justify="center">
          <Button
            onClick={() => {
              onCloseBillOrders();
              onOpenMemberChecker();
            }}
          >
            {"Continue"}
          </Button>
        </Flex>
      </Modal>

      <Modal
        opened={openedMemberChecker}
        onClose={onCloseMemberChecker}
        overlayProps={{ blur: 2 }}
        title={
          <Typography fz="h4" fw={500}>
            {"You have membership ?"}
          </Typography>
        }
      >
        <MemberChecker
          loading={loading}
          withMemberClick={goToCheckout}
          noMemberClick={goToCheckout}
        />
      </Modal>
    </MainLayout>
  );
}
