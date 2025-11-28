/**
 * 格式化数字为货币字符串
 * @param value - 要格式化的数字
 * @returns 格式化后的货币字符串，如 $1.2M, $500K
 */
export function formatCurrency(value: number): string {
  if (value >= 1000000) {
    return `$${(value / 1000000).toFixed(1)}M`;
  }
  if (value >= 1000) {
    return `$${(value / 1000).toFixed(1)}K`;
  }
  return `$${value.toFixed(2)}`;
}

/**
 * 格式化百分比
 * @param value - 百分比值（0-100）
 * @param decimals - 小数位数，默认为1
 * @returns 格式化后的百分比字符串
 */
export function formatPercentage(value: number, decimals: number = 1): string {
  return `${value.toFixed(decimals)}%`;
}
