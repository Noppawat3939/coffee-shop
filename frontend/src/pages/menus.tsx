import {
  Accordion,
  Typography,
  Flex,
  Image,
  ActionIcon,
  Modal,
  Button,
} from "@mantine/core";
import { useDisclosure, useMap } from "@mantine/hooks";
import { useNavigate } from "@tanstack/react-router";
import { Coffee } from "lucide-react";
import { Fragment, useEffect, useMemo } from "react";
import {
  BillOrders,
  IncreaseDecreaseInput,
  MainLayout,
  MemberChecker,
} from "~/components";
import { priceFormat, sum } from "~/helper";
import { useAxios } from "~/hooks";
import type { IMember } from "~/interfaces/member.interface";
import type { IMenu } from "~/interfaces/menu.interface";
import { menu } from "~/services";

export default function MenusPage() {
  const { execute, data } = useAxios(menu.getMenus);

  const navigation = useNavigate();

  const orderMap: Map<string, number> = useMap();

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
    // console.log(member);
    const searchParams = {} as Record<string, number>;
    console.log(orderMap.entries());
    for (const [key, amount] of orderMap.entries()) {
      const [, variationId] = key.split("_");

      if (amount > 0) {
        searchParams[`variation_id_${variationId}`] = amount;
      }
    }

    // navigation({ to: "/checkout", search: searchParams });
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
          withMemberClick={goToCheckout}
          noMemberClick={goToCheckout}
        />
      </Modal>
    </MainLayout>
  );
}
