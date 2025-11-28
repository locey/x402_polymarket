'use client';

import { CATEGORIES } from '@/constants/categories';

import type { Category } from '@/constants/categories';

interface CategoryFilterProps {
  activeCategory: Category | string;
  onCategoryChange: (category: Category | string) => void;
}

export default function CategoryFilter({ activeCategory, onCategoryChange }: CategoryFilterProps) {
  return (
    <div className="flex items-center space-x-2 overflow-x-auto pb-2 scrollbar-hide">
      {CATEGORIES.map(category => (
        <button
          key={category}
          onClick={() => onCategoryChange(category)}
          className={`flex items-center space-x-2 px-4 py-2 rounded-lg font-medium whitespace-nowrap transition-all duration-200 cursor-pointer ${
            activeCategory === category
              ? 'bg-[#00d4ff] text-white'
              : 'bg-[#151d2a] text-gray-400 hover:text-white hover:bg-[#1e2937]'
          }`}
        >
          {category === 'Trending' && (
            <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"
              />
            </svg>
          )}
          <span>{category}</span>
        </button>
      ))}
    </div>
  );
}
