export const CATEGORIES = [
  'Trending',
  'New',
  'Politics',
  'Sports',
  'Finance',
  'Crypto',
  'Tech',
  'Culture',
] as const;

export type Category = (typeof CATEGORIES)[number];
