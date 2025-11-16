import { getDefaultConfig } from '@rainbow-me/rainbowkit';
import { arbitrum, base, mainnet, optimism, polygon, polygonAmoy, sepolia } from 'wagmi/chains';

export const config = getDefaultConfig({
  appName: 'X402 Prediction Market',
  projectId: 'YOUR_PROJECT_ID' /*|| process.env.NEXT_PUBLIC_PROJECT_ID!*/,
  chains: [
    mainnet,
    polygon,
    optimism,
    arbitrum,
    base,
    ...(process.env.NEXT_PUBLIC_ENABLE_TESTNETS === 'true' ? [sepolia, polygonAmoy] : []),
  ],
  ssr: true,
});
