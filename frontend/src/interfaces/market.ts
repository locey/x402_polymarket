export interface Market {
  id: number;
  question: string;
  category: string;
  volume: number;
  yesPercentage: number;
  noPercentage: number;
  yesPrice: number;
  noPrice: number;
  description?: string;
  endDate?: Date;
}

export interface BetOption {
  type: 'yes' | 'no';
  percentage: number;
  price: number;
}

export interface UserBet {
  marketId: number;
  option: 'yes' | 'no';
  amount: number;
  timestamp: Date;
}
