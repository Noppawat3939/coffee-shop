import { Button, Flex } from "@mantine/core";
import type { ComponentProps } from "react";

type TCustomButtonProps = ComponentProps<typeof Button>;

type SearchWithResetButtonProps = {
  searchProps: TCustomButtonProps;
  resetProps: TCustomButtonProps;
  gap?: number;
};

export default function SearchWithResetButton({
  searchProps,
  resetProps,
  gap = 10,
}: SearchWithResetButtonProps) {
  return (
    <Flex gap={gap}>
      <Button itemID="search-btn" w={100} {...searchProps}>
        Search
      </Button>
      <Button itemID="reset-btn" variant="outline" w={100} {...resetProps}>
        Reset
      </Button>
    </Flex>
  );
}
