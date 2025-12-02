'use client';

import { useState } from 'react';

import CategoryFilter from '@/components/CategoryFilter';
import MarketCard from '@/components/MarketCard';
import { mockMarkets } from '@/constants/mockData';
import { filterMarketsByCategory } from '@/utils/market';

import type { Category } from '@/constants/categories';

export default function MarketList() {
  const [activeCategory, setActiveCategory] = useState<Category | string>('Trending');

  const filteredMarkets = filterMarketsByCategory(mockMarkets, activeCategory);

  return (
    <div>
      {/* 分类过滤器 */}
      <div className="mb-8">
        <CategoryFilter activeCategory={activeCategory} onCategoryChange={setActiveCategory} />
      </div>

      {/* 市场列表 */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {filteredMarkets.map((market, index) => (
          <MarketCard key={market.id} market={market} index={index} />
        ))}
      </div>

      {filteredMarkets.length === 0 && (
        <div className="text-center py-12">
          <p className="text-gray-400 text-lg">No markets found in this category.</p>
        </div>
      )}
    </div>
  );
}
