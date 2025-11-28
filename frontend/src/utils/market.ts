import type { Market } from '@/interfaces/market';

/**
 * 根据 ID 获取市场数据
 * @param markets - 市场数据数组
 * @param id - 市场 ID
 * @returns 找到的市场数据或 undefined
 */
export function getMarketById(markets: Market[], id: number): Market | undefined {
  return markets.find(market => market.id === id);
}

/**
 * 根据分类筛选市场
 * @param markets - 市场数据数组
 * @param category - 分类名称
 * @returns 筛选后的市场数组
 */
export function filterMarketsByCategory(markets: Market[], category: string): Market[] {
  if (category === 'Trending' || category === 'New') {
    return markets;
  }
  return markets.filter(market => market.category === category);
}

/**
 * 搜索市场
 * @param markets - 市场数据数组
 * @param query - 搜索关键词
 * @returns 搜索结果数组
 */
export function searchMarkets(markets: Market[], query: string): Market[] {
  const lowerQuery = query.toLowerCase();
  return markets.filter(
    market =>
      market.question.toLowerCase().includes(lowerQuery) ||
      market.category.toLowerCase().includes(lowerQuery)
  );
}
