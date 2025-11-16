import DAppProvider from '@/components/Layouts/DAppProvider';
import UIProvider from '@/components/Layouts/UIProvider';

import './globals.css';

import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'X402 Prediction Market',
  description: 'This is a Hackathon Project.',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <DAppProvider>
          <UIProvider>{children}</UIProvider>
        </DAppProvider>
      </body>
    </html>
  );
}
