'use client';

import Link from 'next/link';

import { motion } from 'framer-motion';

import { formatCurrency } from '@/utils/format';

import type { Market } from '@/interfaces/market';

interface MarketCardProps {
  market: Market;
  index?: number;
}

export default function MarketCard({ market, index = 0 }: MarketCardProps) {
  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.4, delay: index * 0.05 }}
      whileHover={{ scale: 1.02 }}
      className="h-full"
    >
      <Link href={`/market/${market.id}`} className="block h-full">
        <div className="card-hover h-full flex flex-col">
          {/* 标题和信息 */}
          <div className="flex items-start justify-between mb-4">
            <div className="flex-1 min-h-[80px]">
              <h3 className="text-lg font-semibold text-white mb-2 hover:text-[#00d4ff] transition-colors line-clamp-2">
                {market.question}
              </h3>
              <div className="flex items-center space-x-3 text-sm">
                <span className="badge-primary">{market.category}</span>
                <div className="flex items-center text-gray-400">
                  <svg
                    className="w-4 h-4 mr-1"
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
                  <span>{formatCurrency(market.volume)} Vol.</span>
                </div>
              </div>
            </div>
            <motion.button
              whileHover={{ scale: 1.2 }}
              whileTap={{ scale: 0.9 }}
              onClick={e => {
                e.preventDefault();
                e.stopPropagation();
              }}
              className="text-gray-400 hover:text-[#00d4ff] transition-colors cursor-pointer"
            >
              <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
                />
              </svg>
            </motion.button>
          </div>

          {/* Yes 选项 */}
          <div className="space-y-3 flex-1">
            <div className="flex items-center justify-between text-sm mb-1">
              <span className="text-gray-400">Yes</span>
              <span className="text-white font-semibold">{market.yesPercentage}%</span>
            </div>
            <div className="flex items-center space-x-3">
              <motion.button
                whileHover={{ scale: 1.05 }}
                whileTap={{ scale: 0.95 }}
                onClick={e => {
                  e.preventDefault();
                  e.stopPropagation();
                }}
                className="flex-1 px-4 py-3 bg-green-600/20 hover:bg-green-600/30 border border-green-600/50 hover:border-green-600 text-green-400 font-semibold rounded-lg transition-all duration-200 cursor-pointer"
              >
                Yes
              </motion.button>
              <div className="px-4 py-3 bg-[#0a0f1a] border border-[#1e2937] rounded-lg text-white font-semibold min-w-[80px] text-center">
                {market.yesPercentage.toFixed(1)}%
              </div>
            </div>
          </div>

          {/* No 选项 */}
          <div className="space-y-3 mt-3">
            <div className="flex items-center justify-between text-sm mb-1">
              <span className="text-gray-400">No</span>
              <span className="text-white font-semibold">{market.noPercentage}%</span>
            </div>
            <div className="flex items-center space-x-3">
              <motion.button
                whileHover={{ scale: 1.05 }}
                whileTap={{ scale: 0.95 }}
                onClick={e => {
                  e.preventDefault();
                  e.stopPropagation();
                }}
                className="flex-1 px-4 py-3 bg-red-600/20 hover:bg-red-600/30 border border-red-600/50 hover:border-red-600 text-red-400 font-semibold rounded-lg transition-all duration-200 cursor-pointer"
              >
                No
              </motion.button>
              <div className="px-4 py-3 bg-[#0a0f1a] border border-[#1e2937] rounded-lg text-white font-semibold min-w-[80px] text-center">
                {market.noPercentage.toFixed(1)}%
              </div>
            </div>
          </div>
        </div>
      </Link>
    </motion.div>
  );
}
