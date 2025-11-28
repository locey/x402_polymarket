'use client';

import { motion } from 'framer-motion';

interface HomeSectionProps {
  children?: React.ReactNode;
}

export default function HomeSection({ children }: HomeSectionProps) {
  return (
    <section className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <motion.div
        initial={{ opacity: 0, y: 30 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6 }}
        className="text-center mb-12"
      >
        <motion.h1
          initial={{ opacity: 0, scale: 0.9 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ duration: 0.6, delay: 0.2 }}
          className="text-5xl md:text-6xl font-bold text-white mb-4"
        >
          <span className="bg-linear-to-r from-[#00d4ff] to-[#00b8e6] bg-clip-text text-transparent">
            X402 PolyMarket
          </span>
        </motion.h1>
        <motion.p
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ duration: 0.6, delay: 0.4 }}
          className="text-xl text-gray-400 max-w-2xl mx-auto"
        >
          Trade on real-world events with x402 instant payments. No accounts, no friction.
        </motion.p>
      </motion.div>

      {/* 市场列表 */}
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6, delay: 0.6 }}
      >
        {children}
      </motion.div>
    </section>
  );
}
