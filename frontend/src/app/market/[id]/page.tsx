import Link from 'next/link';
import { notFound } from 'next/navigation';

import BettingPanel from '@/components/BettingPanel';
import { mockMarkets } from '@/constants/mockData';
import { formatCurrency } from '@/utils/format';
import { getMarketById } from '@/utils/market';

interface MarketPageProps {
  params: Promise<{
    id: string;
  }>;
}

export default async function MarketPage({ params }: MarketPageProps) {
  const { id } = await params;
  const marketId = parseInt(id);
  const market = getMarketById(mockMarkets, marketId);

  if (!market) {
    notFound();
  }

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      {/* 返回按钮 */}
      <Link
        href="/"
        className="inline-flex items-center space-x-2 text-gray-400 hover:text-white transition-colors mb-8 cursor-pointer"
      >
        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
        </svg>
        <span>Back to Markets</span>
      </Link>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* 左侧：市场信息 */}
        <div className="lg:col-span-2 flex">
          <div className="card w-full">
            {/* 标题 */}
            <div className="mb-6">
              <h1 className="text-3xl md:text-4xl font-bold text-white mb-4">{market.question}</h1>
              <div className="flex items-center space-x-4">
                <span className="badge-primary">{market.category}</span>
                <div className="flex items-center text-gray-400">
                  <svg
                    className="w-5 h-5 mr-2"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"
                    />
                  </svg>
                  <span className="text-lg font-semibold">
                    {formatCurrency(market.volume)} Volume
                  </span>
                </div>
              </div>
            </div>

            {/* 市场选项 */}
            <div className="mb-8">
              <h2 className="text-2xl font-bold text-white mb-4">Market Options</h2>
              <div className="space-y-4">
                {/* Yes 选项 */}
                <div className="p-6 bg-[#0a0f1a] border-2 border-green-600/30 rounded-xl">
                  <div className="flex items-center justify-between mb-2">
                    <span className="text-2xl font-bold text-white">Yes</span>
                    <span className="text-3xl font-bold text-green-400">
                      {market.yesPercentage}%
                    </span>
                  </div>
                  <div className="progress-bar">
                    <div
                      className="progress-fill-success"
                      style={{ width: `${market.yesPercentage}%` }}
                    />
                  </div>
                  <p className="text-gray-400 mt-2">
                    {market.yesPercentage.toFixed(1)}% probability
                  </p>
                </div>

                {/* No 选项 */}
                <div className="p-6 bg-[#0a0f1a] border-2 border-red-600/30 rounded-xl">
                  <div className="flex items-center justify-between mb-2">
                    <span className="text-2xl font-bold text-white">No</span>
                    <span className="text-3xl font-bold text-red-400">{market.noPercentage}%</span>
                  </div>
                  <div className="progress-bar">
                    <div
                      className="progress-fill-danger"
                      style={{ width: `${market.noPercentage}%` }}
                    />
                  </div>
                  <p className="text-gray-400 mt-2">
                    {market.noPercentage.toFixed(1)}% probability
                  </p>
                </div>
              </div>
            </div>

            {/* 市场描述 */}
            <div>
              <h2 className="text-2xl font-bold text-white mb-4">About This Market</h2>
              <p className="text-gray-400 leading-relaxed">
                {market.description ||
                  'This is a prediction market where users can bet on the outcome of real-world events. The market will be resolved when the outcome is known.'}
              </p>
            </div>
          </div>
        </div>

        {/* 右侧：下注面板 */}
        <div className="lg:col-span-1 flex">
          <div className="w-full">
            <BettingPanel market={market} />
          </div>
        </div>
      </div>
    </div>
  );
}
