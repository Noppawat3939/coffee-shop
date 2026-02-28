import { Table, type TableTdProps, type TableThProps } from "@mantine/core";
import type { ReactNode } from "react";

type Column<T> = {
  key: keyof T;
  header: ReactNode;
  data: T[];
  isIndex?: boolean;
  thProps?: Partial<TableThProps>;
  tdProps?: Partial<TableTdProps>;
  render?: (row: T, index: number) => ReactNode;
};

type CustomTableProps<T> = {
  columns: Column<T>[];
  actions?: (row: T, index: number) => ReactNode;
  actionsHeader?: ReactNode;
};

export default function CustomTable<TData>({
  columns,
  actions,
  actionsHeader = "Action",
}: CustomTableProps<TData>) {
  const [{ data = [] }] = columns;

  if (!data.length)
    return (
      <Table>
        <Table.Thead>
          <Table.Tr>
            {columns.map((col) => (
              <Table.Th key={`empt-table-header-${col.key.toString()}`}>
                {col.header}
              </Table.Th>
            ))}
          </Table.Tr>
        </Table.Thead>
      </Table>
    );

  return (
    <Table.ScrollContainer minWidth={200} type="native">
      <Table verticalSpacing="xs" stickyHeader stickyHeaderOffset={0}>
        <Table.Thead>
          <Table.Tr>
            {columns.map((col) => (
              <Table.Th
                key={`table-header-${String(col.key)}`}
                {...col.thProps}
              >
                {col.header}
              </Table.Th>
            ))}

            {actions && (
              <Table.Th ta="center" key="header-actions">
                {actionsHeader}
              </Table.Th>
            )}
          </Table.Tr>
        </Table.Thead>

        <Table.Tbody>
          {data.map((row, rowIndex) => (
            <Table.Tr key={rowIndex}>
              {columns.map((col) => (
                <Table.Td
                  key={`cell-${rowIndex}-${String(col.key)}`}
                  {...col.tdProps}
                >
                  {col.isIndex
                    ? rowIndex + 1
                    : col.render
                      ? col.render(row, rowIndex)
                      : col.key
                        ? String(row[col.key] ?? "")
                        : null}
                </Table.Td>
              ))}

              {actions && (
                <Table.Td key={`actions-${rowIndex}`}>
                  {actions(row, rowIndex)}
                </Table.Td>
              )}
            </Table.Tr>
          ))}
        </Table.Tbody>
      </Table>
    </Table.ScrollContainer>
  );
}
