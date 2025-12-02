import Link from 'next/link';

export default function Footer() {
  return (
    <footer className="mt-20 border-t border-[#1e2937] bg-[#0a0f1a]">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="flex flex-col md:flex-row justify-between items-center space-y-4 md:space-y-0">
          <p className="text-gray-400 text-sm">
            © 2025 X402 PolyMarket. Powered by x402 Protocol.
          </p>
          <div className="flex space-x-6">
            <Link
              href="#"
              className="text-gray-400 hover:text-[#00d4ff] transition-colors text-sm cursor-pointer"
            >
              Terms
            </Link>
            <Link
              href="#"
              className="text-gray-400 hover:text-[#00d4ff] transition-colors text-sm cursor-pointer"
            >
              Privacy
            </Link>
            <Link
              href="#"
              className="text-gray-400 hover:text-[#00d4ff] transition-colors text-sm cursor-pointer"
            >
              Docs
            </Link>
          </div>
        </div>
      </div>
    </footer>
  );
}
