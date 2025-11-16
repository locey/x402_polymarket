import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  /* config options here */
  experimental: {
    // optimize package imports to reduce bundle size
    optimizePackageImports: ['@rainbow-me/rainbowkit', '@tanstack/react-query', 'viem', 'wagmi'],
  },
};

export default nextConfig;
