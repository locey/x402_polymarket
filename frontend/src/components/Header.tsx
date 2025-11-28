import Link from 'next/link';

import CustomConnectButton from '@/components/CustomConnectButton';
import SearchBox from '@/components/SearchBox';

export default function Header() {
  return (
    <header className="sticky top-0 z-50 bg-[#0a0f1a]/95 backdrop-blur-sm border-b border-[#1e2937]">
      <nav className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          {/* Logo */}
          <Link
            href="/"
            className="flex items-center space-x-2 hover:opacity-80 transition-opacity cursor-pointer"
          >
            <div className="w-8 h-8 bg-[#00d4ff] rounded-lg flex items-center justify-center font-bold text-white">
              X
            </div>
            <span className="text-xl font-bold text-white">X402 PolyMarket</span>
          </Link>

          {/* 搜索框 */}
          <SearchBox />

          {/* 右侧按钮 */}
          <div className="flex items-center space-x-4">
            <button className="hidden lg:block px-4 py-2 text-white hover:text-[#00d4ff] transition-colors font-medium cursor-pointer">
              How it works
            </button>
            <CustomConnectButton />
          </div>
        </div>
      </nav>
    </header>
  );
}
