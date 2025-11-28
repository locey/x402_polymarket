'use client';

import { useState } from 'react';

import type { Market } from '@/interfaces/market';

interface BettingPanelProps {
  market: Market;
  userBalance?: number;
}

export default function BettingPanel({ market, userBalance = 1000 }: BettingPanelProps) {
  const [selectedOption, setSelectedOption] = useState<'yes' | 'no'>('yes');
  const [betAmount, setBetAmount] = useState<string>('10');

  const quickAmounts = [10, 50, 100];

  const handlePlaceBet = () => {
    // TODO: 实现下注逻辑
    console.log('Placing bet:', { marketId: market.id, option: selectedOption, amount: betAmount });
  };

  return (
    <div className="card h-full">
      <h2 className="text-2xl font-bold text-white mb-6">Place Your Bet</h2>

      {/* 用户余额 */}
      <div className="mb-6">
        <div className="flex items-center justify-between mb-2">
          <span className="text-gray-400">Your Balance</span>
          <span className="text-2xl font-bold text-[#00d4ff]">${userBalance.toFixed(2)}</span>
        </div>
      </div>

      {/* 选择选项 */}
      <div className="mb-6">
        <label className="block text-gray-400 mb-3">Select Option</label>
        <div className="grid grid-cols-2 gap-4">
          <button
            onClick={() => setSelectedOption('yes')}
            className={`p-4 rounded-lg border-2 transition-all duration-200 cursor-pointer ${
              selectedOption === 'yes'
                ? 'border-green-600 bg-green-600/20 text-green-400'
                : 'border-[#1e2937] bg-[#0a0f1a] text-gray-400 hover:border-green-600/50'
            }`}
          >
            <div className="text-lg font-semibold mb-1">Yes</div>
            <div className="text-2xl font-bold">{market.yesPercentage}%</div>
          </button>
          <button
            onClick={() => setSelectedOption('no')}
            className={`p-4 rounded-lg border-2 transition-all duration-200 cursor-pointer ${
              selectedOption === 'no'
                ? 'border-red-600 bg-red-600/20 text-red-400'
                : 'border-[#1e2937] bg-[#0a0f1a] text-gray-400 hover:border-red-600/50'
            }`}
          >
            <div className="text-lg font-semibold mb-1">No</div>
            <div className="text-2xl font-bold">{market.noPercentage}%</div>
          </button>
        </div>
      </div>

      {/* 下注金额 */}
      <div className="mb-6">
        <label className="block text-gray-400 mb-3">Bet Amount ($)</label>
        <input
          type="number"
          value={betAmount}
          onChange={e => setBetAmount(e.target.value)}
          className="input mb-3"
          placeholder="Enter amount"
          min="1"
          max={userBalance}
        />
        <div className="flex space-x-2">
          {quickAmounts.map(amount => (
            <button
              key={amount}
              onClick={() => setBetAmount(amount.toString())}
              className="flex-1 px-4 py-2 bg-[#151d2a] hover:bg-[#1e2937] text-white rounded-lg border border-[#1e2937] hover:border-[#00d4ff] transition-all duration-200 cursor-pointer"
            >
              ${amount}
            </button>
          ))}
        </div>
      </div>

      {/* 预期收益 */}
      <div className="mb-6 p-4 bg-[#0a0f1a] border border-[#1e2937] rounded-lg">
        <div className="flex items-center justify-between mb-2">
          <span className="text-gray-400">Potential Payout</span>
          <span className="text-xl font-bold text-white">
            $
            {(
              parseFloat(betAmount || '0') /
              (selectedOption === 'yes' ? market.yesPrice : market.noPrice)
            ).toFixed(2)}
          </span>
        </div>
        <div className="flex items-center justify-between">
          <span className="text-gray-400">Potential Profit</span>
          <span className="text-lg font-semibold text-[#00d4ff]">
            $
            {(
              parseFloat(betAmount || '0') /
                (selectedOption === 'yes' ? market.yesPrice : market.noPrice) -
              parseFloat(betAmount || '0')
            ).toFixed(2)}
          </span>
        </div>
      </div>

      {/* 下注按钮 */}
      <button
        onClick={handlePlaceBet}
        disabled={!betAmount || parseFloat(betAmount) <= 0 || parseFloat(betAmount) > userBalance}
        className="w-full btn-primary disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer"
      >
        Place Bet
      </button>

      <p className="text-center text-gray-500 text-sm mt-4">Connect your wallet to start betting</p>
    </div>
  );
}
